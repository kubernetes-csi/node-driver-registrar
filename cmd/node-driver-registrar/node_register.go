/*
Copyright 2018 The Kubernetes Authors.

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
	"fmt"
	"net"
	"os"

	"google.golang.org/grpc"

	"github.com/golang/glog"
	"golang.org/x/sys/unix"
	registerapi "k8s.io/kubernetes/pkg/kubelet/apis/pluginregistration/v1alpha1"

	"github.com/kubernetes-csi/node-driver-registrar/pkg/connection"
)

func nodeRegister(
	csiConn connection.CSIConnection,
	csiDriverName string,
) {
	// Fetch node name from environment variable
	k8sNodeName := os.Getenv("KUBE_NODE_NAME")
	if k8sNodeName == "" {
		glog.Error("Node name not found. The environment variable KUBE_NODE_NAME is empty.")
		os.Exit(1)
	}

	// Get CSI Driver Node ID
	glog.V(1).Infof("Calling CSI driver to discover node ID.")
	ctx, cancel := context.WithTimeout(context.Background(), csiTimeout)
	defer cancel()
	csiDriverNodeId, err := csiConn.NodeGetId(ctx)
	if err != nil {
		glog.Error(err.Error())
		os.Exit(1)
	}
	glog.V(2).Infof("CSI driver node ID: %q", csiDriverNodeId)

	// When kubeletRegistrationPath is specified then driver-registrar ONLY acts
	// as gRPC server which replies to registration requests initiated by kubelet's
	// pluginswatcher infrastructure. Node labeling is done by kubelet's csi code.
	registrar := newRegistrationServer(csiDriverName, *kubeletRegistrationPath, supportedVersions)
	socketPath := fmt.Sprintf("/registration/%s-reg.sock", csiDriverName)
	fi, err := os.Stat(socketPath)
	if err == nil && (fi.Mode()&os.ModeSocket) != 0 {
		// Remove any socket, stale or not, but fall through for other files
		if err := os.Remove(socketPath); err != nil {
			glog.Errorf("failed to remove stale socket %s with error: %+v", socketPath, err)
			os.Exit(1)
		}
	}
	if err != nil && !os.IsNotExist(err) {
		glog.Errorf("failed to stat the socket %s with error: %+v", socketPath, err)
		os.Exit(1)
	}
	// Default to only user accessible socket, caller can open up later if desired
	oldmask := unix.Umask(0077)

	glog.Infof("Starting Registration Server at: %s\n", socketPath)
	lis, err := net.Listen("unix", socketPath)
	if err != nil {
		glog.Errorf("failed to listen on socket: %s with error: %+v", socketPath, err)
		os.Exit(1)
	}
	unix.Umask(oldmask)
	glog.Infof("Registration Server started at: %s\n", socketPath)
	grpcServer := grpc.NewServer()
	// Registers kubelet plugin watcher api.
	registerapi.RegisterRegistrationServer(grpcServer, registrar)

	// Starts service
	if err := grpcServer.Serve(lis); err != nil {
		glog.Errorf("Registration Server stopped serving: %v", err)
		os.Exit(1)
	}
	// If gRPC server is gracefully shutdown, exit
	os.Exit(0)
}
