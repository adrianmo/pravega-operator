package k8sutil

import (
	"fmt"
	"github.com/pravega/pravega-operator/pkg/apis/pravega/v1alpha1"
)

func ConfigMapNameForBookie(clusterName string) string {
	return fmt.Sprintf("%s-bookie-config", clusterName)
}

func StatefulSetNameForBookie(clusterName string) string {
	return fmt.Sprintf("%s-bookie", clusterName)
}

func ConfigMapNameForController(clusterName string) string {
	return fmt.Sprintf("%s-controller-config", clusterName)
}

func ServiceNameForController(clusterName string) string {
	return fmt.Sprintf("%s-pravega-controller", clusterName)
}

func HeadlessServiceNameForBookie(clusterName string) string {
	return fmt.Sprintf("%s-bookie-headless", clusterName)
}

func DeploymentNameForController(clusterName string) string {
	return fmt.Sprintf("%s-pravega-controller", clusterName)
}

func ConfigMapNameForSegmentstore(clusterName string) string {
	return fmt.Sprintf("%s-segmentstore", clusterName)
}

func StatefulSetNameForSegmentstore(clusterName string) string {
	return fmt.Sprintf("%s-segmentstore", clusterName)
}

func LabelsForBookie(pravegaCluster *v1alpha1.PravegaCluster) map[string]string {
	return LabelsForPravegaCluster(pravegaCluster, "bookie")
}

func LabelsForController(pravegaCluster *v1alpha1.PravegaCluster) map[string]string {
	return LabelsForPravegaCluster(pravegaCluster, "pravega-controller")
}

func LabelsForSegmentStore(pravegaCluster *v1alpha1.PravegaCluster) map[string]string {
	return LabelsForPravegaCluster(pravegaCluster, "pravega-segmentstore")
}

func LabelsForPravegaCluster(pravegaCluster *v1alpha1.PravegaCluster, component string) map[string]string {
	return map[string]string{
		"app":             "pravega-cluster",
		"pravega_cluster": pravegaCluster.Name,
		"component":       component,
	}
}

func PravegaControllerServiceURL(pravegaCluster v1alpha1.PravegaCluster) string {
	return fmt.Sprintf("tcp://%v.%v:%v", ServiceNameForController(pravegaCluster.Name), pravegaCluster.Namespace, "9090")
}
