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
	registerapi "k8s.io/kubernetes/pkg/kubelet/apis/pluginregistration/v1alpha1"

	"github.com/kubernetes-csi/node-driver-registrar/pkg/connection"
)

const (
	// Name of node annotation that contains JSON map of driver names to node
	// names
	annotationKey = "csi.volume.kubernetes.io/nodeid"

	// Default timeout of short CSI calls like GetPluginInfo
	csiTimeout = time.Second

	// Verify (and update, if needed) the node ID at this freqeuency.
	sleepDuration = 2 * time.Minute
)

// Command line flags
var (
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

	if *kubeletRegistrationPath == "" {
		glog.Error("kubelet-registration-path is a required parameter")
		os.Exit(1)
	}

	if *showVersion {
		fmt.Println(os.Args[0], version)
		return
	}
	glog.Infof("Version: %s", version)

	// Once https://github.com/container-storage-interface/spec/issues/159 is
	// resolved, if plugin does not support PUBLISH_UNPUBLISH_VOLUME, then we
	// can skip adding mapping to "csi.volume.kubernetes.io/nodeid" annotation.

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

	// Run forever
	nodeRegister(csiConn, csiDriverName)
}
