package k8sutil

import (
	"context"
	"fmt"
	"net"
	"os"
	"strings"
	"sync"
	"time"

	"github.com/pelletier/go-toml"
	"github.com/sirupsen/logrus"
	"github.com/wutong-paas/wutong-operator/util/commonutil"
	"github.com/wutong-paas/wutong-operator/util/constants"
	corev1 "k8s.io/api/core/v1"
	networkingv1 "k8s.io/api/networking/v1"
	k8sErrors "k8s.io/apimachinery/pkg/api/errors"
	"k8s.io/apimachinery/pkg/api/resource"
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	"k8s.io/apimachinery/pkg/types"
	utilversion "k8s.io/apimachinery/pkg/util/version"
	"k8s.io/client-go/kubernetes"
	"k8s.io/client-go/kubernetes/scheme"
	"k8s.io/client-go/rest"
	"k8s.io/client-go/tools/reference"
	ctrl "sigs.k8s.io/controller-runtime"
	"sigs.k8s.io/controller-runtime/pkg/client"
)

var once sync.Once
var clientset kubernetes.Interface

// GetClientSet -
func GetClientSet() kubernetes.Interface {
	if clientset == nil {
		once.Do(func() {
			config := ctrl.GetConfigOrDie()
			clientset = kubernetes.NewForConfigOrDie(config)
		})
	}
	return clientset
}

// NewKubeConfig -
func NewKubeConfig() (*rest.Config, error) {
	cfg, err := InClusterConfig()
	if err != nil {
		return nil, err
	}
	return cfg, nil
}

// InClusterConfig -
func InClusterConfig() (*rest.Config, error) {
	// Work around https://github.com/kubernetes/kubernetes/issues/40973
	// See https://github.com/coreos/etcd-operator/issues/731#issuecomment-283804819
	if len(os.Getenv("KUBERNETES_SERVICE_HOST")) == 0 {
		addrs, err := net.LookupHost("kubernetes.default.svc")
		if err != nil {
			panic(err)
		}
		os.Setenv("KUBERNETES_SERVICE_HOST", addrs[0])
	}
	if len(os.Getenv("KUBERNETES_SERVICE_PORT")) == 0 {
		os.Setenv("KUBERNETES_SERVICE_PORT", "443")
	}
	cfg, err := rest.InClusterConfig()
	if err != nil {
		return nil, err
	}
	return cfg, nil
}

// HostPath returns a pointer to the HostPathType value passed in.
func HostPath(hostpath corev1.HostPathType) *corev1.HostPathType {
	return &hostpath
}

// IngressPathType returns a pointer to the PathType value passed in.
func IngressPathType(pathType networkingv1.PathType) *networkingv1.PathType {
	return &pathType
}

// HostPathDirectoryOrCreate returns a pointer to the HostPathType value passed in.
func HostPathDirectoryOrCreate() *corev1.HostPathType {
	var hpdoc = corev1.HostPathDirectoryOrCreate
	return &hpdoc
}

// MountPropagationMode returns a pointer to the MountPropagationMode value passed in.
func MountPropagationMode(moundPropagationMode corev1.MountPropagationMode) *corev1.MountPropagationMode {
	return &moundPropagationMode
}

// PersistentVolumeReclaimPolicy returns a pointer to the PersistentVolumeReclaimPolicy value passed in.
func PersistentVolumeReclaimPolicy(persistentVolumeReclaimPolicy corev1.PersistentVolumeReclaimPolicy) *corev1.PersistentVolumeReclaimPolicy {
	return &persistentVolumeReclaimPolicy
}

// UpdateCRStatus udpate cr status
func UpdateCRStatus(client client.Client, obj client.Object) error {
	ctx, cancel := context.WithTimeout(context.Background(), time.Second*5)
	defer cancel()
	err := client.Status().Update(ctx, obj)
	if err != nil {
		err = client.Status().Update(ctx, obj)
		if err != nil {
			return fmt.Errorf("update custom resource status: %v", err)
		}
	}
	return nil
}

// PersistentVolumeClaimForWTData -
func PersistentVolumeClaimForWTData(ns, claimName string, accessModes []corev1.PersistentVolumeAccessMode, labels map[string]string, storageClassName string, storageRequest int64) *corev1.PersistentVolumeClaim {
	size := resource.NewQuantity(storageRequest*1024*1024*1024, resource.BinarySI)
	pvc := &corev1.PersistentVolumeClaim{
		ObjectMeta: metav1.ObjectMeta{
			Name:      claimName,
			Namespace: ns,
			Labels:    labels,
		},
		Spec: corev1.PersistentVolumeClaimSpec{
			AccessModes: accessModes,
			Resources: corev1.ResourceRequirements{
				Requests: map[corev1.ResourceName]resource.Quantity{
					corev1.ResourceStorage: *size,
				},
			},
			StorageClassName: commonutil.String(storageClassName),
		},
	}

	return pvc
}

// EventsForPersistentVolumeClaim -
func EventsForPersistentVolumeClaim(pvc *corev1.PersistentVolumeClaim) (*corev1.EventList, error) {
	clientset := GetClientSet()
	ref, err := reference.GetReference(scheme.Scheme, pvc)
	if err != nil {
		return nil, err
	}
	ref.Kind = ""
	if _, isMirrorPod := pvc.Annotations[corev1.MirrorPodAnnotationKey]; isMirrorPod {
		ref.UID = types.UID(pvc.Annotations[corev1.MirrorPodAnnotationKey])
	}
	events, err := clientset.CoreV1().Events(pvc.GetNamespace()).Search(scheme.Scheme, ref)
	return events, err
}

