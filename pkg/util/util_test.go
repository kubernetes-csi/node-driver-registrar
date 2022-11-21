/*
Copyright 2020 The Kubernetes Authors.

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

package util

import (
	"net"
	"os"
	"path/filepath"
	"runtime"
	"testing"

	utiltesting "k8s.io/client-go/util/testing"
	"k8s.io/klog/v2"
)

var socketFileName = "reg.sock"
var kubeletRegistrationPath = "/var/lib/kubelet/plugins/csi-dummy/registration"

// TestSocketFileDoesNotExist - Test1: file does not exist. So clean up should be successful.
func TestSocketFileDoesNotExist(t *testing.T) {
	// Create a temp directory
	testDir, err := utiltesting.MkTmpdir("csi-test")
	if err != nil {
		t.Fatalf("could not create temp dir: %v", err)
	}
	defer os.RemoveAll(testDir)

	socketPath := filepath.Join(testDir, socketFileName)
	socketExists, err := DoesSocketExist(socketPath)
	if err != nil {
		t.Fatalf("check for existence returned error: %+v", err)
	}
	if socketExists {
		t.Fatalf("socket exists when it should not")
	}
	// Negative test, file is not created. So file name in current path used.
	err = CleanupSocketFile(socketPath)
	if err != nil {
		t.Fatalf("cleanup returned error: %+v", err)
	}
}

// TestSocketPathDoesNotExist - Test2: directory does not exist. So clean up should be successful.
func TestSocketPathDoesNotExist(t *testing.T) {
	// Create a temp directory and delete it. This way we know the directory
	// does not exist.
	testDir, err := utiltesting.MkTmpdir("csi-test")
	if err != nil {
		t.Fatalf("could not create temp dir: %v", err)
	}
	os.RemoveAll(testDir)

	socketPath := filepath.Join(testDir, socketFileName)
	socketExists, err := DoesSocketExist(socketPath)
	if err != nil {
		t.Fatalf("check for existence returned error: %+v", err)
	}
	if socketExists {
		t.Fatalf("socket exists when it should not")
	}
	err = CleanupSocketFile(socketPath)
	if err != nil {
		t.Fatalf("cleanup returned error: %+v", err)
	}
}

// TestSocketPathSimple - Test3: +ve test create socket and call delete. Regular happy path scenario.
func TestSocketPathSimple(t *testing.T) {
	// Create a temp directory
	testDir, err := utiltesting.MkTmpdir("csi-test")
	if err != nil {
		t.Fatalf("could not create temp dir: %v", err)
	}
	defer os.RemoveAll(testDir)

	socketPath := filepath.Join(testDir, socketFileName)

	_, err = net.Listen("unix", socketPath)
	if err != nil {
		klog.Errorf("failed to listen on socket: %s with error: %+v", socketPath, err)
		os.Exit(1)
	}

	socketExists, err := DoesSocketExist(socketPath)
	if err != nil {
		t.Fatalf("check for existence returned error: %+v", err)
	}
	if !socketExists {
		t.Fatalf("socket does not exist when it should")
	}

	err = CleanupSocketFile(socketPath)
	if err != nil {
		t.Fatalf("cleanup returned error: %+v", err)
	}

	_, err = os.Lstat(socketPath)
	if err != nil {
		if !os.IsNotExist(err) {
			t.Fatalf("lstat error on file %s ", socketPath)
		}
	} else {
		t.Fatalf("socket file %s exists", socketPath)
	}
}

// TestSocketRegularFile - Test 4: Create a regular file and check the deletion logic
func TestSocketRegularFile(t *testing.T) {
	// Create a temp directory
	testDir, err := utiltesting.MkTmpdir("csi-test")
	if err != nil {
		t.Fatalf("could not create temp dir: %v", err)
	}
	defer os.RemoveAll(testDir)

	socketPath := filepath.Join(testDir, socketFileName)
	f, err := os.Create(socketPath)
	if err != nil {
		t.Fatalf("create file failed: %s", socketPath)
	}
	f.Close()

	socketExists, err := DoesSocketExist(socketPath)
	if err == nil && runtime.GOOS != "windows" {
		t.Fatalf("check for existence should returned error on linux: %+v", err)
	}
	// See comments in CleanupSocketFile for differences in windows and linux behavior checking for sockets.
	if runtime.GOOS == "windows" {
		if !socketExists {
			t.Fatalf("socket does not exist when it should")
		}
	} else if socketExists {
		t.Fatalf("socket exists when it should not")
	}

	err = CleanupSocketFile(socketPath)
	if err == nil && runtime.GOOS != "windows" {
		t.Fatalf("cleanup returned should return error on linux: %+v", err)
	}

	if runtime.GOOS == "windows" {
		// In windows a regular file will be deleted without checking whether
		// its a socket.
		_, err = os.Lstat(socketPath)
		if err != nil {
			if !os.IsNotExist(err) {
				t.Fatalf("lstat error on file %s ", socketPath)
			}
		} else {
			t.Fatalf("regular file %s exists", socketPath)
		}
	} else {
		_, err = os.Lstat(socketPath)
		if err != nil {
			if os.IsNotExist(err) {
				t.Fatalf("regular file %s got deleted", socketPath)
			} else {
				t.Fatalf("lstat error on file %s ", socketPath)
			}
		}
	}
}

// TestTouchFile creates a file if it doesn't exist
func TestTouchFile(t *testing.T) {
	// Create a temp directory
	testDir, err := utiltesting.MkTmpdir("csi-test")
	if err != nil {
		t.Fatalf("could not create temp dir: %v", err)
	}
	defer os.RemoveAll(testDir)

	filePath := filepath.Join(testDir, kubeletRegistrationPath)
	fileExists, err := DoesFileExist(filePath)
	if err != nil {
		t.Fatalf("Failed to execute file exist: %+v", err)
	}
	if fileExists {
		t.Fatalf("File %s must not exist", filePath)
	}

	// returns an error only if it failed to clean the file, not if the file didn't exist
	err = CleanupFile(filePath)
	if err != nil {
		t.Fatalf("Failed to execute file cleanup: %+v", err)
	}

	err = TouchFile(filePath)
	if err != nil {
		t.Fatalf("Failed to execute file touch: %+v", err)
	}

	fileExists, err = DoesFileExist(filePath)
	if err != nil {
		t.Fatalf("Failed to execute file exist: %+v", err)
	}
	if !fileExists {
		t.Fatalf("File %s must exist", filePath)
	}
}
