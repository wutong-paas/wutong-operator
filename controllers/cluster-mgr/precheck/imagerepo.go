package precheck

import (
	"context"
	"fmt"
	"time"

	"github.com/wutong-paas/wutong-operator/v2/util/constants"
	"github.com/wutong-paas/wutong-operator/v2/util/repositoryutil"
	"github.com/wutong-paas/wutong-operator/v2/util/wtutil"

	"github.com/go-logr/logr"
	wutongv1alpha1 "github.com/wutong-paas/wutong-operator/v2/api/v1alpha1"
	corev1 "k8s.io/api/core/v1"
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
)

type imagerepo struct {
	ctx     context.Context
	log     logr.Logger
	cluster *wutongv1alpha1.WutongCluster
}

// NewImageRepoPrechecker creates a new prechecker.
func NewImageRepoPrechecker(ctx context.Context, log logr.Logger, cluster *wutongv1alpha1.WutongCluster) PreChecker {
	l := log.WithName("ImageRepoPreChecker")
	return &imagerepo{
		ctx:     ctx,
		log:     l,
		cluster: cluster,
	}
}

func (d *imagerepo) Check() wutongv1alpha1.WutongClusterCondition {
	condition := wutongv1alpha1.WutongClusterCondition{
		Type:              wutongv1alpha1.WutongClusterConditionTypeImageRepository,
		Status:            corev1.ConditionTrue,
		LastHeartbeatTime: metav1.NewTime(time.Now()),
	}

	imageRepo := wtutil.GetImageRepository(d.cluster)

	if idx, cdt := d.cluster.Status.GetCondition(wutongv1alpha1.WutongClusterConditionTypeImageRepository); (idx == -1 || cdt.Reason == "DefaultImageRepoFailed") && imageRepo != constants.DefImageRepository {
		condition.Status = corev1.ConditionFalse
		condition.Reason = "InProgress"
		condition.Message =
			fmt.Sprintf("precheck for %s is in progress", wutongv1alpha1.WutongClusterConditionTypeImageRepository)
	}

	// Verify that the image repository is available
	d.log.V(6).Info("login repository", "repository", getImageRepositoryDomain(d.cluster), "user", d.cluster.Spec.ImageHub.Username)

	if err := repositoryutil.LoginRepository(getImageRepositoryDomain(d.cluster), d.cluster.Spec.ImageHub.Username, d.cluster.Spec.ImageHub.Password); err != nil {
		return d.failConditoin(condition, err)
	}
	return condition
}

func (d *imagerepo) failConditoin(condition wutongv1alpha1.WutongClusterCondition, err error) wutongv1alpha1.WutongClusterCondition {
	return failConditoin(condition, "ImageRepoFailed", err.Error())
}

// getImageRepositoryDomain returns image repository domain based on wutongcluster.
func getImageRepositoryDomain(cluster *wutongv1alpha1.WutongCluster) string {
	if cluster.Spec.ImageHub == nil {
		return constants.DefImageRepository
	}
	return cluster.Spec.ImageHub.Domain
}
