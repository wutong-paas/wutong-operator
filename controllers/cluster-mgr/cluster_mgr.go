package clustermgr

import (
	"context"
	"encoding/base64"
	"fmt"
	"time"

	"github.com/go-logr/logr"
	"github.com/pquerna/ffjson/ffjson"
	wutongv1alpha1 "github.com/wutong-paas/wutong-operator/api/v1alpha1"
	"github.com/wutong-paas/wutong-operator/controllers/cluster-mgr/precheck"
	"github.com/wutong-paas/wutong-operator/util/constants"
	"github.com/wutong-paas/wutong-operator/util/wtutil"
	corev1 "k8s.io/api/core/v1"
	storagev1 "k8s.io/api/storage/v1"
	k8sErrors "k8s.io/apimachinery/pkg/api/errors"
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	"k8s.io/apimachinery/pkg/runtime"
	"k8s.io/apimachinery/pkg/types"
	"k8s.io/client-go/util/retry"
	"sigs.k8s.io/controller-runtime/pkg/client"
	"sigs.k8s.io/controller-runtime/pkg/controller/controllerutil"
)

const (
	// WtHubCredentialsName name for wt-hub-credentials
	WtHubCredentialsName = "wt-hub-credentials"
)

var provisionerAccessModes = map[string]corev1.PersistentVolumeAccessMode{
	// Kubernetes Internal Provisioner.
	// More info: https://github.com/kubernetes/kubernetes/tree/v1.17.3/pkg/volume
	"kubernetes.io/aws-ebs":         corev1.ReadWriteOnce,
	"kubernetes.io/azure-disk":      corev1.ReadWriteOnce,
	"kubernetes.io/azure-file":      corev1.ReadWriteMany,
	"kubernetes.io/cephfs":          corev1.ReadWriteMany,
	"kubernetes.io/cinder":          corev1.ReadWriteOnce,
	"kubernetes.io/fc":              corev1.ReadWriteOnce,
	"kubernetes.io/flocker":         corev1.ReadWriteOnce,
	"kubernetes.io/gce-pd":          corev1.ReadWriteOnce,
	"kubernetes.io/glusterfs":       corev1.ReadWriteMany,
	"kubernetes.io/iscsi":           corev1.ReadWriteOnce,
	"kubernetes.io/nfs":             corev1.ReadWriteMany,
	"kubernetes.io/portworx-volume": corev1.ReadWriteMany,
	"kubernetes.io/quobyte":         corev1.ReadWriteMany,
	"kubernetes.io/wt":              corev1.ReadWriteMany,
	"kubernetes.io/scaleio":         corev1.ReadWriteMany,
	"kubernetes.io/storageos":       corev1.ReadWriteMany,
	// Alibaba csi plugins for kubernetes.
	// More info: https://github.com/kubernetes-sigs/alibaba-cloud-csi-driver/tree/master/pkg
	"cpfsplugin.csi.alibabacloud.com": corev1.ReadWriteMany,
	"diskplugin.csi.alibabacloud.com": corev1.ReadWriteOnce,
	"alicloud/disk":                   corev1.ReadWriteOnce,
	"lvmplugin.csi.alibabacloud.com":  corev1.ReadWriteMany,
	"memplugin.csi.alibabacloud.com":  corev1.ReadWriteMany,
	"nasplugin.csi.alibabacloud.com":  corev1.ReadWriteMany,
	"ossplugin.csi.alibabacloud.com":  corev1.ReadWriteMany,
}

// WutongClusteMgr -
type WutongClusteMgr struct {
	ctx    context.Context
	client client.Client
	scheme *runtime.Scheme
	log    logr.Logger

	cluster *wutongv1alpha1.WutongCluster
	sclist  []*wutongv1alpha1.StorageClass // storage class list
	defsc   string                         // default storage class
}

// NewClusterMgr new Cluster Mgr
func NewClusterMgr(ctx context.Context, client client.Client, log logr.Logger, cluster *wutongv1alpha1.WutongCluster, scheme *runtime.Scheme) *WutongClusteMgr {
	mgr := &WutongClusteMgr{
		ctx:     ctx,
		client:  client,
		log:     log,
		cluster: cluster,
		scheme:  scheme,
	}
	mgr.setStorageStorageClasses()
	return mgr
}

