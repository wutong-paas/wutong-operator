package handler

import (
	"context"
	"fmt"
	"path"
	"strings"

	"github.com/wutong-paas/wutong-operator/v2/util/probeutil"
	"github.com/wutong-paas/wutong-operator/v2/util/wtutil"

	"github.com/wutong-paas/wutong-operator/v2/util/commonutil"

	wutongv1alpha1 "github.com/wutong-paas/wutong-operator/v2/api/v1alpha1"
	"github.com/wutong-paas/wutong-operator/v2/util/constants"
	"github.com/wutong-paas/wutong-operator/v2/util/k8sutil"

	appsv1 "k8s.io/api/apps/v1"
	corev1 "k8s.io/api/core/v1"
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	"k8s.io/apimachinery/pkg/util/intstr"
	"sigs.k8s.io/controller-runtime/pkg/client"
)

// ChaosName name for wt-chaos
var ChaosName = "wt-chaos"

type chaos struct {
	ctx        context.Context
	client     client.Client
	component  *wutongv1alpha1.WutongComponent
	cluster    *wutongv1alpha1.WutongCluster
	labels     map[string]string
	db         *wutongv1alpha1.Database
	etcdSecret *corev1.Secret

	pvcParametersRWX     *pvcParameters
	cacheStorageRequest  int64
	wtdataStorageRequest int64
}

var _ ComponentHandler = &chaos{}
var _ StorageClassRWXer = &chaos{}
var _ Replicaser = &chaos{}

// NewChaos creates a new wt-chaos handler.
func NewChaos(ctx context.Context, client client.Client, component *wutongv1alpha1.WutongComponent, cluster *wutongv1alpha1.WutongCluster) ComponentHandler {
	return &chaos{
		ctx:                  ctx,
		client:               client,
		component:            component,
		cluster:              cluster,
		labels:               LabelsForWutongComponent(component),
		cacheStorageRequest:  getStorageRequest("CHAOS_CACHE_STORAGE_REQUEST", 10),
		wtdataStorageRequest: getStorageRequest("WTDATA_STORAGE_REQUEST", 40),
	}
}

func (c *chaos) Before() error {
	db, err := getDefaultDBInfo(c.ctx, c.client, c.cluster.Spec.RegionDatabase, c.component.Namespace, DBName)
	if err != nil {
		return fmt.Errorf("get db info: %v", err)
	}
	if db.Name == "" {
		db.Name = RegionDatabaseName
	}
	c.db = db

	secret, err := etcdSecret(c.ctx, c.client, c.cluster)
	if err != nil {
		return fmt.Errorf("failed to get etcd secret: %v", err)
	}
	c.etcdSecret = secret

	if c.component.Labels["persistentVolumeClaimAccessModes"] == string(corev1.ReadWriteOnce) {
		sc, err := storageClassNameFromWutongVolumeRWO(c.ctx, c.client, c.component.Namespace)
		if err != nil {
			return err
		}
		c.SetStorageClassNameRWX(sc)
		return nil
	}
	return setStorageCassName(c.ctx, c.client, c.component.Namespace, c)
}

func (c *chaos) Resources() []client.Object {
	return []client.Object{
		c.deployment(),
		c.service(),
		c.defaultMavenSetting(),
	}
}

func (c *chaos) After() error {
	return nil
}
func (c *chaos) ListPods() ([]corev1.Pod, error) {
	return listPods(c.ctx, c.client, c.component.Namespace, c.labels)
}

func (c *chaos) SetStorageClassNameRWX(pvcParametersRWX *pvcParameters) {
	c.pvcParametersRWX = pvcParametersRWX
}

func (c *chaos) ResourcesCreateIfNotExists() []client.Object {
	if c.component.Labels["persistentVolumeClaimAccessModes"] == string(corev1.ReadWriteOnce) {
		result := []client.Object{
			createPersistentVolumeClaimRWO(c.component.Namespace, constants.WTDataPVC, c.pvcParametersRWX, c.labels, c.wtdataStorageRequest),
		}
		if c.cluster.Spec.CacheMode != "hostpath" {
			result = append(result, createPersistentVolumeClaimRWO(c.component.Namespace, constants.CachePVC, c.pvcParametersRWX, c.labels, c.cacheStorageRequest))
		}
		return result
	} else {

		result := []client.Object{
			createPersistentVolumeClaimRWX(c.component.Namespace, constants.WTDataPVC, c.pvcParametersRWX, c.labels, c.wtdataStorageRequest),
		}
		if c.cluster.Spec.CacheMode != "hostpath" {
			result = append(result, createPersistentVolumeClaimRWX(c.component.Namespace, constants.CachePVC, c.pvcParametersRWX, c.labels, c.cacheStorageRequest))
		}
		return result
	}
}

