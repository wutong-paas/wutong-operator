package wtutil

import (
	"context"
	"fmt"
	"net"
	"path"
	"sort"
	"time"

	"github.com/wutong-paas/wutong-operator/util/constants"
	"github.com/wutong-paas/wutong-operator/util/k8sutil"

	wutongv1alpha1 "github.com/wutong-paas/wutong-operator/api/v1alpha1"
	corev1 "k8s.io/api/core/v1"
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	"k8s.io/apimachinery/pkg/labels"
	"k8s.io/apimachinery/pkg/selection"
)

type NodesSortByName []*wutongv1alpha1.K8sNode

func (s NodesSortByName) Len() int           { return len(s) }
func (s NodesSortByName) Swap(i, j int)      { s[i], s[j] = s[j], s[i] }
func (s NodesSortByName) Less(i, j int) bool { return s[i].Name < s[j].Name }

// GetImageRepository returns image repository name based on WutongCluster.
func GetImageRepository(cluster *wutongv1alpha1.WutongCluster) string {
	if cluster.Spec.ImageHub == nil {
		return constants.DefImageRepository
	}
	return path.Join(cluster.Spec.ImageHub.Domain, cluster.Spec.ImageHub.Namespace)
}

func GetImageRepositoryDomain(cluster *wutongv1alpha1.WutongCluster) string {
	if cluster.Spec.ImageHub == nil {
		return constants.DefImageRepository
	}
	return cluster.Spec.ImageHub.Domain
}

// LabelsForWutong returns labels for resources created by wutong operator.
func LabelsForWutong(labels map[string]string) map[string]string {
	wtLabels := map[string]string{
		"creator":  "Wutong",
		"belongTo": "wutong-operator",
	}
	for key, val := range labels {
		// wtLabels has priority over labels
		if wtLabels[key] != "" {
			continue
		}
		wtLabels[key] = val
	}
	return wtLabels
}

func ListMasterNodes() []*wutongv1alpha1.K8sNode {
	ctx, cancel := context.WithTimeout(context.Background(), time.Second*10)
	defer cancel()
	nodeList, err := k8sutil.GetClientSet().CoreV1().Nodes().List(ctx, metav1.ListOptions{})
	if err != nil {
		return nil
	}

	findIP := func(addresses []corev1.NodeAddress, addressType corev1.NodeAddressType) string {
		for _, address := range addresses {
			if address.Type == addressType {
				return address.Address
			}
		}
		return ""
	}

	var k8sNodes []*wutongv1alpha1.K8sNode
	for _, node := range nodeList.Items {
		_, ok := node.Labels[constants.MasterNodeLabelKey]
		if !ok {
			continue
		}
		k8sNode := &wutongv1alpha1.K8sNode{
			Name:       node.Name,
			InternalIP: findIP(node.Status.Addresses, corev1.NodeInternalIP),
			ExternalIP: findIP(node.Status.Addresses, corev1.NodeExternalIP),
		}
		k8sNodes = append(k8sNodes, k8sNode)
	}

	return k8sNodes
}

func ListNodesByLabels(labelKey string) []*wutongv1alpha1.K8sNode {
	nodeList := &corev1.NodeList{}
	ctx, cancel := context.WithTimeout(context.Background(), time.Second*10)
	defer cancel()

	re, err := labels.NewRequirement(labelKey, selection.Exists, []string{})
	if err == nil {
		if nodeList, err = k8sutil.GetClientSet().CoreV1().Nodes().List(ctx, metav1.ListOptions{
			LabelSelector: labels.NewSelector().Add(*re).String(),
		}); err != nil {
			return nil
		}
	} else {
		if nodeList, err = k8sutil.GetClientSet().CoreV1().Nodes().List(ctx, metav1.ListOptions{}); err != nil {
			return nil
		}
		for _, node := range nodeList.Items {
			if _, ok := node.Labels[labelKey]; ok {
				nodeList.Items = append(nodeList.Items, node)
			}
		}
	}

	findIP := func(addresses []corev1.NodeAddress, addressType corev1.NodeAddressType) string {
		for _, address := range addresses {
			if address.Type == addressType {
				return address.Address
			}
		}
		return ""
	}

	var k8sNodes []*wutongv1alpha1.K8sNode
	for _, node := range nodeList.Items {
		k8sNode := &wutongv1alpha1.K8sNode{
			Name:       node.Name,
			InternalIP: findIP(node.Status.Addresses, corev1.NodeInternalIP),
			ExternalIP: findIP(node.Status.Addresses, corev1.NodeExternalIP),
		}
		k8sNodes = append(k8sNodes, k8sNode)
	}

	sort.Sort(NodesSortByName(k8sNodes))

	return k8sNodes
}

func ListMasterNodesForGateway() []*wutongv1alpha1.K8sNode {
	nodes := ListMasterNodes()
	// Filtering nodes with port conflicts
	// check gateway ports
	return FilterNodesWithPortConflicts(nodes)
}

// FilterNodesWithPortConflicts -
func FilterNodesWithPortConflicts(nodes []*wutongv1alpha1.K8sNode) []*wutongv1alpha1.K8sNode {
	var result []*wutongv1alpha1.K8sNode
	gatewayPorts := []int{80, 443, 10254, 18080, 18081, 8443, 6060, 7070}
	for idx := range nodes {
		node := nodes[idx]
		ok := true
		for _, port := range gatewayPorts {
			if isPortOccupied(fmt.Sprintf("%s:%d", node.InternalIP, port)) {
				ok = false
				break
			}
		}
		if ok {
			result = append(result, node)
		}
	}
	return result
}

func isPortOccupied(address string) bool {
	conn, err := net.Dial("tcp", address)
	if err != nil {
		return false
	}
	defer func() { _ = conn.Close() }()
	return true
}

func GatewayIngressIP(cluster *wutongv1alpha1.WutongCluster) string {
	result := cluster.GatewayIngressIP()
	if result == "" {
		masterNodes := ListMasterNodes()
		if len(masterNodes) > 0 {
			result = masterNodes[0].ExternalIP
			if result == "" {
				result = masterNodes[0].InternalIP
			}
		}
	}
	return result
}

func TCPIngressAnnotationsFromPort(port string) map[string]string {
	return map[string]string{
		"nginx.ingress.kubernetes.io/l4-enable": "true",
		"nginx.ingress.kubernetes.io/l4-host":   "0.0.0.0",
		"nginx.ingress.kubernetes.io/l4-port":   port,
	}
}
