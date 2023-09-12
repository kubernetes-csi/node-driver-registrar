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
	"errors"
	"flag"
	"fmt"
	_ "net/http/pprof"
	"os"
	"path/filepath"
	"strconv"
	"time"

	"github.com/kubernetes-csi/csi-lib-utils/metrics"
	"github.com/kubernetes-csi/node-driver-registrar/pkg/util"
	"github.com/spf13/pflag"
	"k8s.io/klog/v2"

	"github.com/kubernetes-csi/csi-lib-utils/connection"
	csirpc "github.com/kubernetes-csi/csi-lib-utils/rpc"
	"k8s.io/component-base/featuregate"
	logsapi "k8s.io/component-base/logs/api/v1"
	_ "k8s.io/component-base/logs/json/register" // Enable JSON output format.
	registerapi "k8s.io/kubelet/pkg/apis/pluginregistration/v1"
)

const (
	// Name of node annotation that contains JSON map of driver names to node
	// names
	annotationKey = "csi.volume.kubernetes.io/nodeid"

	// Verify (and update, if needed) the node ID at this frequency.
	sleepDuration = 2 * time.Minute
)

const (
	// ModeRegistration runs node-driver-registrar as a long running process
	ModeRegistration = "registration"

	// ModeKubeletRegistrationProbe makes node-driver-registrar act as an exec probe
	// that checks if the kubelet plugin registration succeeded.
	ModeKubeletRegistrationProbe = "kubelet-registration-probe"
)

var (
	// The registration probe path, set when the program runs and used as the path of the file
	// to create when the kubelet plugin registration succeeds.
	registrationProbePath = ""
)

