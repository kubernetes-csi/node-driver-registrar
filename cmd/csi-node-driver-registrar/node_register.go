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
	"net/http"
	"net/http/pprof"
	"os"
	"os/signal"
	"runtime"
	"syscall"

	"google.golang.org/grpc"

	"github.com/kubernetes-csi/csi-lib-utils/connection"
	"github.com/kubernetes-csi/node-driver-registrar/pkg/util"
	"k8s.io/klog/v2"
	registerapi "k8s.io/kubelet/pkg/apis/pluginregistration/v1"
)

func nodeRegister(csiDriverName, httpEndpoint string) {
	// When kubeletRegistrationPath is specified then driver-registrar ONLY acts
	// as gRPC server which replies to registration requests initiated by kubelet's
	// plugins watcher infrastructure. Node labeling is done by kubelet's csi code.
	registrar := newRegistrationServer(csiDriverName, *kubeletRegistrationPath, supportedVersions)
	socketPath := buildSocketPath(csiDriverName)
	if err := util.CleanupSocketFile(socketPath); err != nil {
		klog.ErrorS(err, "")
		klog.FlushAndExit(klog.ExitFlushTimeout, 1)
	}

	var oldmask int
	if runtime.GOOS == "linux" {
		// Default to only user accessible socket, caller can open up later if desired
		oldmask, _ = util.Umask(0077)
	}

	klog.InfoS("Starting Registration Server", "socketPath", socketPath)
	lis, err := net.Listen("unix", socketPath)
	if err != nil {
		klog.ErrorS(err, "Failed to listen on socket", "socketPath", socketPath)
		klog.FlushAndExit(klog.ExitFlushTimeout, 1)
	}
	if runtime.GOOS == "linux" {
		util.Umask(oldmask)
	}
	klog.InfoS("Registration Server started", "socketPath", socketPath)
	grpcServer := grpc.NewServer()

	// Registers kubelet plugin watcher api.
	registerapi.RegisterRegistrationServer(grpcServer, registrar)

	go httpServer(socketPath, httpEndpoint, csiDriverName)
	go removeRegSocket(csiDriverName)
	// Starts service
	if err := grpcServer.Serve(lis); err != nil {
		klog.ErrorS(err, "Registration Server stopped serving")
		klog.FlushAndExit(klog.ExitFlushTimeout, 1)
	}

	// If gRPC server is gracefully shutdown, cleanup and exit
	os.Exit(0)
}

func buildSocketPath(csiDriverName string) string {
	return fmt.Sprintf("%s/%s-reg.sock", *pluginRegistrationPath, csiDriverName)
}

func httpServer(socketPath string, httpEndpoint string, csiDriverName string) {
	if httpEndpoint == "" {
		klog.InfoS("Skipping HTTP server")
		return
	}
	klog.InfoS("Starting HTTP server", "endpoint", httpEndpoint)

	// Prepare http endpoint for healthz + profiling (if enabled)
	mux := http.NewServeMux()
	mux.HandleFunc("/healthz", func(w http.ResponseWriter, req *http.Request) {
		socketExists, err := util.DoesSocketExist(socketPath)
		if err == nil && socketExists {
			grpcSocketCheckError := checkLiveRegistrationSocket(socketPath, csiDriverName)
			if grpcSocketCheckError != nil {
				w.WriteHeader(http.StatusInternalServerError)
				w.Write([]byte(grpcSocketCheckError.Error()))
				klog.ErrorS(grpcSocketCheckError, "Health check failed")
			} else {
				w.WriteHeader(http.StatusOK)
				w.Write([]byte(`ok`))
				klog.V(5).InfoS("Health check succeeded")
			}
		} else if err != nil {
			w.WriteHeader(http.StatusInternalServerError)
			w.Write([]byte(err.Error()))
			klog.ErrorS(err, "Health check failed")
		} else if !socketExists {
			w.WriteHeader(http.StatusNotFound)
			w.Write([]byte("registration socket does not exist"))
			klog.ErrorS(nil, "Health check failed, registration socket does not exist")
		}
	})

	if *enableProfile {
		klog.InfoS("Starting profiling", "endpoint", httpEndpoint)

		mux.HandleFunc("/debug/pprof/", pprof.Index)
		mux.HandleFunc("/debug/pprof/cmdline", pprof.Cmdline)
		mux.HandleFunc("/debug/pprof/profile", pprof.Profile)
		mux.HandleFunc("/debug/pprof/symbol", pprof.Symbol)
		mux.HandleFunc("/debug/pprof/trace", pprof.Trace)
	}

	klog.ErrorS(http.ListenAndServe(httpEndpoint, mux), "")
	klog.FlushAndExit(klog.ExitFlushTimeout, 1)
}

func checkLiveRegistrationSocket(socketFile, csiDriverName string) error {
	klog.V(2).InfoS("Attempting to open a gRPC connection", "socketfile", socketFile)
	grpcConn, err := connection.ConnectWithoutMetrics(socketFile)
	if err != nil {
		return fmt.Errorf("error connecting to node-registrar socket %s: %v", socketFile, err)
	}

	defer closeGrpcConnection(socketFile, grpcConn)

	klog.V(2).InfoS("Calling node registrar to check if it still responds")
	ctx, cancel := context.WithTimeout(context.Background(), *operationTimeout)
	defer cancel()

	client := registerapi.NewRegistrationClient(grpcConn)

	infoRequest := &registerapi.InfoRequest{}

	info, err := client.GetInfo(ctx, infoRequest)
	if err != nil {
		return fmt.Errorf("error getting info from node-registrar socket: %v", err)
	}

	if info.Name == csiDriverName {
		return nil
	}
	return fmt.Errorf("invalid driver name %s", info.Name)
}

func closeGrpcConnection(socketFile string, conn *grpc.ClientConn) {
	err := conn.Close()
	if err != nil {
		klog.ErrorS(err, "Error closing socket", "socketfile", socketFile)
		klog.FlushAndExit(klog.ExitFlushTimeout, 1)
	}
}

func removeRegSocket(csiDriverName string) {
	sigc := make(chan os.Signal, 1)
	signal.Notify(sigc, syscall.SIGTERM)
	<-sigc
	socketPath := buildSocketPath(csiDriverName)
	err := os.Remove(socketPath)
	if err != nil && !os.IsNotExist(err) {
		klog.ErrorS(err, "Failed to remove socket with error", "socket", socketPath)
		klog.FlushAndExit(klog.ExitFlushTimeout, 1)
	}
	os.Exit(0)
}
