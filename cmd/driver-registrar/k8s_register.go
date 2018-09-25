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
	"os"
	"os/signal"
	"time"

	"github.com/golang/glog"
	crdclient "k8s.io/apiextensions-apiserver/pkg/client/clientset/clientset"
	apierrors "k8s.io/apimachinery/pkg/api/errors"
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	"k8s.io/client-go/rest"
	"k8s.io/client-go/util/retry"
	k8scsi "k8s.io/csi-api/pkg/apis/csi/v1alpha1"
	k8scsiclient "k8s.io/csi-api/pkg/client/clientset/versioned"
	k8scsicrd "k8s.io/csi-api/pkg/crd"

	"github.com/kubernetes-csi/driver-registrar/pkg/connection"
)

func kubernetesRegister(
	config *rest.Config,
	csiConn connection.CSIConnection,
	csiDriverName string,
) {
	// Get client info to CSIDriver
	clientset, err := k8scsiclient.NewForConfig(config)
	if err != nil {
		glog.Error(err.Error())
		os.Exit(1)
	}

	// Set spec
	spec := &k8scsi.CSIDriverSpec{
		AttachRequired:        k8sAttachmentRequired,
		PodInfoOnMountVersion: k8sPodInfoOnMountVersion,
	}
	glog.V(1).Infof("AttachRequired: %v", *k8sAttachmentRequired)
	glog.V(1).Infof("PodInfoOnMountVersion: %v", *k8sPodInfoOnMountVersion)

	// Register CRD
	glog.V(1).Info("Registering " + k8scsi.CsiDriverResourcePlural)
	crdclient, err := crdclient.NewForConfig(config)
	if err != nil {
		glog.Error(err.Error())
		os.Exit(1)
	}
	crdv1beta1client := crdclient.ApiextensionsV1beta1().CustomResourceDefinitions()
	_, err = crdv1beta1client.Create(k8scsicrd.CSIDriverCRD())
	if err == nil {
		glog.V(1).Info("CSIDriver CRD registered")
	} else if apierrors.IsAlreadyExists(err) {
		glog.V(1).Info("CSIDriver CRD already had been registered")
	} else if err != nil {
		glog.Error(err.Error())
		os.Exit(1)
	}

	// Set up goroutine to cleanup (aka deregister) on termination.
	c := make(chan os.Signal, 1)
	signal.Notify(c, os.Interrupt)
	go func() {
		<-c
		verifyAndDeleteCSIDriverInfo(
			clientset,
			csiDriverName,
			spec)
		os.Exit(1)
	}()

	// Run forever
	for {
		verifyAndAddCSIDriverInfo(clientset, csiDriverName, spec)
		time.Sleep(sleepDuration)
	}
}

// Registers CSI driver by creating a CSIDriver object
func verifyAndAddCSIDriverInfo(
	csiClientset *k8scsiclient.Clientset,
	csiDriverName string,
	spec *k8scsi.CSIDriverSpec,
) error {
	retryErr := retry.RetryOnConflict(retry.DefaultRetry, func() error {
		csidrivers := csiClientset.CsiV1alpha1().CSIDrivers()

		// Create it
		csiDriver := &k8scsi.CSIDriver{
			ObjectMeta: metav1.ObjectMeta{
				Name: csiDriverName,
			},
			Spec: *spec,
		}

		_, err := csidrivers.Create(csiDriver)
		if err == nil {
			glog.V(1).Infof("CSIDRiver object created for driver %s", csiDriverName)
			return nil
		} else if apierrors.IsAlreadyExists(err) {
			return nil
		} else {
			glog.Errorf("Failed to create CSIDriver object: %v", err)
			return err
		}
	})
	if retryErr != nil {
		return retryErr
	}
	return nil
}

// Deregister CSI Driver by deleting CSIDriver object
func verifyAndDeleteCSIDriverInfo(
	csiClientset *k8scsiclient.Clientset,
	csiDriverName string,
	spec *k8scsi.CSIDriverSpec,
) error {
	retryErr := retry.RetryOnConflict(retry.DefaultRetry, func() error {
		csidrivers := csiClientset.CsiV1alpha1().CSIDrivers()
		err := csidrivers.Delete(csiDriverName, &metav1.DeleteOptions{})
		if err == nil {
			glog.V(1).Infof("CSIDRiver object deleted for driver %s", csiDriverName)
			return nil
		} else if apierrors.IsNotFound(err) {
			glog.V(1).Info("No need to clean up CSIDriver since it does not exist")
			return nil
		} else {
			glog.Errorf("Failed to create CSIDriver object: %v", err)
			return err
		}
	})
	if retryErr != nil {
		return retryErr
	}
	return nil
}
