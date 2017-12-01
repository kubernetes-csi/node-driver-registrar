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
	"context"
	"encoding/json"
	"flag"
	"fmt"
	"os"
	"time"

	"github.com/golang/glog"
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	"k8s.io/client-go/kubernetes"
	"k8s.io/client-go/rest"
	"k8s.io/client-go/tools/clientcmd"
	"k8s.io/client-go/util/retry"

	"github.com/kubernetes-csi/driver-registrar/pkg/connection"
)

const (
	// Name of node annotation that contains JSON map of driver names to node
	// names
	annotationKey = "csi.volume.kubernetes.io/nodeid"

	// Default timeout of short CSI calls like GetPluginInfo
	csiTimeout = time.Second
)

// Command line flags
var (
	kubeconfig        = flag.String("kubeconfig", "", "Absolute path to the kubeconfig file. Required only when running out of cluster.")
	connectionTimeout = flag.Duration("connection-timeout", 1*time.Minute, "Timeout for waiting for CSI driver socket.")
	csiAddress        = flag.String("csi-address", "/run/csi/socket", "Address of the CSI driver socket.")
)

func main() {
	flag.Set("logtostderr", "true")
	flag.Parse()

	// Fetch node name from environemnt annotation
	nodeName := os.Getenv("KUBE_NODE_NAME")
	if nodeName == "" {
		glog.Error(fmt.Errorf(
			"Node name not found. The environment variable KUBE_NODE_NAME is empty."))
		os.Exit(1)
	}

	// Once https://github.com/container-storage-interface/spec/issues/159 is
	// resolved, if plugin does not support PUBLISH_UNPUBLISH_VOLUME, then we
	// can skip adding mappting to "csi.volume.kubernetes.io/nodeid" annotation.

	// Connect to CSI.
	glog.V(1).Infof("Attempting to open a gRPC connection with: %q", csiAddress)
	csiConn, err := connection.NewConnection(*csiAddress, *connectionTimeout)
	if err != nil {
		glog.Error(err.Error())
		os.Exit(1)
	}

	// Get CSI driver name.
	glog.V(1).Infof("Calling CSI driver to discover driver name.")
	ctx, cancel := context.WithTimeout(context.Background(), csiTimeout)
	defer cancel()
	driverName, err := csiConn.GetDriverName(ctx)
	if err != nil {
		glog.Error(err.Error())
		os.Exit(1)
	}
	glog.V(2).Infof("CSI driver name: %q", driverName)

	// Get CSI Driver Node ID
	glog.V(1).Infof("Calling CSI driver to discover node ID.")
	ctx, cancel = context.WithTimeout(context.Background(), csiTimeout)
	defer cancel()
	csiDriverNodeId, err := csiConn.GetNodeID(ctx)
	if err != nil {
		glog.Error(err.Error())
		os.Exit(1)
	}
	glog.V(2).Infof("CSI driver node ID: %q", csiDriverNodeId)

	// Create the client config. Use kubeconfig if given, otherwise assume
	// in-cluster.
	glog.V(1).Infof("Loading kubeconfig.")
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

	glog.V(1).Infof("Attempt to update node annotation if needed")
	nodesClient := clientset.CoreV1().Nodes()

	// Add or update annotation on Node object
	retryErr := retry.RetryOnConflict(retry.DefaultRetry, func() error {
		// Retrieve the latest version of Node before attempting update, so that
		// existing changes are not overwritten. RetryOnConflict uses
		// exponential backoff to avoid exhausting the apiserver.
		result, getErr := nodesClient.Get(nodeName, metav1.GetOptions{})
		if getErr != nil {
			glog.Error("Failed to get latest version of Node: %v", getErr)
			os.Exit(1)
		}

		var previousAnnotationValue string
		if result.ObjectMeta.Annotations != nil {
			previousAnnotationValue =
				result.ObjectMeta.Annotations[annotationKey]
			glog.V(3).Infof(
				"previousAnnotationValue=%q", previousAnnotationValue)
		}

		existingDriverMap := map[string]string{}
		if previousAnnotationValue != "" {
			// Parse previousAnnotationValue as JSON
			if err := json.Unmarshal([]byte(previousAnnotationValue), &existingDriverMap); err != nil {
				glog.Errorf(
					"Failed to parse node's %q annotation value (%q) err=%v",
					annotationKey,
					previousAnnotationValue,
					nodeName,
					err)
				os.Exit(1)
			}
		}

		if val, ok := existingDriverMap[driverName]; ok {
			if val == csiDriverNodeId {
				// Value already exists in node annotation, nothing more to do
				glog.V(1).Infof(
					"The key value {%q: %q} alredy eixst in node %q annotation: %v",
					driverName,
					csiDriverNodeId,
					annotationKey,
					previousAnnotationValue)
				os.Exit(0)
			}
		}

		// Add/update annotation value
		existingDriverMap[driverName] = csiDriverNodeId
		jsonObj, err := json.Marshal(existingDriverMap)
		if err != nil {
			glog.Errorf(
				"Failed while trying to add key value {%q: %q} to node %q annotation. Existing value: %v",
				driverName,
				csiDriverNodeId,
				annotationKey,
				previousAnnotationValue)
			os.Exit(1)
		}

		result.ObjectMeta.Annotations = cloneAndAddAnnotation(
			result.ObjectMeta.Annotations,
			annotationKey,
			string(jsonObj))
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
// Returns the given map, if annotationKey is empty.
func cloneAndAddAnnotation(
	annotations map[string]string,
	annotationKey,
	annotationValue string) map[string]string {
	if annotationKey == "" {
		// Don't need to add an annotation.
		return annotations
	}
	// Clone.
	newAnnotations := map[string]string{}
	for key, value := range annotations {
		newAnnotations[key] = value
	}
	newAnnotations[annotationKey] = annotationValue
	return newAnnotations
}
