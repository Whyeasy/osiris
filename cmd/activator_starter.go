package cmd

import (
	"context"

	"github.com/golang/glog"

	deployments "github.com/dailymotion-oss/osiris/pkg/deployments/starter"
	"github.com/dailymotion-oss/osiris/pkg/kubernetes"
	"github.com/dailymotion-oss/osiris/pkg/version"
)

func RunActivatorStarter(ctx context.Context) {
	glog.Infof(
		"Starting Osiris Activator Starter -- version %s -- commit %s",
		version.Version(),
		version.Commit(),
	)

	client, err := kubernetes.Client()
	if err != nil {
		glog.Fatalf("Error building kubernetes clientset: %s", err)
	}

	if err != nil {
		glog.Fatalf("Error retrieving activator configuration: %s", err)
	}

	activatorCfg, err := deployments.GetConfigFromEnvironment()
	if err != nil {
		glog.Fatalf(
			"Error retrieving endpoints controller configuration: %s",
			err,
		)
	}

	// Run the activator
	deployments.NewActivatorStarter(activatorCfg, client).Run(ctx)
}