// Command line flags
var (
	connectionTimeout       = flag.Duration("connection-timeout", 0, "The --connection-timeout flag is deprecated")
	operationTimeout        = flag.Duration("timeout", time.Second, "Timeout for waiting for communication with driver")
	csiAddress              = flag.String("csi-address", "/run/csi/socket", "Path of the CSI driver socket that the node-driver-registrar will connect to.")
	pluginRegistrationPath  = flag.String("plugin-registration-path", "/registration", "Path to Kubernetes plugin registration directory.")
	kubeletRegistrationPath = flag.String("kubelet-registration-path", "", "Path of the CSI driver socket on the Kubernetes host machine.")
	healthzPort             = flag.Int("health-port", 0, "(deprecated) TCP port for healthz requests. Set to 0 to disable the healthz server. Only one of `--health-port` and `--http-endpoint` can be set.")
	httpEndpoint            = flag.String("http-endpoint", "", "The TCP network address where the HTTP server for diagnostics, including pprof and the health check indicating whether the registration socket exists, will listen (example: `:8080`). The default is empty string, which means the server is disabled. Only one of `--health-port` and `--http-endpoint` can be set.")
	showVersion             = flag.Bool("version", false, "Show version.")
	mode                    = flag.String("mode", ModeRegistration, `The running mode of node-driver-registrar. "registration" runs node-driver-registrar as a long running process. "kubelet-registration-probe" runs as a health check and returns a status code of 0 if the driver was registered successfully, in the probe definition make sure that the value of --kubelet-registration-path is the same as in the container.`)
	enableProfile           = flag.Bool("enable-pprof", false, "enable pprof profiling")

	// Set during compilation time
	version = "unknown"

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

// newRegistrationServer returns an initialized registrationServer instance
func newRegistrationServer(driverName string, endpoint string, versions []string) registerapi.RegistrationServer {
	return &registrationServer{
		driverName: driverName,
		endpoint:   endpoint,
		version:    versions,
	}
}

// GetInfo is the RPC invoked by plugin watcher
func (e registrationServer) GetInfo(ctx context.Context, req *registerapi.InfoRequest) (*registerapi.PluginInfo, error) {
	logger := klog.FromContext(ctx)
	logger.Info("Received GetInfo call", "request", req)

	// on successful registration, create the registration probe file
	err := util.TouchFile(registrationProbePath)
	if err != nil {
		logger.Error(err, "Failed to create registration probe file", "registrationProbePath", registrationProbePath)
	} else {
		logger.Info("Kubelet registration probe created", "path", registrationProbePath)
	}

	return &registerapi.PluginInfo{
		Type:              registerapi.CSIPlugin,
		Name:              e.driverName,
		Endpoint:          e.endpoint,
		SupportedVersions: e.version,
	}, nil
}

func (e registrationServer) NotifyRegistrationStatus(ctx context.Context, status *registerapi.RegistrationStatus) (*registerapi.RegistrationStatusResponse, error) {
	logger := klog.FromContext(ctx)
	logger.Info("Received NotifyRegistrationStatus call", "status", status)
	if !status.PluginRegistered {
		logger.Error(errors.New(status.Error), "Registration process failed with error, restarting registration container")
		os.Exit(1)
	}

	return &registerapi.RegistrationStatusResponse{}, nil
}

func modeIsKubeletRegistrationProbe() bool {
	return *mode == ModeKubeletRegistrationProbe
}

func main() {
	c := logsapi.NewLoggingConfiguration()
	var pFlagSet pflag.FlagSet
	logsapi.AddFlags(c, &pFlagSet)
	var flagSet flag.FlagSet
	logsapi.AddGoFlags(c, &flagSet)
	pFlagSet.VisitAll(func(f *pflag.Flag) {
		flag.Var(f.Value, f.Name, f.Usage)
	})

	flag.Parse()
	if *showVersion {
		fmt.Println(os.Args[0], version)
		return
	}

	// This command has no --feature-gate parameter, which would make it
	// impossible to use features that are considered alpha and thus
	// disable by default. Therefore those features get enabled here. They
	// are still marked as ALPHA in the command line help.
	fg := featuregate.NewFeatureGate()
	logsapi.AddFeatureGates(fg)
	fg.SetFromMap(map[string]bool{
		string(logsapi.ContextualLogging):   true,
		string(logsapi.LoggingAlphaOptions): true,
	})

	// Switch to desired output format.
	if err := logsapi.ValidateAndApply(c, fg); err != nil {
		klog.ErrorS(err, "")
		klog.FlushAndExit(klog.ExitFlushTimeout, 1)
	}

	if *kubeletRegistrationPath == "" {
		klog.ErrorS(nil, "kubelet-registration-path is a required parameter")
		klog.FlushAndExit(klog.ExitFlushTimeout, 1)
	}
	// set after we made sure that *kubeletRegistrationPath exists
	kubeletRegistrationPathDir := filepath.Dir(*kubeletRegistrationPath)
	registrationProbePath = filepath.Join(kubeletRegistrationPathDir, "registration")

	// with the mode kubelet-registration-probe
	if modeIsKubeletRegistrationProbe() {
		lockfileExists, err := util.DoesFileExist(registrationProbePath)
		if err != nil {
			klog.ErrorS(err, "Failed to check if registration path exists", "registrationProbePath", registrationProbePath)
			klog.FlushAndExit(klog.ExitFlushTimeout, 1)
		}
		if !lockfileExists {
			klog.ErrorS(err, "Kubelet plugin registration hasn't succeeded yets, file doesn't exist", "registrationProbePath", registrationProbePath)
			klog.FlushAndExit(klog.ExitFlushTimeout, 1)
		}
		klog.InfoS("Kubelet plugin registration succeeded")
		os.Exit(0)
	}

	klog.InfoS("Version", "version", version)
	klog.InfoS("Running node-driver-registrar", "mode", *mode)

	if *healthzPort > 0 && *httpEndpoint != "" {
		klog.ErrorS(nil, "Only one of `--health-port` and `--http-endpoint` can be set")
		klog.FlushAndExit(klog.ExitFlushTimeout, 1)
	}
	var addr string
	if *healthzPort > 0 {
		addr = ":" + strconv.Itoa(*healthzPort)
	} else {
		addr = *httpEndpoint
	}

	if *connectionTimeout != 0 {
		klog.InfoS("--connection-timeout is deprecated and will have no effect")
	}

	// Unused metrics manager, necessary for connection.Connect below
	cmm := metrics.NewCSIMetricsManagerForSidecar("")

	// Once https://github.com/container-storage-interface/spec/issues/159 is
	// resolved, if plugin does not support PUBLISH_UNPUBLISH_VOLUME, then we
	// can skip adding mapping to "csi.volume.kubernetes.io/nodeid" annotation.

	klog.V(1).InfoS("Attempting to open a gRPC connection", "csiAddress", *csiAddress)
	csiConn, err := connection.Connect(*csiAddress, cmm)
	if err != nil {
		klog.ErrorS(err, "Error connecting to CSI driver")
		klog.FlushAndExit(klog.ExitFlushTimeout, 1)
	}

	klog.V(1).InfoS("Calling CSI driver to discover driver name")
	ctx, cancel := context.WithTimeout(context.Background(), *operationTimeout)
	defer cancel()

	csiDriverName, err := csirpc.GetDriverName(ctx, csiConn)
	if err != nil {
		klog.ErrorS(err, "Error retreiving CSI driver name")
		klog.FlushAndExit(klog.ExitFlushTimeout, 1)
	}

	klog.V(2).InfoS("CSI driver name", "csiDriverName", csiDriverName)
	cmm.SetDriverName(csiDriverName)

	// Run forever
	nodeRegister(csiDriverName, addr)
}
