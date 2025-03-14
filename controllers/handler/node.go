package handler

import (
	"context"
	"fmt"
	"strings"

	"github.com/go-logr/logr"
	wutongv1alpha1 "github.com/wutong-paas/wutong-operator/v2/api/v1alpha1"
	"github.com/wutong-paas/wutong-operator/v2/util/commonutil"
	"github.com/wutong-paas/wutong-operator/v2/util/constants"
	"github.com/wutong-paas/wutong-operator/v2/util/k8sutil"
	"github.com/wutong-paas/wutong-operator/v2/util/probeutil"
	"github.com/wutong-paas/wutong-operator/v2/util/wtutil"
	appsv1 "k8s.io/api/apps/v1"
	corev1 "k8s.io/api/core/v1"
	"k8s.io/apimachinery/pkg/api/resource"
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	"k8s.io/apimachinery/pkg/util/intstr"
	"sigs.k8s.io/controller-runtime/pkg/client"
)

// NodeName name for wt-node
var NodeName = "wt-node"
var NodeXDSServiceName = NodeName + "-xds"

type node struct {
	ctx    context.Context
	client client.Client
	log    logr.Logger

	labels     map[string]string
	etcdSecret *corev1.Secret
	cluster    *wutongv1alpha1.WutongCluster
	component  *wutongv1alpha1.WutongComponent

	pvcParametersRWX     *pvcParameters
	wtdataStorageRequest int64
}

var _ ComponentHandler = &node{}
var _ StorageClassRWXer = &node{}
var _ ResourcesCreator = &node{}
var _ Replicaser = &node{}

// NewNode creates a new wt-node handler.
func NewNode(ctx context.Context, client client.Client, component *wutongv1alpha1.WutongComponent, cluster *wutongv1alpha1.WutongCluster) ComponentHandler {
	return &node{
		ctx:                  ctx,
		client:               client,
		log:                  log.WithValues("Name: %s", component.Name),
		component:            component,
		cluster:              cluster,
		labels:               LabelsForWutongComponent(component),
		wtdataStorageRequest: getStorageRequest("WTDATA_STORAGE_REQUEST", 40),
	}
}

func (n *node) Before() error {
	secret, err := etcdSecret(n.ctx, n.client, n.cluster)
	if err != nil {
		return fmt.Errorf("failed to get etcd secret: %v", err)
	}
	n.etcdSecret = secret

	if n.component.Labels["persistentVolumeClaimAccessModes"] == string(corev1.ReadWriteOnce) {
		sc, err := storageClassNameFromWutongVolumeRWO(n.ctx, n.client, n.component.Namespace)
		if err != nil {
			return err
		}
		n.SetStorageClassNameRWX(sc)
		return nil
	}
	return setStorageCassName(n.ctx, n.client, n.component.Namespace, n)
}

func (n *node) Resources() []client.Object {
	return []client.Object{
		n.daemonSetForWutongNode(),
		n.serviceFroWutongNode(),
		n.service(),
	}
}

func (n *node) After() error {
	return nil
}

func (n *node) ListPods() ([]corev1.Pod, error) {
	return listPods(n.ctx, n.client, n.component.Namespace, n.labels)
}

func (n *node) SetStorageClassNameRWX(pvcParameters *pvcParameters) {
	n.pvcParametersRWX = pvcParameters
}

func (n *node) ResourcesCreateIfNotExists() []client.Object {
	if n.component.Labels["persistentVolumeClaimAccessModes"] == string(corev1.ReadWriteOnce) {
		return []client.Object{
			createPersistentVolumeClaimRWO(n.component.Namespace, constants.WTDataPVC, n.pvcParametersRWX, n.labels, n.wtdataStorageRequest),
		}
	}
	return []client.Object{
		// pvc is immutable after creation except resources.requests for bound claims
		createPersistentVolumeClaimRWX(n.component.Namespace, constants.WTDataPVC, n.pvcParametersRWX, n.labels, n.wtdataStorageRequest),
	}
}