// setStorageStorageClasses set sclist and defsc for WutongClusteMgr
func (r *WutongClusteMgr) setStorageStorageClasses() {
	r.log.V(6).Info("start listing available storage classes")

	storageClassList := &storagev1.StorageClassList{}
	var opts []client.ListOption
	ctx, cancel := context.WithTimeout(r.ctx, time.Second*10)
	defer cancel()
	if err := r.client.List(ctx, storageClassList, opts...); err != nil {
		r.log.Error(err, "list storageclass")
		return
	}

	var storageClasses []*wutongv1alpha1.StorageClass
	for _, sc := range storageClassList.Items {
		v, ok := sc.Annotations[constants.DefaultStorageClassAnnotationKey]
		if ok && v == "true" {
			r.defsc = sc.Name
		}
		storageClass := &wutongv1alpha1.StorageClass{
			Name:        sc.Name,
			Provisioner: sc.Provisioner,
			AccessMode:  provisionerAccessModes[sc.Provisioner],
		}
		storageClasses = append(storageClasses, storageClass)
	}
	r.log.V(6).Info("listing available storage classes success")
	r.sclist = storageClasses
}

// CheckOrUpdateWutongCluster return update or not and error
func (r *WutongClusteMgr) CheckOrUpdateWutongCluster() (bool, error) {
	update := false
	// set default storage class for rwx
	if r.cluster.Spec.WutongVolumeSpecRWX == nil {
		r.cluster.Spec.WutongVolumeSpecRWX = &wutongv1alpha1.WutongVolumeSpec{
			StorageClassName: r.defsc,
		}
		update = true
	} else {
		if r.cluster.Spec.WutongVolumeSpecRWX.StorageClassName == "" {
			r.cluster.Spec.WutongVolumeSpecRWX.StorageClassName = r.defsc
			update = true
		}
	}

	if r.cluster.Spec.InstallVersion == "" {
		r.cluster.Spec.InstallVersion = constants.DefaultInstallVersion
		update = true
	}

	if update {
		if err := retry.RetryOnConflict(retry.DefaultRetry, func() error {
			rc := &wutongv1alpha1.WutongCluster{}
			if err := r.client.Get(r.ctx, types.NamespacedName{Name: r.cluster.Name, Namespace: r.cluster.Namespace}, rc); err != nil {
				return err
			}
			rc.Spec = r.cluster.Spec
			return r.client.Update(r.ctx, rc)
		}); err != nil {
			r.log.Error(err, "update wutongcluster status")
			return update, err
		}
	}

	return update, nil
}

// GenerateWutongClusterStatus creates the final WutongCluster status for a WutongCluster, given the
// internal WutongCluster status.
func (r *WutongClusteMgr) GenerateWutongClusterStatus() (*wutongv1alpha1.WutongClusterStatus, error) {
	r.log.V(6).Info("start generating status")

	s := &wutongv1alpha1.WutongClusterStatus{
		MasterRoleLabel: constants.MasterNodeLabelKey,
		StorageClasses:  r.sclist,
	}

	if r.checkIfImagePullSecretExists() {
		s.ImagePullSecret = &corev1.LocalObjectReference{Name: WtHubCredentialsName}
	}

	s.GatewayAvailableNodes = &wutongv1alpha1.AvailableNodes{
		SpecifiedNodes: r.listSpecifiedGatewayNodes(),
		MasterNodes:    wtutil.ListMasterNodesForGateway(),
	}
	s.ChaosAvailableNodes = &wutongv1alpha1.AvailableNodes{
		SpecifiedNodes: r.listSpecifiedChaosNodes(),
		MasterNodes:    wtutil.ListMasterNodes(),
	}

	// conditions for wutong cluster status
	s.Conditions = r.generateConditions()
	r.log.V(6).Info("generating status success")
	return s, nil
}

func (r *WutongClusteMgr) listSpecifiedGatewayNodes() []*wutongv1alpha1.K8sNode {
	nodes := wtutil.ListNodesByLabels(constants.SpecialGatewayLabelKey)
	// Filtering nodes with port conflicts
	// check gateway ports
	return wtutil.FilterNodesWithPortConflicts(nodes)
}

