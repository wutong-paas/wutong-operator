package precheck

import (
	wutongv1alpha1 "github.com/wutong-paas/wutong-operator/v2/api/v1alpha1"
)

// PreChecker checks the environment and parameters required to install the wutong cluster
type PreChecker interface {
	Check() wutongv1alpha1.WutongClusterCondition
}
