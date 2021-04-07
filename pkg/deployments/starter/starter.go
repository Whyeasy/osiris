package starter

import (
	"context"

	corev1 "k8s.io/api/core/v1"
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	"k8s.io/client-go/kubernetes"
	"k8s.io/client-go/tools/cache"

	"github.com/dailymotion-oss/osiris/pkg/healthz"
	k8s "github.com/dailymotion-oss/osiris/pkg/kubernetes"
	"github.com/golang/glog"
)

type ActivatorStarter interface {
	Run(ctx context.Context)
}

type activatorStarter struct {
	cfg                Config
	kubeClient         kubernetes.Interface
	namespacesInformer cache.SharedInformer
	ctx                context.Context
}

func NewActivatorStarter(config Config,
	kubeClient kubernetes.Interface) ActivatorStarter {
	as := &activatorStarter{
		cfg:        config,
		kubeClient: kubeClient,
		namespacesInformer: k8s.NamespaceIndexInformer(
			kubeClient,
			metav1.NamespaceAll),
	}
	as.namespacesInformer.AddEventHandler(cache.ResourceEventHandlerFuncs{
		AddFunc: as.syncNamespace,
		UpdateFunc: func(_, newObj interface{}) {
			as.syncNamespace(newObj)
		},
		DeleteFunc: as.syncDeletedNamespace,
	})
	return as
}

func (as *activatorStarter) Run(ctx context.Context) {
	ctx, cancel := context.WithCancel(ctx)
	defer cancel()
	as.ctx = ctx
	go func() {
		<-ctx.Done()
		glog.Infof("Activator Starter is shutting down")
	}()
	glog.Infof("Activator Starter is started")
	go func() {
		as.namespacesInformer.Run(ctx.Done())
		cancel()
	}()
	healthz.RunServer(ctx, 5000)
	cancel()
}

func (as *activatorStarter) syncNamespace(obj interface{}) {
	namespace := obj.(*corev1.Namespace)
	if k8s.NamespaceIsEligibleForActivatorDeployment(namespace.Annotations) {
		glog.Infof(
			"Notified about new or updated Osiris-enabled namespace %s",
			namespace.Name,
		)
		//Deploy activator to namespace
		glog.Infof(
			"Deploying activator in namespace %s",
			namespace.Name,
		)
		as.createActivator(namespace.Name)
	} else {
		glog.Infof(
			"Notified about new or updated Osiris-enabled namespace %s",
			namespace.Name,
		)
		// Remove activator in namespace.
		glog.Infof(
			"Removing activator in namespace %s",
			namespace.Name,
		)
		as.deleteActivator(namespace.Name)
	}
}

func (as *activatorStarter) syncDeletedNamespace(obj interface{}) {
	namesapce := obj.(*corev1.Namespace)
	if k8s.NamespaceIsEligibleForActivatorDeployment(namesapce.Annotations) {
		glog.Infof(
			"Notified about deleted Osiris-enabled namespace %s",
			namesapce.Name,
		)
	}
}