func (r *WutongClusteMgr) listSpecifiedChaosNodes() []*wutongv1alpha1.K8sNode {
	return wtutil.ListNodesByLabels(constants.SpecialChaosLabelKey)
}

// CreateImagePullSecret create image pull secret
func (r *WutongClusteMgr) CreateImagePullSecret() error {
	var secret corev1.Secret
	if err := r.client.Get(r.ctx, types.NamespacedName{Namespace: r.cluster.Namespace, Name: WtHubCredentialsName}, &secret); err != nil {
		if !k8sErrors.IsNotFound(err) {
			return err
		}
	}

	if config, exist := secret.Data[".dockerconfigjson"]; exist && string(config) == string(r.generateDockerConfig()) {
		r.log.V(5).Info("dockerconfig not change")
		return nil
	}
	secret = corev1.Secret{
		ObjectMeta: metav1.ObjectMeta{
			Name:      WtHubCredentialsName,
			Namespace: r.cluster.Namespace,
		},
		Data: map[string][]byte{
			".dockerconfigjson": r.generateDockerConfig(),
		},
		Type: corev1.SecretTypeDockerConfigJson,
	}

	if err := controllerutil.SetControllerReference(r.cluster, &secret, r.scheme); err != nil {
		return fmt.Errorf("set controller reference for secret %s: %v", WtHubCredentialsName, err)
	}

	err := r.client.Create(r.ctx, &secret)
	if err != nil {
		if k8sErrors.IsAlreadyExists(err) {
			r.log.V(7).Info("update image pull secret", "name", WtHubCredentialsName)
			err = r.client.Update(r.ctx, &secret)
			if err == nil {
				return nil
			}
		}
		return fmt.Errorf("create secret for pulling images: %v", err)
	}

	return nil
}

func (r *WutongClusteMgr) checkIfImagePullSecretExists() bool {
	secret := &corev1.Secret{}
	err := r.client.Get(r.ctx, types.NamespacedName{Namespace: r.cluster.Namespace, Name: WtHubCredentialsName}, secret)
	if err != nil {
		if !k8sErrors.IsNotFound(err) {
			r.log.Info(fmt.Sprintf("get secret %s: %v", WtHubCredentialsName, err))
		}
		return false
	}
	return true
}

func (r *WutongClusteMgr) generateDockerConfig() []byte {
	type dockerConfig struct {
		Auths map[string]map[string]string `json:"auths"`
	}

	username, password := r.cluster.Spec.ImageHub.Username, r.cluster.Spec.ImageHub.Password
	auth := map[string]string{
		"username": username,
		"password": password,
		"auth":     base64.StdEncoding.EncodeToString([]byte(username + ":" + password)),
	}

	dockercfg := dockerConfig{
		Auths: map[string]map[string]string{
			r.cluster.Spec.ImageHub.Domain: auth,
		},
	}

	bytes, _ := ffjson.Marshal(dockercfg)
	return bytes
}

