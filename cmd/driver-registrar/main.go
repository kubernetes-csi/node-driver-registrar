/*
Copyright 2017 The Kubernetes Authors.

Licensed under the Apache License, Version 2.0 (the "License");
you may not use this file except in compliance with the License.
You may obtain a copy of the License at

    http://www.apache.org/licenses/LICENSE-2.0

Unless required by applicable law or agreed to in writing, software
distributed under the License is distributed on an "AS IS" BASIS,
WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
See the License for the specific language governing permissions and
limitations under the License.
*/

package main

import (
	"flag"
	"fmt"
	"os"

	"github.com/golang/glog"
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	"k8s.io/client-go/kubernetes"
	"k8s.io/client-go/rest"
	"k8s.io/client-go/tools/clientcmd"
	"k8s.io/client-go/util/retry"
)

const (
	// Name of node label that contains JSON map of driver names to node names
	labelKey = "csi.volume.kubernetes.io/nodeid"
)

// Command line flags
var (
	kubeconfig = flag.String("kubeconfig", "", "Absolute path to the kubeconfig file. Required only when running out of cluster.")
)

func main() {
	flag.Set("logtostderr", "true")
	flag.Parse()

	// Fetch node name from environemnt label
	nodeName := os.Getenv("KUBE_NODE_NAME")
	if nodeName == "" {
		glog.Error(fmt.Errorf("Node name not found. The environment variable KUBE_NODE_NAME is empty."))
		os.Exit(1)
	}

	// Create the client config. Use kubeconfig if given, otherwise assume
	// in-cluster.
	config, err := buildConfig(*kubeconfig)
	if err != nil {
		glog.Error(err.Error())
		os.Exit(1)
	}

	clientset, err := kubernetes.NewForConfig(config)
	if err != nil {
		glog.Error(err.Error())
		os.Exit(1)
	}

	nodesClient := clientset.CoreV1().Nodes()
	// nodeObj, err := nodesClient.Get(nodeName, metav1.GetOptions{})
	// if errors.IsNotFound(err) {
	// 	glog.Errorf("Node %q not found\n", nodeName)
	// 	os.Exit(1)
	// } else if statusError, isStatus := err.(*errors.StatusError); isStatus {
	// 	glog.Errorf("Error getting node %q %v\n", nodeName, statusError.ErrStatus.Message)
	// 	os.Exit(1)
	// } else if err != nil {
	// 	glog.Error(err.Error())
	// 	os.Exit(1)
	// } else {
	// 	fmt.Printf("Found node:\n%v\n", nodeObj)
	// }

	retryErr := retry.RetryOnConflict(retry.DefaultRetry, func() error {
		// Retrieve the latest version of Node before attempting update
		// RetryOnConflict uses exponential backoff to avoid exhausting the apiserver
		result, getErr := nodesClient.Get(nodeName, metav1.GetOptions{})
		if getErr != nil {
			glog.Error("Failed to get latest version of Node: %v", getErr)
			os.Exit(1)
		}

		var previousLabelValue string
		if result.ObjectMeta.Labels != nil {
			previousLabelValue := result.ObjectMeta.Labels[labelKey]
			glog.V(3).Infof("previousLabelValue=%q", previousLabelValue)
		}

		result.ObjectMeta.Labels = cloneAndAddLabel(
			result.ObjectMeta.Labels, labelKey, previousLabelValue+"newValue")
		_, updateErr := nodesClient.Update(result)
		return updateErr
	})
	if retryErr != nil {
		panic(fmt.Errorf("Update failed: %v", retryErr))
	}
	fmt.Println("Updated node...")

}

func buildConfig(kubeconfig string) (*rest.Config, error) {
	if kubeconfig != "" {
		return clientcmd.BuildConfigFromFlags("", kubeconfig)
	}

	// Return config object which uses the service account kubernetes gives to
	// pods. It's intended for clients that are running inside a pod running on
	// kubernetes.
	return rest.InClusterConfig()
}

// Clones the given map and returns a new map with the given key and value added.
// Returns the given map, if labelKey is empty.
func cloneAndAddLabel(
	labels map[string]string, labelKey, labelValue string) map[string]string {
	if labelKey == "" {
		// Don't need to add a label.
		return labels
	}
	// Clone.
	newLabels := map[string]string{}
	for key, value := range labels {
		newLabels[key] = value
	}
	newLabels[labelKey] = labelValue
	return newLabels
}
