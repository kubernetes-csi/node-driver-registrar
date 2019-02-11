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
	"testing"

	"github.com/golang/mock/gomock"
	"github.com/kubernetes-csi/csi-test/driver"
)

func createMockServer(t *testing.T) (
	*gomock.Controller,
	*driver.MockCSIDriver,
	*driver.MockIdentityServer,
	*driver.MockControllerServer,
	*driver.MockNodeServer,
	CSIConnection,
	error) {
	// Start the mock server
	mockController := gomock.NewController(t)
	identityServer := driver.NewMockIdentityServer(mockController)
	controllerServer := driver.NewMockControllerServer(mockController)
	nodeServer := driver.NewMockNodeServer(mockController)
	drv := driver.NewMockCSIDriver(&driver.MockCSIDriverServers{
		Identity:   identityServer,
		Controller: controllerServer,
		Node:       nodeServer,
	})
	drv.Start()

	// Create a client connection to it
	addr := drv.Address()
	csiConn, err := NewConnection(addr)
	if err != nil {
		return nil, nil, nil, nil, nil, nil, err
	}

	return mockController, drv, identityServer, controllerServer, nodeServer, csiConn, nil
}
