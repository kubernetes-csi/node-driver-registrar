[![Build Status](https://travis-ci.org/kubernetes-csi/driver-registrar.svg?branch=master)](https://travis-ci.org/kubernetes-csi/driver-registrar)
# Driver Registrar

A sidecar container that

1. Registers the containerized CSI driver with kubelet (in the future).
2. Adds the drivers custom `NodeId` (retrieved via `GetNodeID` call) to an annotation on the Kubernetes Node API Object.