func (r *WutongClusteMgr) generateConditions() []wutongv1alpha1.WutongClusterCondition {
	// region database
	spec := r.cluster.Spec
	if spec.RegionDatabase != nil && !r.isConditionTrue(wutongv1alpha1.WutongClusterConditionTypeDatabaseRegion) {
		preChecker := precheck.NewDatabasePrechecker(wutongv1alpha1.WutongClusterConditionTypeDatabaseRegion, spec.RegionDatabase)
		condition := preChecker.Check()
		r.cluster.Status.UpdateCondition(&condition)
	}

	// console database
	if spec.UIDatabase != nil && !r.isConditionTrue(wutongv1alpha1.WutongClusterConditionTypeDatabaseConsole) {
		preChecker := precheck.NewDatabasePrechecker(wutongv1alpha1.WutongClusterConditionTypeDatabaseConsole, spec.UIDatabase)
		condition := preChecker.Check()
		r.cluster.Status.UpdateCondition(&condition)
	}

	// image repository
	if spec.ImageHub != nil && !r.isConditionTrue(wutongv1alpha1.WutongClusterConditionTypeImageRepository) {
		preChecker := precheck.NewImageRepoPrechecker(r.ctx, r.log, r.cluster)
		condition := preChecker.Check()
		r.cluster.Status.UpdateCondition(&condition)
	}

	// kubernetes version
	if !r.isConditionTrue(wutongv1alpha1.WutongClusterConditionTypeKubernetesVersion) {
		k8sVersion := precheck.NewK8sVersionPrechecker(r.ctx, r.log, r.client)
		condition := k8sVersion.Check()
		r.cluster.Status.UpdateCondition(&condition)
	}

	storagePreChecker := precheck.NewStorage(r.ctx, r.client, r.cluster.GetNamespace(), r.cluster.Spec.WutongVolumeSpecRWX, r.defsc)
	storageCondition := storagePreChecker.Check()
	r.cluster.Status.UpdateCondition(&storageCondition)

	if r.cluster.Spec.InstallMode != wutongv1alpha1.InstallationModeOffline {
		dnsPrechecker := precheck.NewDNSPrechecker(r.cluster, r.log)
		dnsCondition := dnsPrechecker.Check()
		r.cluster.Status.UpdateCondition(&dnsCondition)
	}

	memory := precheck.NewMemory(r.ctx, r.log, r.client)
	memoryCondition := memory.Check()
	r.cluster.Status.UpdateCondition(&memoryCondition)

	// container network
	if r.cluster.Spec.SentinelImage != "" {
		containerNetworkPrechecker := precheck.NewContainerNetworkPrechecker(r.ctx, r.client, r.scheme, r.log, r.cluster)
		containerNetworkCondition := containerNetworkPrechecker.Check()
		r.cluster.Status.UpdateCondition(&containerNetworkCondition)
	}

	if idx, condition := r.cluster.Status.GetCondition(wutongv1alpha1.WutongClusterConditionTypeRunning); idx == -1 || condition.Status != corev1.ConditionTrue {
		running := r.runningCondition()
		r.cluster.Status.UpdateCondition(&running)
	}

	return r.cluster.Status.Conditions
}

func (r *WutongClusteMgr) isConditionTrue(typ3 wutongv1alpha1.WutongClusterConditionType) bool {

	_, condition := r.cluster.Status.GetCondition(typ3)

	if condition != nil && condition.Status == corev1.ConditionTrue {
		return true
	}
	return false
}

func (r *WutongClusteMgr) runningCondition() wutongv1alpha1.WutongClusterCondition {
	condition := wutongv1alpha1.WutongClusterCondition{
		Type:              wutongv1alpha1.WutongClusterConditionTypeRunning,
		Status:            corev1.ConditionTrue,
		LastHeartbeatTime: metav1.NewTime(time.Now()),
	}

	// list all WutongComponents
	WutongComponents, err := r.listWutongComponents()
	if err != nil {
		return FailCondition(condition, "ListWutongComponentFailed", err.Error())
	}

	for _, cpt := range WutongComponents {
		idx, c := cpt.Status.GetCondition(wutongv1alpha1.WutongComponentReady)
		if idx == -1 {
			return FailCondition(condition, "WutongComponentReadyNotFound",
				fmt.Sprintf("condition 'WutongComponentReady' not found for %s", cpt.GetName()))
		}
		if c.Status == corev1.ConditionFalse {
			return FailCondition(condition, "WutongComponentNotReady",
				fmt.Sprintf("WutongComponent(%s) not ready", cpt.GetName()))
		}
	}

	return condition
}

func (r *WutongClusteMgr) listWutongComponents() ([]wutongv1alpha1.WutongComponent, error) {
	WutongComponentList := &wutongv1alpha1.WutongComponentList{}
	err := r.client.List(r.ctx, WutongComponentList, client.InNamespace(r.cluster.Namespace))
	if err != nil {
		return nil, err
	}
	return WutongComponentList.Items, nil
}

// FailCondition -
func FailCondition(condition wutongv1alpha1.WutongClusterCondition, reason, msg string) wutongv1alpha1.WutongClusterCondition {
	condition.Status = corev1.ConditionFalse
	condition.Reason = reason
	condition.Message = msg
	return condition
}