func (n *node) Replicas() *int32 {
	nodeList := &corev1.NodeList{}
	if err := n.client.List(n.ctx, nodeList); err != nil {
		n.log.V(6).Info(fmt.Sprintf("list nodes: %v", err))
		return nil
	}
	return commonutil.Int32(int32(len(nodeList.Items)))
}

func (n *node) volumesByContainerRuntime(containerRuntime, sock string) ([]corev1.Volume, []corev1.VolumeMount) {
	var volumes []corev1.Volume
	var volumeMounts []corev1.VolumeMount
	switch containerRuntime {
	case constants.ContainerRuntimeContainerd:
		volumes = []corev1.Volume{
			{
				Name: "containerd-sock",
				VolumeSource: corev1.VolumeSource{
					HostPath: &corev1.HostPathVolumeSource{
						Path: sock,
						Type: k8sutil.HostPath(corev1.HostPathSocket),
					},
				},
			},
			{
				Name: "varlog",
				VolumeSource: corev1.VolumeSource{
					HostPath: &corev1.HostPathVolumeSource{
						Path: "/var/log", // for container logs
						Type: k8sutil.HostPath(corev1.HostPathDirectoryOrCreate),
					},
				},
			},
		}
		volumeMounts = []corev1.VolumeMount{
			{
				Name:      "containerd-sock", // default using containerd
				MountPath: sock,
			},
			{
				Name:      "varlog",
				MountPath: "/var/log",
			},
		}
	case constants.ContainerRuntimeDocker:
		volumes = []corev1.Volume{
			{
				Name: "docker",
				VolumeSource: corev1.VolumeSource{
					HostPath: &corev1.HostPathVolumeSource{
						Path: "/var/lib/docker",
						Type: k8sutil.HostPath(corev1.HostPathDirectoryOrCreate),
					},
				},
			},
			{
				Name: "vardocker",
				VolumeSource: corev1.VolumeSource{
					HostPath: &corev1.HostPathVolumeSource{
						Path: "/var/docker/lib",
						Type: k8sutil.HostPath(corev1.HostPathDirectoryOrCreate),
					},
				},
			},
			{
				Name: "docker-cert",
				VolumeSource: corev1.VolumeSource{
					HostPath: &corev1.HostPathVolumeSource{
						Path: constants.DefaultDockerCertsDir,
						Type: k8sutil.HostPath(corev1.HostPathDirectoryOrCreate),
					},
				},
			},
			{
				Name: "docker-sock",
				VolumeSource: corev1.VolumeSource{
					HostPath: &corev1.HostPathVolumeSource{
						Path: sock,
						Type: k8sutil.HostPath(corev1.HostPathSocket),
					},
				},
			},
		}
		volumeMounts = []corev1.VolumeMount{
			{
				Name:      "docker-sock",
				MountPath: sock,
			},
			{
				Name:      "docker", // for container logs, ubuntu
				MountPath: "/var/lib/docker",
			},
			{
				Name:      "vardocker", // for container logs, centos
				MountPath: "/var/docker/lib",
			},
			{
				Name:      "docker-cert",
				MountPath: constants.DefaultDockerCertsDir,
			},
		}
		if sock != constants.DockerSock {
			volumes = append(volumes, corev1.Volume{
				Name: "docker-base-sock",
				VolumeSource: corev1.VolumeSource{
					HostPath: &corev1.HostPathVolumeSource{
						Path: constants.DockerSock,
					},
				},
			})

			volumeMounts = append(volumeMounts, corev1.VolumeMount{
				Name:      "docker-base-sock",
				MountPath: constants.DockerSock,
			})
		}
	}

	return volumes, volumeMounts
}

