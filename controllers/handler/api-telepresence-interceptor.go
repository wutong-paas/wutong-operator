package handler

import (
	"context"

	wutongv1alpha1 "github.com/wutong-paas/wutong-operator/v2/api/v1alpha1"
	"github.com/wutong-paas/wutong-operator/v2/util/commonutil"
	"github.com/wutong-paas/wutong-operator/v2/util/constants"
	appsv1 "k8s.io/api/apps/v1"
	corev1 "k8s.io/api/core/v1"
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	"sigs.k8s.io/controller-runtime/pkg/client"
)

// APITelepresenceInterceptorName name
var APITelepresenceInterceptorName = "wt-api-telepresence-interceptor"

type apiTelepresenceInterceptor struct {
	ctx       context.Context
	client    client.Client
	labels    map[string]string
	component *wutongv1alpha1.WutongComponent
	cluster   *wutongv1alpha1.WutongCluster
}

var _ ComponentHandler = &api{}
var _ StorageClassRWXer = &api{}

// NewAPITelepresenceInterceptor new apiTelepresenceInterceptor handle
func NewAPITelepresenceInterceptor(ctx context.Context, client client.Client, component *wutongv1alpha1.WutongComponent, cluster *wutongv1alpha1.WutongCluster) ComponentHandler {
	return &apiTelepresenceInterceptor{
		ctx:       ctx,
		client:    client,
		component: component,
		cluster:   cluster,
		labels:    LabelsForWutongComponent(component),
	}
}

func (a *apiTelepresenceInterceptor) Before() error {
	return nil
}

func (a *apiTelepresenceInterceptor) Resources() []client.Object {
	resources := []client.Object{a.deployment()}
	return resources
}

func (a *apiTelepresenceInterceptor) After() error {
	return nil
}

func (a *apiTelepresenceInterceptor) ListPods() ([]corev1.Pod, error) {
	return listPods(a.ctx, a.client, a.component.Namespace, a.labels)
}

func (a *apiTelepresenceInterceptor) deployment() client.Object {
	volumeMounts := []corev1.VolumeMount{
		{
			Name:      "wt-management-cluster-kubeconfig",
			MountPath: "/root/.kube/config",
			SubPath:   "kubeconfig",
		},
	}
	volumes := []corev1.Volume{
		{
			Name: "wt-management-cluster-kubeconfig",
			VolumeSource: corev1.VolumeSource{
				Secret: &corev1.SecretVolumeSource{
					SecretName: "wt-management-cluster-kubeconfig",
				},
			},
		},
	}

	a.labels["name"] = APITelepresenceInterceptorName

	volumeMounts = mergeVolumeMounts(volumeMounts, a.component.Spec.VolumeMounts)
	volumes = mergeVolumes(volumes, a.component.Spec.Volumes)

	ds := &appsv1.Deployment{
		ObjectMeta: metav1.ObjectMeta{
			Name:      APITelepresenceInterceptorName,
			Namespace: a.component.Namespace,
			Labels:    a.labels,
		},
		Spec: appsv1.DeploymentSpec{
			Replicas: a.component.Spec.Replicas,
			Selector: &metav1.LabelSelector{
				MatchLabels: a.labels,
			},
			Template: corev1.PodTemplateSpec{
				ObjectMeta: metav1.ObjectMeta{
					Name:   APITelepresenceInterceptorName,
					Labels: a.labels,
				},
				Spec: corev1.PodSpec{
					ImagePullSecrets:              imagePullSecrets(a.component, a.cluster),
					TerminationGracePeriodSeconds: commonutil.Int64(120),
					Containers: []corev1.Container{
						{
							Name:  APITelepresenceInterceptorName,
							Image: a.component.Spec.Image,
							// ImagePullPolicy: a.component.ImagePullPolicy(),
							ImagePullPolicy: corev1.PullAlways,
							Env:             a.component.Spec.Env,
							Args:            a.component.Spec.Args,
							VolumeMounts:    volumeMounts,
							Resources:       a.component.Spec.Resources,
							SecurityContext: &corev1.SecurityContext{
								Capabilities: &corev1.Capabilities{
									Add: []corev1.Capability{
										"NET_ADMIN",
									},
								},
							},
							Lifecycle: &corev1.Lifecycle{
								PreStop: &corev1.LifecycleHandler{
									Exec: &corev1.ExecAction{
										Command: []string{"./leave-and-quit.sh"},
									},
								},
							},
							StartupProbe: &corev1.Probe{
								InitialDelaySeconds: 10,
								PeriodSeconds:       10,
								FailureThreshold:    10,
								ProbeHandler: corev1.ProbeHandler{
									Exec: &corev1.ExecAction{
										Command: []string{"./health-check.sh"},
									},
								},
							},
							LivenessProbe: &corev1.Probe{
								InitialDelaySeconds: 10,
								PeriodSeconds:       10,
								FailureThreshold:    3,
								ProbeHandler: corev1.ProbeHandler{
									Exec: &corev1.ExecAction{
										Command: []string{"./health-check.sh"},
									},
								},
							},
						},
					},
					PriorityClassName: constants.WutongPlatformComponentPriorityClassName,
					// ServiceAccountName: "wutong-operator",
					Volumes: volumes,
				},
			},
		},
	}

	return ds
}
