//go:build windows
// +build windows

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
	"errors"
	"fmt"
	"os"
	"path/filepath"
)

func Umask(mask int) (int, error) {
	return -1, errors.New("umask not supported in Windows")
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
	// TODO: Until the bug - https://github.com/golang/go/issues/33357 is fixed, os.stat wouldn't return the
	// right mode(socket) on windows. Hence deleting the file, without checking whether
	// its a socket, on windows.
	if _, err := os.Lstat(socketPath); err != nil {
		if os.IsNotExist(err) {
			return false, nil
		}
		return false, fmt.Errorf("failed to lstat the socket %s with error: %+v", socketPath, err)
	}
	return true, nil
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
	info, err := os.Lstat(filePath)
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
