//go:build !windows
// +build !windows

/*
Copyright 2019 The Kubernetes Authors.

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
	"fmt"
	"os"
	"path/filepath"

	"golang.org/x/sys/unix"
)

func Umask(mask int) (int, error) {
	return unix.Umask(mask), nil
}

func CleanupSocketFile(socketPath string) error {
	socketExists, err := DoesSocketExist(socketPath)
	if err != nil {
		return err
	}
	if socketExists {
		if err := os.Remove(socketPath); err != nil {
			return fmt.Errorf("failed to remove stale socket %s with error: %+v", socketPath, err)
		}
	}
	return nil
}

func DoesSocketExist(socketPath string) (bool, error) {
	fi, err := os.Stat(socketPath)
	if err == nil {
		if isSocket := (fi.Mode()&os.ModeSocket != 0); isSocket {
			return true, nil
		}
		return false, fmt.Errorf("file exists in socketPath %s but it's not a socket.: %+v", socketPath, fi)
	}
	if err != nil && !os.IsNotExist(err) {
		return false, fmt.Errorf("failed to stat the socket %s with error: %+v", socketPath, err)
	}
	return false, nil
}

func CleanupFile(filePath string) error {
	fileExists, err := DoesFileExist(filePath)
	if err != nil {
		return err
	}
	if fileExists {
		if err := os.Remove(filePath); err != nil {
			return fmt.Errorf("failed to remove stale file=%s with error: %+v", filePath, err)
		}
	}
	return nil
}

func DoesFileExist(filePath string) (bool, error) {
	info, err := os.Stat(filePath)
	if err == nil {
		return info.Mode().IsRegular(), nil
	}
	if err != nil && !os.IsNotExist(err) {
		return false, fmt.Errorf("Failed to stat the file=%s with error: %+v", filePath, err)
	}
	return false, nil
}

func TouchFile(filePath string) error {
	exists, err := DoesFileExist(filePath)
	if err != nil {
		return err
	}
	if !exists {
		err := os.MkdirAll(filepath.Dir(filePath), 0755)
		if err != nil {
			return err
		}

		file, err := os.Create(filePath)
		if err != nil {
			return err
		}
		file.Close()
	}
	return nil
}
