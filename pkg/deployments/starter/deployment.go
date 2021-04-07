package starter

import (
	appsv1 "k8s.io/api/apps/v1"
	corev1 "k8s.io/api/core/v1"
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	"k8s.io/apimachinery/pkg/util/intstr"
)

func (as *activatorStarter) createActivator(namespace string) {

	labels := map[string]string{
		"app.kubernetes.io/name":      "osiris-activator",
		"app.kubernetes.io/component": "activator",
	}

	matchLabels := map[string]string{
		"app.kubernetes.io/name":      "osiris-activator",
		"app.kubernetes.io/component": "activator",
	}

	activatorDep := *&appsv1.Deployment{
		ObjectMeta: metav1.ObjectMeta{
			Name:      "osiris-activator",
			Namespace: namespace,
			Labels:    labels,
		},
		Spec: appsv1.DeploymentSpec{
			Selector: &metav1.LabelSelector{
				MatchLabels: matchLabels,
			},
			Template: corev1.PodTemplateSpec{
				ObjectMeta: metav1.ObjectMeta{
					Labels: matchLabels,
				},
				Spec: corev1.PodSpec{
					Containers: []corev1.Container{{
						Image: "ghcr.io/dailymotion-oss/osiris:test21",
						Name:  "activator",
						Env: []corev1.EnvVar{
							{
								Name: "POD_NAMESPACE",
								ValueFrom: &corev1.EnvVarSource{
									FieldRef: &corev1.ObjectFieldSelector{
										FieldPath: "metadata.namespace",
									},
								},
							},
						},
						Ports: []corev1.ContainerPort{
							{
								Name:          "proxy",
								ContainerPort: 5000,
								Protocol:      corev1.ProtocolTCP,
							},
							{
								Name:          "healthz",
								ContainerPort: 5001,
								Protocol:      corev1.ProtocolTCP,
							},
						},
						Command: []string{"/osiris/bin/osiris"},
						Args: []string{
							"--logtostderr=true",
							"activator",
						},
						LivenessProbe: &corev1.Probe{
							Handler: corev1.Handler{
								HTTPGet: &corev1.HTTPGetAction{
									Path: "/healthz",
									Port: intstr.FromString("healthz"),
								},
							},
						},
						ReadinessProbe: &corev1.Probe{
							Handler: corev1.Handler{
								HTTPGet: &corev1.HTTPGetAction{
									Path: "/healthz",
									Port: intstr.FromString("healthz"),
								},
							},
						},
					}},
				},
			},
		},
	}

	as.kubeClient.AppsV1().Deployments(namespace).Create(as.ctx, &activatorDep, metav1.CreateOptions{})
}

func (as *activatorStarter) deleteActivator(namespace string) {
	as.kubeClient.AppsV1().Deployments(namespace).Delete(as.ctx, "osiris-activator", metav1.DeleteOptions{})
}
