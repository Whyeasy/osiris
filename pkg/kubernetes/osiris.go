package kubernetes

import (
	"strconv"
	"strings"
)

const (
	IgnoredPathsAnnotationName         = "osiris.dm.gg/ignoredPaths"
	MetricsCollectorAnnotationName     = "osiris.dm.gg/metricsCollector"
	MetricsCheckIntervalAnnotationName = "osiris.dm.gg/metricsCheckInterval"
	enableScalingAnnotationName        = "osiris.dm.gg/enableScaling"
	collectMetricsAnnotationName       = "osiris.dm.gg/collectMetrics"
	manageEndpointsAnnotationName      = "osiris.dm.gg/manageEndpoints"
	namespaceActivatorAnnotationName   = "osiris.dm.gg/enableActivator"
)

// WorkloadIsEligibleForAutoScaling checks the annotations to see if the
// workload (deployment or statefulset) is eligible for auto-scaling with osiris or not.
func WorkloadIsEligibleForAutoScaling(annotations map[string]string) bool {
	return annotationBooleanValue(annotations, enableScalingAnnotationName)
}

// PodIsEligibleForProxyInjection checks the annotations to see if the
// pod is eligible for proxy injection or not.
func PodIsEligibleForProxyInjection(annotations map[string]string) bool {
	return annotationBooleanValue(annotations, collectMetricsAnnotationName)
}

// ServiceIsEligibleForEndpointsManagement checks the annotations to see if the
// service is eligible for management of its endpoints by osiris or not.
func ServiceIsEligibleForEndpointsManagement(annotations map[string]string) bool {
	return annotationBooleanValue(annotations, manageEndpointsAnnotationName)
}

func NamespaceIsEligibleForActivatorDeployment(annotations map[string]string) bool {
	return annotationBooleanValue(annotations, namespaceActivatorAnnotationName)
}

func annotationBooleanValue(annotations map[string]string, key string) bool {
	enabled, ok := annotations[key]
	if !ok {
		return false
	}
	switch strings.ToLower(enabled) {
	case "y", "yes", "true", "on", "1":
		return true
	default:
		return false
	}
}

// GetMinReplicas gets the minimum number of replicas required for scale up
// from the annotations. If it fails to do so, it returns the default value
// instead.
func GetMinReplicas(annotations map[string]string, defaultVal int32) int32 {
	val, ok := annotations["osiris.dm.gg/minReplicas"]
	if !ok {
		return defaultVal
	}
	minReplicas, err := strconv.Atoi(val)
	if err != nil {
		return defaultVal
	}
	return int32(minReplicas)
}
