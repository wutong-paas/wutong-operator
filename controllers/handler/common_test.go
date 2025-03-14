package handler

import (
	"context"
	"testing"

	"github.com/stretchr/testify/assert"
	wutongv1alpha1 "github.com/wutong-paas/wutong-operator/v2/api/v1alpha1"
	"github.com/wutong-paas/wutong-operator/v2/util/k8sutil"
	corev1 "k8s.io/api/core/v1"
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	"k8s.io/apimachinery/pkg/runtime"
	"sigs.k8s.io/controller-runtime/pkg/client/fake"
)

func TestGetDefaultInfo(t *testing.T) {
	ctx := context.Background()

	scheme := runtime.NewScheme()
	if err := corev1.AddToScheme(scheme); err != nil {
		t.Fatal(err)
	}

	namespace := "wt-system"
	secret := &corev1.Secret{
		ObjectMeta: metav1.ObjectMeta{
			Name:      DBName,
			Namespace: "wt-system",
		},
		Data: map[string][]byte{
			mysqlPasswordKey: []byte("foobar"),
			mysqlUserKey:     []byte("write"),
		},
	}
	clientset := fake.NewClientBuilder().WithScheme(scheme).WithRuntimeObjects(secret).Build()

	dbInfo, err := getDefaultDBInfo(ctx, clientset, nil, namespace, DBName)
	if err != nil {
		t.Errorf("get db info: %v", err)
		t.FailNow()
	}
	assert.NotNil(t, dbInfo)
	assert.Equal(t, "foobar", dbInfo.Password)
	assert.Equal(t, "write", dbInfo.Username)
}

func TestStorageClassRWXVolumeNotFound(t *testing.T) {
	scheme := runtime.NewScheme()
	if err := wutongv1alpha1.AddToScheme(scheme); err != nil {
		t.Fatal(err)
	}

	cli := fake.NewClientBuilder().WithScheme(scheme).Build()
	ctx := context.Background()
	ns := "wt-system"
	_, err := storageClassNameFromWutongVolumeRWX(ctx, cli, ns)
	assert.NotNil(t, err)
	assert.True(t, IsWutongVolumeNotFound(err))
	assert.Equal(t, WutongVolumeNotFound, err.Error())
}

func TestStorageClassRWXVolumeRWXNotReady(t *testing.T) {
	scheme := runtime.NewScheme()
	if err := wutongv1alpha1.AddToScheme(scheme); err != nil {
		t.Fatal(err)
	}

	ns := "wt-system"
	labels := k8sutil.LabelsForAccessModeRWX()
	volume := &wutongv1alpha1.WutongVolume{
		ObjectMeta: metav1.ObjectMeta{
			Namespace: ns,
			Labels:    labels,
		},
	}
	cli := fake.NewClientBuilder().WithScheme(scheme).WithRuntimeObjects(volume).Build()
	ctx := context.Background()
	_, err := storageClassNameFromWutongVolumeRWX(ctx, cli, ns)
	assert.NotNil(t, err)
	assert.True(t, IsIgnoreError(err))
	assert.Equal(t, "storage class not ready", err.Error())
}

func TestStorageClassRWXVolumeRWXOK(t *testing.T) {
	scheme := runtime.NewScheme()
	if err := wutongv1alpha1.AddToScheme(scheme); err != nil {
		t.Fatal(err)
	}

	ns := "wt-system"
	labels := k8sutil.LabelsForAccessModeRWX()
	sc := "foobar.csi.wutong.io"
	volume := &wutongv1alpha1.WutongVolume{
		ObjectMeta: metav1.ObjectMeta{
			Namespace: ns,
			Labels:    labels,
		},
		Spec: wutongv1alpha1.WutongVolumeSpec{
			StorageClassName: sc,
		},
	}
	cli := fake.NewClientBuilder().WithScheme(scheme).WithRuntimeObjects(volume).Build()
	ctx := context.Background()
	got, err := storageClassNameFromWutongVolumeRWX(ctx, cli, ns)
	assert.Nil(t, err)
	assert.Equal(t, sc, got.storageClassName)
}

func TestStorageClassRWXVolumeRWONotFoundAndRWXNotFound(t *testing.T) {
	scheme := runtime.NewScheme()
	if err := wutongv1alpha1.AddToScheme(scheme); err != nil {
		t.Fatal(err)
	}

	ns := "wt-system"
	cli := fake.NewClientBuilder().WithScheme(scheme).Build()
	ctx := context.Background()
	_, err := storageClassNameFromWutongVolumeRWO(ctx, cli, ns)
	assert.NotNil(t, err)
	assert.True(t, IsWutongVolumeNotFound(err))
	assert.Equal(t, WutongVolumeNotFound, err.Error())
}

func TestStorageClassRWXVolumeRWONotFoundButRWXFound(t *testing.T) {
	scheme := runtime.NewScheme()
	if err := wutongv1alpha1.AddToScheme(scheme); err != nil {
		t.Fatal(err)
	}

	ns := "wt-system"
	labels := k8sutil.LabelsForAccessModeRWX()
	sc := "foobar.csi.wutong.io"
	volume := &wutongv1alpha1.WutongVolume{
		ObjectMeta: metav1.ObjectMeta{
			Namespace: ns,
			Labels:    labels,
		},
		Spec: wutongv1alpha1.WutongVolumeSpec{
			StorageClassName: sc,
		},
	}
	cli := fake.NewClientBuilder().WithScheme(scheme).WithRuntimeObjects(volume).Build()
	ctx := context.Background()
	got, err := storageClassNameFromWutongVolumeRWO(ctx, cli, ns)
	assert.Nil(t, err)
	assert.Equal(t, sc, got.storageClassName)
}