// IsPodReady checks if the given pod is ready or not.
func IsPodReady(pod *corev1.Pod) bool {
	for _, condition := range pod.Status.Conditions {
		if condition.Type == corev1.PodReady && condition.Status == corev1.ConditionTrue {
			return true
		}
	}
	return false
}

// IsPodCompleted checks if the given pod is ready or not.
func IsPodCompleted(pod *corev1.Pod) bool {
	for _, condition := range pod.Status.Conditions {
		if condition.Type == corev1.PodReady && condition.Status == corev1.ConditionFalse && condition.Reason == "PodCompleted" {
			return true
		}
	}
	return false
}

// CreateIfNotExists -
func CreateIfNotExists(ctx context.Context, c client.Client, obj client.Object) error {
	err := c.Get(ctx, types.NamespacedName{Name: obj.GetName(), Namespace: obj.GetNamespace()}, obj)
	if err != nil {
		if !k8sErrors.IsNotFound(err) {
			return err
		}
		return c.Create(ctx, obj)
	}
	return nil
}

// ListNodes returns all nodes.
func ListNodes(ctx context.Context, c client.Client) ([]corev1.Node, error) {
	nodeList := &corev1.NodeList{}
	if err := c.List(ctx, nodeList); err != nil {
		return nil, err
	}
	return nodeList.Items, nil
}

// GetKubeVersion returns the version of k8s
func GetKubeVersion() *utilversion.Version {
	var serverVersion, err = GetClientSet().Discovery().ServerVersion()
	if err != nil {
		logrus.Errorf("Get Kubernetes Version failed [%+v]", err)
		return utilversion.MustParseSemantic("v1.19.6")
	}
	return utilversion.MustParseSemantic(serverVersion.GitVersion)
}

type containerRuntime struct {
	Name     string
	Endpoint string
}

var containerRuntimeInstance *containerRuntime

var (
	containerdSocks = []string{
		constants.DefaultContainerdSock,
		constants.K3sContainerdSock,
	}

	dockerSocks = []string{
		constants.DefaultContainerdSock,
		constants.CriDockerdSock,
		constants.DockershimSock,
		constants.DockerSock,
	}
)

// GetContainerRuntime 获取容器运行时以及容器运行时的 sock
// 1、从 Node 节点信息获取容器运行时，参考：kubectl get node -o wide，如果是 kubeadm 部署的集群，那么 Node 节点信息中会有容器运行时
// 2、如果 Node 节点信息中没有容器运行时，那么默认使用 containerd
// 3、如果容器运行时是 containerd，那么从 /etc/containerd/config.toml 中获取 sock，如果仍然没有，那么遍历常见的 containerd sock 路径
// 4、如果容器运行时是 docker，那么遍历常见的 docker sock 路径
func GetContainerRuntime() *containerRuntime {
	if containerRuntimeInstance != nil {
		return containerRuntimeInstance
	}

	containerRuntimeInstance = new(containerRuntime)

	nl, err := GetClientSet().CoreV1().Nodes().List(context.Background(), metav1.ListOptions{
		LabelSelector: constants.MasterNodeLabelKey,
	})
	if err == nil && len(nl.Items) > 0 {
		containerRuntimeInstance.Name = strings.Split(nl.Items[0].Status.NodeInfo.ContainerRuntimeVersion, ":")[0]
		// Note: Just worked for kubernetes installed by kubeadm
		if v, ok := nl.Items[0].Annotations[constants.KubeadmContainerRuntimeEndpointAnnoKey]; ok && v != "" {
			v = strings.TrimPrefix(v, "unix://")
			containerRuntimeInstance.Endpoint = v
		}
	}

	if containerRuntimeInstance.Name == "" {
		containerRuntimeInstance.Name = constants.ContainerRuntimeContainerd
	}

	if containerRuntimeInstance.Endpoint == "" {
		switch containerRuntimeInstance.Name {
		case constants.ContainerRuntimeContainerd:
			var endpoint string
			containerdConf, err := toml.LoadFile(constants.ContainerdConfigPath)
			if err == nil {
				endpoint = containerdConf.Get("grpc.address").(string)
			}

			if endpoint == "" {
				for _, sock := range containerdSocks {
					if _, err := os.Stat(sock); err == nil {
						if err == nil {
							endpoint = sock
							break
						}
					}
				}
			}
			containerRuntimeInstance.Endpoint = endpoint

		case constants.ContainerRuntimeDocker:
			for _, sock := range dockerSocks {
				_, err := os.Stat(sock)
				if err == nil {
					containerRuntimeInstance.Endpoint = sock
					break
				}
			}
			return containerRuntimeInstance
		}
	}

	if containerRuntimeInstance.Endpoint == "" {
		containerRuntimeInstance.Endpoint = constants.DefaultContainerdSock
	}

	return containerRuntimeInstance
}

// LabelsForAccessModeRWX returns wutong labels with access mode rwx.
func LabelsForAccessModeRWX() map[string]string {
	return map[string]string{
		"accessModes": "rwx",
	}
}

// LabelsForAccessModeRWO returns wutong labels with access mode rwo.
func LabelsForAccessModeRWO() map[string]string {
	return map[string]string{
		"accessModes": "rwo",
	}
}
