package precheck

import (
	"context"
	"fmt"
	"time"

	wutongv1alpha1 "github.com/wutong-paas/wutong-operator/v2/api/v1alpha1"
	corev1 "k8s.io/api/core/v1"
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	"sigs.k8s.io/controller-runtime/pkg/client"
)

type storage struct {
	ctx                     context.Context
	client                  client.Client
	ns                      string
	rwx                     *wutongv1alpha1.WutongVolumeSpec
	defaultStorageClassName string
}

// NewStorage -
func NewStorage(ctx context.Context, client client.Client, ns string, rwx *wutongv1alpha1.WutongVolumeSpec, defaultStorageClassName string) PreChecker {
	return &storage{
		ctx:                     ctx,
		client:                  client,
		ns:                      ns,
		rwx:                     rwx,
		defaultStorageClassName: defaultStorageClassName,
	}
}

func (s *storage) Check() wutongv1alpha1.WutongClusterCondition {
	condition := wutongv1alpha1.WutongClusterCondition{
		Type:              wutongv1alpha1.WutongClusterConditionTypeStorage,
		Status:            corev1.ConditionTrue,
		LastHeartbeatTime: metav1.NewTime(time.Now()),
	}

	if s.rwx.StorageClassName == "" {
		condition.Status = corev1.ConditionFalse
		condition.Reason = "InProgress"
		condition.Message = fmt.Sprintf("precheck for %s is in progress", wutongv1alpha1.WutongClusterConditionTypeStorage)
	}

	return condition
}

// func (s *storage) isPVCBound(pvc *corev1.PersistentVolumeClaim) bool {
// 	return pvc.Status.Phase == corev1.ClaimBound
// }

// func (s *storage) pvcForWTData(accessModes []corev1.PersistentVolumeAccessMode, storageClassName string) *corev1.PersistentVolumeClaim {
// 	labels := wtutil.LabelsForWutong(nil)
// 	return k8sutil.PersistentVolumeClaimForWTData(s.ns, constants.WTDataPVC, accessModes, labels, storageClassName, 20)
// }

// func (s *storage) failConditoin(condition wutongv1alpha1.WutongClusterCondition, msg string) wutongv1alpha1.WutongClusterCondition {
// 	return failConditoin(condition, "StorageFailed", msg)
// }

// func eventListToString(eventList *corev1.EventList) string {
// 	var res []string
// 	for _, event := range eventList.Items {
// 		res = append(res, fmt.Sprintf("%s: %s", event.Reason, event.Message))
// 	}
// 	return strings.Join(res, ",")
// }
