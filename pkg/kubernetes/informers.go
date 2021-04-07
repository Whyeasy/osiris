package kubernetes

import (
	"context"

	appsv1 "k8s.io/api/apps/v1"
	corev1 "k8s.io/api/core/v1"
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	"k8s.io/apimachinery/pkg/fields"
	"k8s.io/apimachinery/pkg/labels"
	"k8s.io/apimachinery/pkg/runtime"
	"k8s.io/apimachinery/pkg/watch"
	"k8s.io/client-go/kubernetes"
	"k8s.io/client-go/tools/cache"
)

func DeploymentsIndexInformer(
	client kubernetes.Interface,
	namespace string,
	fieldSelector fields.Selector,
	labelSelector labels.Selector,
) cache.SharedIndexInformer {
	deploymentsClient := client.AppsV1().Deployments(namespace)
	return cache.NewSharedIndexInformer(
		&cache.ListWatch{
			ListFunc: func(options metav1.ListOptions) (runtime.Object, error) {
				if fieldSelector != nil {
					options.FieldSelector = fieldSelector.String()
				}
				if labelSelector != nil {
					options.LabelSelector = labelSelector.String()
				}
				return deploymentsClient.List(context.TODO(), options)
			},
			WatchFunc: func(options metav1.ListOptions) (watch.Interface, error) {
				if fieldSelector != nil {
					options.FieldSelector = fieldSelector.String()
				}
				if labelSelector != nil {
					options.LabelSelector = labelSelector.String()
				}
				return deploymentsClient.Watch(context.TODO(), options)
			},
		},
		&appsv1.Deployment{},
		0,
		cache.Indexers{},
	)
}

func StatefulSetsIndexInformer(
	client kubernetes.Interface,
	namespace string,
	fieldSelector fields.Selector,
	labelSelector labels.Selector,
) cache.SharedIndexInformer {
	statefulSetsClient := client.AppsV1().StatefulSets(namespace)
	return cache.NewSharedIndexInformer(
		&cache.ListWatch{
			ListFunc: func(options metav1.ListOptions) (runtime.Object, error) {
				if fieldSelector != nil {
					options.FieldSelector = fieldSelector.String()
				}
				if labelSelector != nil {
					options.LabelSelector = labelSelector.String()
				}
				return statefulSetsClient.List(context.TODO(), options)
			},
			WatchFunc: func(options metav1.ListOptions) (watch.Interface, error) {
				if fieldSelector != nil {
					options.FieldSelector = fieldSelector.String()
				}
				if labelSelector != nil {
					options.LabelSelector = labelSelector.String()
				}
				return statefulSetsClient.Watch(context.TODO(), options)
			},
		},
		&appsv1.StatefulSet{},
		0,
		cache.Indexers{},
	)
}

func PodsIndexInformer(
	client kubernetes.Interface,
	namespace string,
	fieldSelector fields.Selector,
	labelSelector labels.Selector,
) cache.SharedIndexInformer {
	podsClient := client.CoreV1().Pods(namespace)
	return cache.NewSharedIndexInformer(
		&cache.ListWatch{
			ListFunc: func(options metav1.ListOptions) (runtime.Object, error) {
				if fieldSelector != nil {
					options.FieldSelector = fieldSelector.String()
				}
				if labelSelector != nil {
					options.LabelSelector = labelSelector.String()
				}
				return podsClient.List(context.TODO(), options)
			},
			WatchFunc: func(options metav1.ListOptions) (watch.Interface, error) {
				if fieldSelector != nil {
					options.FieldSelector = fieldSelector.String()
				}
				if labelSelector != nil {
					options.LabelSelector = labelSelector.String()
				}
				return podsClient.Watch(context.TODO(), options)
			},
		},
		&corev1.Pod{},
		0,
		cache.Indexers{},
	)
}

func ServicesIndexInformer(
	client kubernetes.Interface,
	namespace string,
	fieldSelector fields.Selector,
	labelSelector labels.Selector,
) cache.SharedIndexInformer {
	servicesClient := client.CoreV1().Services(namespace)
	return cache.NewSharedIndexInformer(
		&cache.ListWatch{
			ListFunc: func(options metav1.ListOptions) (runtime.Object, error) {
				if fieldSelector != nil {
					options.FieldSelector = fieldSelector.String()
				}
				if labelSelector != nil {
					options.LabelSelector = labelSelector.String()
				}
				return servicesClient.List(context.TODO(), options)
			},
			WatchFunc: func(options metav1.ListOptions) (watch.Interface, error) {
				if fieldSelector != nil {
					options.FieldSelector = fieldSelector.String()
				}
				if labelSelector != nil {
					options.LabelSelector = labelSelector.String()
				}
				return servicesClient.Watch(context.TODO(), options)
			},
		},
		&corev1.Service{},
		0,
		cache.Indexers{},
	)
}

func EndpointsIndexInformer(
	client kubernetes.Interface,
	namespace string,
	fieldSelector fields.Selector,
	labelSelector labels.Selector,
) cache.SharedIndexInformer {
	endpointsClient := client.CoreV1().Endpoints(namespace)
	return cache.NewSharedIndexInformer(
		&cache.ListWatch{
			ListFunc: func(options metav1.ListOptions) (runtime.Object, error) {
				if fieldSelector != nil {
					options.FieldSelector = fieldSelector.String()
				}
				if labelSelector != nil {
					options.LabelSelector = labelSelector.String()
				}
				return endpointsClient.List(context.TODO(), options)
			},
			WatchFunc: func(options metav1.ListOptions) (watch.Interface, error) {
				if fieldSelector != nil {
					options.FieldSelector = fieldSelector.String()
				}
				if labelSelector != nil {
					options.LabelSelector = labelSelector.String()
				}
				return endpointsClient.Watch(context.TODO(), options)
			},
		},
		&corev1.Endpoints{},
		0,
		cache.Indexers{},
	)
}

func NodesIndexInformer(
	client kubernetes.Interface,
	namespace string,
	fieldSelector fields.Selector,
	labelSelector labels.Selector,
) cache.SharedIndexInformer {
	nodesClient := client.CoreV1().Nodes()
	return cache.NewSharedIndexInformer(
		&cache.ListWatch{
			ListFunc: func(options metav1.ListOptions) (runtime.Object, error) {
				if fieldSelector != nil {
					options.FieldSelector = fieldSelector.String()
				}
				if labelSelector != nil {
					options.LabelSelector = labelSelector.String()
				}
				return nodesClient.List(context.TODO(), options)
			},
			WatchFunc: func(options metav1.ListOptions) (watch.Interface, error) {
				if fieldSelector != nil {
					options.FieldSelector = fieldSelector.String()
				}
				if labelSelector != nil {
					options.LabelSelector = labelSelector.String()
				}
				return nodesClient.Watch(context.TODO(), options)
			},
		},
		&corev1.Node{},
		0,
		cache.Indexers{},
	)
}

func NamespaceIndexInformer(
	client kubernetes.Interface,
	namespace string,
) cache.SharedIndexInformer {
	namespaceClient := client.CoreV1().Namespaces()
	return cache.NewSharedIndexInformer(
		&cache.ListWatch{
			ListFunc: func(options metav1.ListOptions) (runtime.Object, error) {
				return namespaceClient.List(context.TODO(), options)
			},
			WatchFunc: func(options metav1.ListOptions) (watch.Interface, error) {
				return namespaceClient.Watch(context.TODO(), options)
			},
		},
		&corev1.Namespace{},
		0,
		cache.Indexers{},
	)
}