func (n *node) daemonSetForWutongNode() client.Object {
	volumeMounts := []corev1.VolumeMount{
		{
			Name:      "wtdata",
			MountPath: "/wtdata",
		},
		{
			Name:      "sys",
			MountPath: "/sys",
		},
		{
			Name:      "etc",
			MountPath: "/newetc",
		},
		{
			Name:      "wtlocaldata",
			MountPath: "/wtlocaldata",
		},
	}
	volumes := []corev1.Volume{
		{
			Name: "wtdata",
			VolumeSource: corev1.VolumeSource{
				PersistentVolumeClaim: &corev1.PersistentVolumeClaimVolumeSource{
					ClaimName: constants.WTDataPVC,
				},
			},
		},
		{
			Name: "sys",
			VolumeSource: corev1.VolumeSource{
				HostPath: &corev1.HostPathVolumeSource{
					Path: "/sys",
					Type: k8sutil.HostPath(corev1.HostPathDirectory),
				},
			},
		},
		{
			Name: "etc",
			VolumeSource: corev1.VolumeSource{
				HostPath: &corev1.HostPathVolumeSource{
					Path: "/etc",
					Type: k8sutil.HostPath(corev1.HostPathDirectory),
				},
			},
		},
		{
			Name: "wtlocaldata",
			VolumeSource: corev1.VolumeSource{
				HostPath: &corev1.HostPathVolumeSource{
					Path: "/wtlocaldata",
					Type: k8sutil.HostPathDirectoryOrCreate(),
				},
			},
		},
	}

	args := []string{
		"--etcd=" + strings.Join(etcdEndpoints(n.cluster), ","),
		"--hostIP=$(POD_IP)",
		"--run-mode master",
		"--noderule manage,compute", // TODO: Let wt-node recognize itself
		"--nodeid=$(NODE_NAME)",
		"--image-repo-host=" + wtutil.GetImageRepository(n.cluster),
		"--hostsfile=/newetc/hosts",
		"--wt-ns=" + n.component.Namespace,
	}
	cr := k8sutil.GetContainerRuntime()
	vs, vms := n.volumesByContainerRuntime(cr.Name, cr.Endpoint)
	args = append(args, "--container-runtime="+cr.Name)
	args = append(args, "--runtime-endpoint="+cr.Endpoint)

	volumes = append(volumes, vs...)
	volumeMounts = append(volumeMounts, vms...)

	if n.etcdSecret != nil {
		volume, mount := volumeByEtcd(n.etcdSecret)
		volumeMounts = append(volumeMounts, mount)
		volumes = append(volumes, volume)
		args = append(args, etcdSSLArgs()...)
	}
	volumeMounts = mergeVolumeMounts(volumeMounts, n.component.Spec.VolumeMounts)
	volumes = mergeVolumes(volumes, n.component.Spec.Volumes)

	envs := []corev1.EnvVar{
		{
			Name: "POD_IP",
			ValueFrom: &corev1.EnvVarSource{
				FieldRef: &corev1.ObjectFieldSelector{
					FieldPath: "status.podIP",
				},
			},
		},
		{
			Name: "NODE_NAME",
			ValueFrom: &corev1.EnvVarSource{
				FieldRef: &corev1.ObjectFieldSelector{
					FieldPath: "spec.nodeName",
				},
			},
		},
		{
			Name:  "WT_NAMESPACE",
			Value: n.component.Namespace,
		},
	}
	if n.cluster.Spec.ImageHub == nil || n.cluster.Spec.ImageHub.Domain == constants.DefImageRepository {
		envs = append(envs, corev1.EnvVar{
			Name:  "WT_REGISTRY_SECRET",
			Value: hubImageRepository,
		})
	}
	envs = mergeEnvs(envs, n.component.Spec.Env)

	// prepare probe
	readinessProbe := probeutil.MakeReadinessProbeHTTP("", "/v2/ping", 6100)
	livenessProbe := probeutil.MakeLivenessProbeHTTP("", "/v2/ping", 6100)
	startupProbe := probeutil.MakeProbe(probeutil.ProbeKindHTTP, "", "/v2/ping", 6100, corev1.URISchemeHTTP, nil)
	probeutil.SetProbeArgs(startupProbe, 10, 10, 10, 1, 30)
	args = mergeArgs(args, n.component.Spec.Args)
	tolerations := []corev1.Toleration{
		{
			Operator: corev1.TolerationOpExists, // tolerate everything.
		},
	}

	if len(n.component.Spec.Tolerations) > 0 {
		tolerations = n.component.Spec.Tolerations
	}
	affinity := &corev1.Affinity{}
	if n.component.Spec.Affinity != nil {
		affinity = n.component.Spec.Affinity
	}
	resources := corev1.ResourceRequirements{
		Limits: corev1.ResourceList{
			corev1.ResourceMemory: resource.MustParse("2Gi"),
			corev1.ResourceCPU:    resource.MustParse("250m"),
		},
		Requests: corev1.ResourceList{
			corev1.ResourceMemory: resource.MustParse("256Mi"),
			corev1.ResourceCPU:    resource.MustParse("100m"),
		},
	}
	resources = mergeResources(resources, n.component.Spec.Resources)
	ds := &appsv1.DaemonSet{
		ObjectMeta: metav1.ObjectMeta{
			Name:      NodeName,
			Namespace: n.component.Namespace,
			Labels:    n.labels,
		},
		Spec: appsv1.DaemonSetSpec{
			Selector: &metav1.LabelSelector{
				MatchLabels: n.labels,
			},
			Template: corev1.PodTemplateSpec{
				ObjectMeta: metav1.ObjectMeta{
					Name:   NodeName,
					Labels: n.labels,
				},
				Spec: corev1.PodSpec{
					ImagePullSecrets:              imagePullSecrets(n.component, n.cluster),
					TerminationGracePeriodSeconds: commonutil.Int64(0),
					ServiceAccountName:            "wutong-operator",
					HostAliases:                   hostsAliases(n.cluster),
					HostPID:                       true,
					DNSPolicy:                     corev1.DNSClusterFirstWithHostNet,
					HostNetwork:                   true,
					Tolerations:                   tolerations,
					Affinity:                      affinity,
					Containers: []corev1.Container{
						{
							Name:            NodeName,
							Image:           n.component.Spec.Image,
							ImagePullPolicy: n.component.ImagePullPolicy(),
							Env:             envs,
							Args:            args,
							VolumeMounts:    volumeMounts,
							ReadinessProbe:  readinessProbe,
							LivenessProbe:   livenessProbe,
							StartupProbe:    startupProbe,
							Resources:       resources,
						},
					},
					PriorityClassName: constants.WutongPlatformComponentPriorityClassName,
					Volumes:           volumes,
				},
			},
		},
	}

	return ds
}

func (n *node) serviceFroWutongNode() client.Object {
	svc := &corev1.Service{
		ObjectMeta: metav1.ObjectMeta{
			Name:      NodeXDSServiceName,
			Namespace: n.component.Namespace,
			Labels:    n.labels,
		},
		Spec: corev1.ServiceSpec{
			Selector: n.labels,
			Ports: []corev1.ServicePort{
				{
					Name:       "tcp-6100",
					Protocol:   corev1.ProtocolTCP,
					Port:       6100,
					TargetPort: intstr.FromInt(6100),
				},
				{
					Name:       "tcp-6101",
					Protocol:   corev1.ProtocolTCP,
					Port:       6101,
					TargetPort: intstr.FromInt(6101),
				},
			},
		},
	}
	return svc
}

func (n *node) service() client.Object {
	return &corev1.Service{
		ObjectMeta: metav1.ObjectMeta{
			Name:      NodeName,
			Namespace: n.component.Namespace,
			Labels:    n.labels,
		},
		Spec: corev1.ServiceSpec{
			Ports: []corev1.ServicePort{
				{
					Name:       NodeName,
					Port:       6100,
					TargetPort: intstr.FromInt(6100),
				},
			},
			Selector: n.labels,
		},
	}
}
