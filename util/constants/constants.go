package constants

// Keys
const (
	DefaultStorageClassAnnotationKey       = "storageclass.kubernetes.io/is-default-class"
	KubeadmContainerRuntimeEndpointAnnoKey = "kubeadm.alpha.kubernetes.io/cri-socket"
	MasterNodeLabelKey                     = "node-role.kubernetes.io/control-plane"
)

// Container Runtime
const (
	// ContainerRuntimeDocker docker runtime
	ContainerRuntimeDocker = "docker"
	// ContainerRuntimeContainerd containerd runtime
	ContainerRuntimeContainerd = "containerd"

	ContainerdConfigPath   = "/etc/containerd/config.toml"
	K3sContainerConfigPath = "/var/lib/rancher/k3s/agent/etc/containerd/config.toml"

	DefaultContainerdSock = "/run/containerd/containerd.sock"
	K3sContainerdSock     = "/run/k3s/containerd/containerd.sock"
	CriDockerdSock        = "/run/cri-dockerd.sock"
	DockershimSock        = "/run/dockershim.sock"
	DockerSock            = "/run/docker.sock"
	CrioSock              = "/run/crio/crio.sock"

	DefaultContainerdCertsDir = "/etc/containerd/certs.d"
	DefaultDockerCertsDir     = "/etc/docker/certs.d"
)

// Wutong
const (
	// WutongSystemNamespace wt-system
	WutongSystemNamespace = "wt-system"
	// DefInstallPkgDestPath  Default destination path of the installation package extraction.
	DefInstallPkgDestPath = "/tmp/DefInstallPkgDestPath"
	// WutongClusterName wutong cluster resource name
	WutongClusterName = "wutongcluster"
	// WutongPackageName wutong package resource name
	WutongPackageName = "wutongpackage"
	// DefImageRepository is the default domain name of the mirror repository that Wutong is installed.
	DefImageRepository = "wutong.me"
	// WTDataPVC -
	WTDataPVC = "wt-cpt-wtdata"
	// CachePVC -
	CachePVC = "wt-chaos-cache"
	// FoobarPVC -
	// FoobarPVC = "foobar"
	// SpecialGatewayLabelKey is a special node label, used to specify where to install the wt-gateway
	SpecialGatewayLabelKey = "wutong.io/gateway"
	// SpecialChaosLabelKey is a special node label, used to specify where to install the wt-chaos
	SpecialChaosLabelKey = "wutong.io/chaos"
	// DefHTTPDomainSuffix -
	DefHTTPDomainSuffix = "wtapps.cn"

	// AliyunCSIDiskPlugin name for aliyun csi disk plugin
	AliyunCSIDiskPlugin = "aliyun-csi-disk-plugin"
	// AliyunCSIDiskProvisioner name for aliyun csi disk provisioner
	AliyunCSIDiskProvisioner = "aliyun-csi-disk-provisioner"
	// AliyunCSINasPlugin name for aliyun csi nas plugin
	AliyunCSINasPlugin = "aliyun-csi-nas-plugin"
	// AliyunCSINasProvisioner name for aliyun csi nas provisioner
	AliyunCSINasProvisioner = "aliyun-csi-nas-provisioner"

	// ServiceAccountName is the name of service account
	ServiceAccountName = "wutong-operator"

	// WutongClusterSettingsConfigMapName is the name of wutong cluster settings configmap
	WutongClusterSettingsConfigMapName      = "wutong-cluster-settings"
	WutongClusterCurrentInstalledVersionKey = "CurrentWutongClusterInstalledVersion"
	WutongClusterEdgeIsolatedClusterCodeKey = "EdgeIsolatedClusterCode"

	// InstallImageRepo install image repo
	InstallImageRepo = "swr.cn-southwest-2.myhuaweicloud.com/wutong"

	// DefaultInstallVersion default install version
	DefaultInstallVersion = "v1.16.1"

	WutongPlatformComponentPriorityClassName = "wutong-platform-component"
)