func TestStorageClassRWXVolumeRWOOK(t *testing.T) {
	scheme := runtime.NewScheme()
	if err := wutongv1alpha1.AddToScheme(scheme); err != nil {
		t.Fatal(err)
	}

	ns := "wt-system"

	volumerwo := getVolume(ns, k8sutil.LabelsForAccessModeRWO())
	volumerwx := getVolume(ns, k8sutil.LabelsForAccessModeRWX())

	cli := fake.NewClientBuilder().WithScheme(scheme).WithRuntimeObjects(volumerwo, volumerwx).Build()
	ctx := context.Background()
	got, err := storageClassNameFromWutongVolumeRWO(ctx, cli, ns)
	assert.Nil(t, err)
	assert.Equal(t, volumerwo.Spec.StorageClassName, got.storageClassName)
}

func TestSetStorageCassNameRWX(t *testing.T) {
	scheme := runtime.NewScheme()
	if err := wutongv1alpha1.AddToScheme(scheme); err != nil {
		t.Fatal(err)
	}

	ns := "wt-system"

	volumerwx := getVolume(ns, k8sutil.LabelsForAccessModeRWX())
	cli := fake.NewClientBuilder().WithScheme(scheme).WithRuntimeObjects(volumerwx).Build()
	ctx := context.Background()

	dummyStorageClassRWX := &dummyStorageClassRWX{}
	err := setStorageCassName(ctx, cli, ns, dummyStorageClassRWX)
	assert.Nil(t, err)
	assert.Equal(t, volumerwx.Spec.StorageClassName, dummyStorageClassRWX.pvcParametersRWX.storageClassName)
}

func TestSetStorageCassNameRWO(t *testing.T) {
	scheme := runtime.NewScheme()
	if err := wutongv1alpha1.AddToScheme(scheme); err != nil {
		t.Fatal(err)
	}

	ns := "wt-system"

	volumerwo := getVolume(ns, k8sutil.LabelsForAccessModeRWO())
	cli := fake.NewClientBuilder().WithScheme(scheme).WithRuntimeObjects(volumerwo).Build()
	ctx := context.Background()

	dummyStorageClassRWO := &dummyStorageClassRWO{}
	err := setStorageCassName(ctx, cli, ns, dummyStorageClassRWO)
	assert.Nil(t, err)
	assert.Equal(t, volumerwo.Spec.StorageClassName, dummyStorageClassRWO.pvcParametersRWO.storageClassName)
}

func TestSetStorageCassNameBothRWXAndRWO(t *testing.T) {
	scheme := runtime.NewScheme()
	if err := wutongv1alpha1.AddToScheme(scheme); err != nil {
		t.Fatal(err)
	}

	ns := "wt-system"

	volumerwx := getVolume(ns, k8sutil.LabelsForAccessModeRWX())
	volumerwo := getVolume(ns, k8sutil.LabelsForAccessModeRWO())
	cli := fake.NewClientBuilder().WithScheme(scheme).WithRuntimeObjects(volumerwo, volumerwx).Build()
	ctx := context.Background()

	dummyStorageClass := &dummyStorageClass{}
	err := setStorageCassName(ctx, cli, ns, dummyStorageClass)
	assert.Nil(t, err)
	assert.Equal(t, volumerwo.Spec.StorageClassName, dummyStorageClass.pvcParametersRWO.storageClassName)
	assert.Equal(t, volumerwx.Spec.StorageClassName, dummyStorageClass.pvcParametersRWX.storageClassName)
}

func getVolume(ns string, labels map[string]string) *wutongv1alpha1.WutongVolume {
	sc := "foo" + labels["accessModes"] + ".csi.wutong.io"
	volume := &wutongv1alpha1.WutongVolume{
		ObjectMeta: metav1.ObjectMeta{
			Name:      labels["accessModes"],
			Namespace: ns,
			Labels:    labels,
		},
		Spec: wutongv1alpha1.WutongVolumeSpec{
			StorageClassName: sc,
		},
	}
	return volume
}

type dummyStorageClassRWX struct {
	pvcParametersRWX *pvcParameters
}

var _ StorageClassRWXer = &dummyStorageClassRWX{}

func (d *dummyStorageClassRWX) SetStorageClassNameRWX(pvcParameters *pvcParameters) {
	d.pvcParametersRWX = pvcParameters
}

type dummyStorageClassRWO struct {
	pvcParametersRWO *pvcParameters
}

var _ StorageClassRWOer = &dummyStorageClassRWO{}

func (d *dummyStorageClassRWO) SetStorageClassNameRWO(pvcParameters *pvcParameters) {
	d.pvcParametersRWO = pvcParameters
}

type dummyStorageClass struct {
	dummyStorageClassRWX
	dummyStorageClassRWO
}

func TestMergeArgs(t *testing.T) {
	commonArgs := []string{
		"--storage.tsdb.no-lockfile",
	}
	priorityArgs := []string{
		"--log-level=debug",
	}
	expectArgs := []string{
		"--storage.tsdb.no-lockfile",
		"--log-level=debug",
	}
	args := mergeArgs(commonArgs, priorityArgs)
	assert.ElementsMatch(t, expectArgs, args)
}
