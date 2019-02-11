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

package connection

import (
	"context"
	"fmt"

	"github.com/container-storage-interface/spec/lib/go/csi"
	"github.com/kubernetes-csi/csi-lib-utils/connection"
	"google.golang.org/grpc"
)

// CSIConnection is gRPC connection to a remote CSI driver and abstracts all
// CSI calls.
type CSIConnection interface {
	// GetDriverName returns driver name as discovered by GetPluginInfo() gRPC
	// call.
	GetDriverName(ctx context.Context) (string, error)
}

type csiConnection struct {
	conn *grpc.ClientConn
}

var (
	_ CSIConnection = &csiConnection{}
)

//NewConnection return a grpc connection object to a remote CSI driver.
func NewConnection(
	address string) (CSIConnection, error) {
	conn, err := connection.Connect(address)
	if err != nil {
		return nil, err
	}
	return &csiConnection{
		conn: conn,
	}, nil
}

func (c *csiConnection) GetDriverName(ctx context.Context) (string, error) {
	client := csi.NewIdentityClient(c.conn)

	req := csi.GetPluginInfoRequest{}

	rsp, err := client.GetPluginInfo(ctx, &req)
	if err != nil {
		return "", err
	}
	name := rsp.GetName()
	if name == "" {
		return "", fmt.Errorf("name is empty")
	}
	return name, nil
}
