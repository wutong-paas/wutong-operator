package handler

import (
	wutongv1alpha1 "github.com/wutong-paas/wutong-operator/api/v1alpha1"
	"github.com/wutong-paas/wutong-operator/util/k8sutil"
)

func GatewayIngressIP(cluster *wutongv1alpha1.WutongCluster) string {
	result := cluster.GatewayIngressIP()
	if result == "" {
		masterNodes := k8sutil.ListMasterNodes()
		if len(masterNodes) > 0 {
			result = masterNodes[0].ExternalIP
			if result == "" {
				result = masterNodes[0].InternalIP
			}
		}
	}
	return result
}