func (c *chaos) Replicas() *int32 {
	return commonutil.Int32(int32(len(c.cluster.Spec.NodesForChaos)))
}

func (c *chaos) deployment() client.Object {
	volumeMounts := []corev1.VolumeMount{
		{
			Name:      "wtdata",
			MountPath: "/wtdata",
		},
		{
			Name:      "cache",
			MountPath: "/cache",
		},
		{
			Name:      "wtdata",
			MountPath: "/root/.ssh",
			SubPath:   "services/ssh",
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
	}
	if c.cluster.Spec.CacheMode == "hostpath" {
		volumes = append(volumes, corev1.Volume{
			Name: "cache",
			VolumeSource: corev1.VolumeSource{
				HostPath: &corev1.HostPathVolumeSource{
					Path: "/cache",
					Type: k8sutil.HostPath(corev1.HostPathDirectoryOrCreate),
				},
			},
		})
	} else {
		volumes = append(volumes, corev1.Volume{
			Name: "cache",
			VolumeSource: corev1.VolumeSource{
				PersistentVolumeClaim: &corev1.PersistentVolumeClaimVolumeSource{
					ClaimName: constants.CachePVC,
				},
			},
		})
	}
	args := []string{
		"--hostIP=$(POD_IP)",
		c.db.RegionDataSource(),
		"--etcd-endpoints=" + strings.Join(etcdEndpoints(c.cluster), ","),
		"--pvc-wtdata-name=" + constants.WTDataPVC,
		"--pvc-cache-name=" + constants.CachePVC,
		"--wt-namespace=" + c.component.Namespace,
		"--wt-repo=" + ResourceProxyName,
	}
	if c.cluster.Spec.CacheMode == "hostpath" {
		args = append(args, "--cache-mode=hostpath")
	}
	rt := k8sutil.GetContainerRuntime()
	vs, vms := volumesByContainerRuntime(rt.Name, rt.Endpoint)
	volumes = append(volumes, vs...)
	volumeMounts = append(volumeMounts, vms...)
	args = append(args, "--container-runtime="+rt.Name)
	args = append(args, "--runtime-endpoint="+rt.Endpoint)

	if c.etcdSecret != nil {
		volume, mount := volumeByEtcd(c.etcdSecret)
		volumeMounts = append(volumeMounts, mount)
		volumes = append(volumes, volume)
		args = append(args, etcdSSLArgs()...)
	}

	var nodeNames []string
	for _, node := range c.cluster.Spec.NodesForChaos {
		nodeNames = append(nodeNames, node.Name)
	}
	var affinity *corev1.Affinity
	if len(nodeNames) > 0 {
		affinity = affinityForRequiredNodes(nodeNames)
	}

	env := []corev1.EnvVar{
		{
			Name:  "CI_VERSION",
			Value: c.cluster.Spec.InstallVersion,
		},
		{
			Name: "POD_IP",
			ValueFrom: &corev1.EnvVarSource{
				FieldRef: &corev1.ObjectFieldSelector{
					FieldPath: "status.podIP",
				},
			},
		},
		{
			Name: "HOST_IP",
			ValueFrom: &corev1.EnvVarSource{
				FieldRef: &corev1.ObjectFieldSelector{
					FieldPath: "spec.nodeName",
				},
			},
		},
		{
			Name:  "SOURCE_DIR",
			Value: "/cache/source",
		},
		{
			Name:  "CACHE_DIR",
			Value: "/cache",
		},
		{
			Name: "IMAGE_PULL_SECRET",
			Value: func() string {
				if c.cluster.Status.ImagePullSecret != nil {
					return c.cluster.Status.ImagePullSecret.Name
				}
				return ""
			}(),
		},
	}
	if imageHub := c.cluster.Spec.ImageHub; imageHub != nil {
		env = append(env, corev1.EnvVar{
			Name:  "BUILD_IMAGE_REPOSTORY_DOMAIN",
			Value: path.Join(imageHub.Domain, imageHub.Namespace),
		})
		env = append(env, corev1.EnvVar{
			Name:  "BUILD_IMAGE_REPOSTORY_USER",
			Value: imageHub.Username,
		})
		env = append(env, corev1.EnvVar{
			Name:  "BUILD_IMAGE_REPOSTORY_PASS",
			Value: imageHub.Password,
		})
	}

	env = mergeEnvs(env, c.component.Spec.Env)
	volumeMounts = mergeVolumeMounts(volumeMounts, c.component.Spec.VolumeMounts)
	volumes = mergeVolumes(volumes, c.component.Spec.Volumes)
	args = mergeArgs(args, c.component.Spec.Args)

	// prepare probe
	readinessProbe := probeutil.MakeReadinessProbeHTTP("", "/v2/builder/health", 3228)
	ds := &appsv1.DaemonSet{
		ObjectMeta: metav1.ObjectMeta{
			Name:      ChaosName,
			Namespace: c.component.Namespace,
			Labels:    c.labels,
		},
		Spec: appsv1.DaemonSetSpec{
			Selector: &metav1.LabelSelector{
				MatchLabels: c.labels,
			},
			Template: corev1.PodTemplateSpec{
				ObjectMeta: metav1.ObjectMeta{
					Name:   ChaosName,
					Labels: c.labels,
				},
				Spec: corev1.PodSpec{
					TerminationGracePeriodSeconds: commonutil.Int64(0),
					ServiceAccountName:            "wutong-operator",
					ImagePullSecrets:              imagePullSecrets(c.component, c.cluster),
					Tolerations: []corev1.Toleration{
						{
							Operator: corev1.TolerationOpExists, // tolerate everything.
						},
					},
					PriorityClassName: constants.WutongPlatformComponentPriorityClassName,
					HostAliases:       hostsAliases(c.cluster),
					Affinity:          affinity,
					Containers: []corev1.Container{
						{
							Name:            ChaosName,
							Image:           c.component.Spec.Image,
							ImagePullPolicy: c.component.ImagePullPolicy(),
							Env:             env,
							Args:            args,
							VolumeMounts:    volumeMounts,
							ReadinessProbe:  readinessProbe,
							Resources:       c.component.Spec.Resources,
						},
					},
					Volumes: volumes,
				},
			},
		},
	}

	return ds
}

func (c *chaos) service() *corev1.Service {
	svc := &corev1.Service{
		ObjectMeta: metav1.ObjectMeta{
			Name:      ChaosName,
			Namespace: c.component.Namespace,
			Labels:    c.labels,
		},
		Spec: corev1.ServiceSpec{
			Ports: []corev1.ServicePort{
				{
					Name: "api",
					Port: 3228,
					TargetPort: intstr.IntOrString{
						IntVal: 3228,
					},
				},
			},
			Selector: c.labels,
		},
	}
	return svc
}

func (c *chaos) defaultMavenSetting() *corev1.ConfigMap {
	var mavensetting = `<settings xmlns="http://maven.apache.org/SETTINGS/1.0.0"
  xmlns:xsi="http://www.w3.org/2001/XMLSchema-instance"
  xsi:schemaLocation="http://maven.apache.org/SETTINGS/1.0.0
                      http://maven.apache.org/xsd/settings-1.0.0.xsd">
  <localRepository/>
  <interactiveMode/>
  <usePluginRegistry/>
  <offline/>
  <pluginGroups/>
  <servers/>
  <mirrors>
    <mirror>
     <id>aliyunmaven</id>
     <mirrorOf>central</mirrorOf>
     <name>阿里云公共仓库</name>
     <url>https://maven.aliyun.com/repository/central</url>
    </mirror>
    <mirror>
      <id>repo1</id>
      <mirrorOf>central</mirrorOf>
      <name>central repo</name>
      <url>http://repo1.maven.org/maven2/</url>
    </mirror>
    <mirror>
     <id>aliyunmaven</id>
     <mirrorOf>apache snapshots</mirrorOf>
     <name>阿里云阿帕奇仓库</name>
     <url>https://maven.aliyun.com/repository/apache-snapshots</url>
    </mirror>
  </mirrors>
  <proxies/>
  <activeProfiles/>
  <profiles>
    <profile>  
        <repositories>
           <repository>
                <id>aliyunmaven</id>
                <name>aliyunmaven</name>
                <url>https://maven.aliyun.com/repository/public</url>
                <layout>default</layout>
                <releases>
                        <enabled>true</enabled>
                </releases>
                <snapshots>
                        <enabled>true</enabled>
                </snapshots>
            </repository>
            <repository>
                <id>MavenCentral</id>
                <url>http://repo1.maven.org/maven2/</url>
            </repository>
            <repository>
                <id>aliyunmavenApache</id>
                <url>https://maven.aliyun.com/repository/apache-snapshots</url>
            </repository>
        </repositories>             
     </profile>
  </profiles>
</settings>
	`
	return &corev1.ConfigMap{
		ObjectMeta: metav1.ObjectMeta{
			Name:      "java-maven-aliyun",
			Namespace: c.component.Namespace,
			Labels: wtutil.LabelsForWutong(map[string]string{
				"configtype": "mavensetting",
				"default":    "true",
			}),
		},
		Data: map[string]string{
			"mavensetting": mavensetting,
		},
	}
}

func volumesByContainerRuntime(containerRuntime, sock string) ([]corev1.Volume, []corev1.VolumeMount) {
	var volumes []corev1.Volume
	var volumeMounts []corev1.VolumeMount

	switch containerRuntime {
	case constants.ContainerRuntimeContainerd:
		volumes = append(volumes, corev1.Volume{
			Name: "containerd-sock",
			VolumeSource: corev1.VolumeSource{
				HostPath: &corev1.HostPathVolumeSource{
					Path: sock,
				},
			},
		})
		volumeMounts = append(volumeMounts, corev1.VolumeMount{
			Name:      "containerd-sock",
			MountPath: sock,
		})

		volumes = append(volumes, corev1.Volume{
			Name: "containerd-hosts",
			VolumeSource: corev1.VolumeSource{
				HostPath: &corev1.HostPathVolumeSource{
					Path: constants.DefaultContainerdCertsDir,
					Type: k8sutil.HostPath(corev1.HostPathDirectoryOrCreate),
				},
			},
		})
		volumeMounts = append(volumeMounts, corev1.VolumeMount{
			Name:      "containerd-hosts",
			MountPath: constants.DefaultContainerdCertsDir,
		})

		if sock == constants.K3sContainerdSock {
			var k3sRegistryConfigPath = "/etc/rancher/k3s/"
			volumes = append(volumes, corev1.Volume{
				Name: "k3s-containerd-registry",
				VolumeSource: corev1.VolumeSource{
					HostPath: &corev1.HostPathVolumeSource{
						Path: k3sRegistryConfigPath,
						Type: k8sutil.HostPath(corev1.HostPathDirectoryOrCreate),
					},
				},
			}, corev1.Volume{
				Name: "k3s-containerd-config",
				VolumeSource: corev1.VolumeSource{
					HostPath: &corev1.HostPathVolumeSource{
						Path: constants.K3sContainerConfigPath,
					},
				},
			})
			volumeMounts = append(volumeMounts, corev1.VolumeMount{
				Name:      "k3s-containerd-registry",
				MountPath: k3sRegistryConfigPath,
			}, corev1.VolumeMount{
				Name:      "k3s-containerd-config",
				MountPath: constants.ContainerdConfigPath,
			})
		} else {
			volumes = append(volumes, corev1.Volume{
				Name: "containerd-config",
				VolumeSource: corev1.VolumeSource{
					HostPath: &corev1.HostPathVolumeSource{
						Path: constants.ContainerdConfigPath,
					},
				},
			})
			volumeMounts = append(volumeMounts, corev1.VolumeMount{
				Name:      "containerd-config",
				MountPath: constants.ContainerdConfigPath,
			})
		}
	case constants.ContainerRuntimeDocker:
		volumes = append(volumes, corev1.Volume{
			Name: "docker-sock",
			VolumeSource: corev1.VolumeSource{
				HostPath: &corev1.HostPathVolumeSource{
					Path: sock,
				},
			},
		})
		volumeMounts = append(volumeMounts, corev1.VolumeMount{
			Name:      "docker-sock",
			MountPath: sock,
		})
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
