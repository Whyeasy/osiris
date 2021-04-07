package main

import (
	"flag"
	"strings"
	"time"

	"github.com/golang/glog"

	"github.com/dailymotion-oss/osiris/cmd"
	"github.com/dailymotion-oss/osiris/pkg/signals"
)

func main() {
	const usageMsg = `usage: must specify Osiris component to start using ` +
		`argument "activator", "endpoints-controller", "endpoints-hijacker", ` +
		`"proxy", "proxy-injector", or "zeroscaler"`

	// We need to parse flags for glog-related options to take effect
	flag.Parse()

	if len(flag.Args()) != 1 {
		glog.Fatal(usageMsg)
	}

	// This context will automatically be canceled on SIGINT or SIGTERM.
	ctx := signals.Context()

	switch strings.ToLower(flag.Arg(0)) {
	case "activator":
		cmd.RunActivator(ctx)
	case "activator-starter":
		cmd.RunActivatorStarter(ctx)
	case "endpoints-controller":
		cmd.RunEndpointsController(ctx)
	case "endpoints-hijacker":
		cmd.RunEndpointsHijacker(ctx)
	case "proxy":
		cmd.RunProxy(ctx)
	case "proxy-injector":
		cmd.RunProxyInjector(ctx)
	case "zeroscaler":
		cmd.RunZeroScaler(ctx)
	default:
		glog.Fatal(usageMsg)
	}

	// A short grace period
	shutdownDuration := 5 * time.Second
	glog.Infof("allowing %s for graceful shutdown to complete", shutdownDuration)
	<-time.After(shutdownDuration)
}
