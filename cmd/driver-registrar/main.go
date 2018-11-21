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
	"flag"
	"fmt"
	"os"
	"time"

	"github.com/golang/glog"
	"k8s.io/client-go/rest"
	"k8s.io/client-go/tools/clientcmd"
	registerapi "k8s.io/kubernetes/pkg/kubelet/apis/pluginregistration/v1alpha1"

	"github.com/kubernetes-csi/driver-registrar/pkg/connection"
)

const (
	// Name of node annotation that contains JSON map of driver names to node
	// names
	annotationKey = "csi.volume.kubernetes.io/nodeid"

	// Default timeout of short CSI calls like GetPluginInfo
	csiTimeout = time.Second

	// Verify (and update, if needed) the node ID at this freqeuency.
	sleepDuration = 2 * time.Minute

	modeNodeRegister = "node-register"
	modeKubeRegister = "kubernetes-register"
	defaultMode      = modeNodeRegister
)

// Command line flags
var (
	kubeconfig = flag.String("kubeconfig", "", "Absolute path to the kubeconfig file. Required only when running out of cluster.")
	runMode    = flag.String("mode", defaultMode, "Mode to run the program. Supported modes: node-register, kubernetes-register.\n"+
		"In "+modeNodeRegister+" mode, the program will register the CSI driver with the "+
		"node, setting up the node information accordingly.\n"+
		"In "+modeKubeRegister+" mode, the program will register the driver with "+
		"Kubernetes. In this mode this program will setup all the necessary information "+
		"to register the CSI driver with Kubernetes. This mode requires that this "+
		"container be run in a StateFul set of 1, and not in a DaemonSet.")
	k8sAttachmentRequired = flag.Bool("driver-requires-attachment",
		true,
		"Indicates this CSI volume driver requires an attach operation (because it "+
			"implements the CSI ControllerPublishVolume() method), and that Kubernetes "+
			"should call attach and wait for any attach operation to complete before "+
			"proceeding to mounting. If value is not specified, default is false meaning "+
			"attach will not be called.")
	k8sPodInfoOnMountVersion = flag.String("pod-info-mount-version",
		"",
		"This indicates that the associated CSI volume driver"+
			"requires additional pod information (like podName, podUID, etc.) during mount."+
			"A version of value \"v1\" will cause the Kubelet send the followings pod information "+
			"during NodePublishVolume() calls to the driver as VolumeAttributes:"+
			"- csi.storage.k8s.io/pod.name: pod.Name\n"+
			"- csi.storage.k8s.io/pod.namespace: pod.Namespace\n"+
			"- csi.storage.k8s.io/pod.uid: string(pod.UID)",
	)
	connectionTimeout       = flag.Duration("connection-timeout", 1*time.Minute, "Timeout for waiting for CSI driver socket.")
	csiAddress              = flag.String("csi-address", "/run/csi/socket", "Address of the CSI driver socket.")
	kubeletRegistrationPath = flag.String("kubelet-registration-path", "",
		`Enables Kubelet Plugin Registration service, and returns the specified path as "endpoint" in "PluginInfo" response.
		 If this option is set, the driver-registrar expose a unix domain socket to handle Kubelet Plugin Registration, 
		 this socket MUST be surfaced on the host in the kubelet plugin registration directory (in addition to the CSI driver socket). 
		 If plugin registration is enabled on kubelet (kubelet flag KubeletPluginsWatcher is set), then this option should be set
		 and the value should be the path of the CSI driver socket on the host machine.`)
	showVersion = flag.Bool("version", false, "Show version.")
	version     = "unknown"
	// List of supported versions
	supportedVersions = []string{"1.0.0"}
)

// registrationServer is a sample plugin to work with plugin watcher
type registrationServer struct {
	driverName string
	endpoint   string
	version    []string
}

var _ registerapi.RegistrationServer = registrationServer{}

// NewregistrationServer returns an initialized registrationServer instance
func newRegistrationServer(driverName string, endpoint string, versions []string) registerapi.RegistrationServer {
	return &registrationServer{
		driverName: driverName,
		endpoint:   endpoint,
		version:    versions,
	}
}

// GetInfo is the RPC invoked by plugin watcher
func (e registrationServer) GetInfo(ctx context.Context, req *registerapi.InfoRequest) (*registerapi.PluginInfo, error) {
	glog.Infof("Received GetInfo call: %+v", req)
	return &registerapi.PluginInfo{
		Type:              registerapi.CSIPlugin,
		Name:              e.driverName,
		Endpoint:          e.endpoint,
		SupportedVersions: e.version,
	}, nil
}

func (e registrationServer) NotifyRegistrationStatus(ctx context.Context, status *registerapi.RegistrationStatus) (*registerapi.RegistrationStatusResponse, error) {
	glog.Infof("Received NotifyRegistrationStatus call: %+v", status)
	if !status.PluginRegistered {
		glog.Errorf("Registration process failed with error: %+v, restarting registration container.", status.Error)
		os.Exit(1)
	}

	return &registerapi.RegistrationStatusResponse{}, nil
}

func main() {
	flag.Set("logtostderr", "true")
	flag.Parse()

	if *showVersion {
		fmt.Println(os.Args[0], version)
		return
	}
	glog.Infof("Version: %s", version)

	// Once https://github.com/container-storage-interface/spec/issues/159 is
	// resolved, if plugin does not support PUBLISH_UNPUBLISH_VOLUME, then we
	// can skip adding mappting to "csi.volume.kubernetes.io/nodeid" annotation.

	// Connect to CSI.
	glog.V(1).Infof("Attempting to open a gRPC connection with: %q", *csiAddress)
	csiConn, err := connection.NewConnection(*csiAddress, *connectionTimeout)
	if err != nil {
		glog.Error(err.Error())
		os.Exit(1)
	}

	// Get CSI driver name.
	glog.V(1).Infof("Calling CSI driver to discover driver name.")
	ctx, cancel := context.WithTimeout(context.Background(), csiTimeout)
	defer cancel()
	csiDriverName, err := csiConn.GetDriverName(ctx)
	if err != nil {
		glog.Error(err.Error())
		os.Exit(1)
	}
	glog.V(2).Infof("CSI driver name: %q", csiDriverName)

	// Create the client config. Use kubeconfig if given, otherwise assume
	// in-cluster.
	glog.V(1).Infof("Loading kubeconfig.")
	config, err := buildConfig(*kubeconfig)
	if err != nil {
		glog.Error(err.Error())
		os.Exit(1)
	}

	// Check mode
	switch *runMode {
	case modeNodeRegister:
		nodeRegister(config, csiConn, csiDriverName)
	case modeKubeRegister:
		kubernetesRegister(config, csiConn, csiDriverName)
	default:
		glog.Errorf("Unknown mode: %s", *runMode)
		fmt.Fprintf(os.Stderr, "Unknown mode: %s", *runMode)
	}
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
