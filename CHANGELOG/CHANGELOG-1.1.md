# Changelog since v1.0.2

## Deprecations
* Command line flag `-connection-timeout` is deprecated and has no effect.

## Notable Features

* The driver registrar now tries to connect to CSI driver indefinitely. ([#29](https://github.com/kubernetes-csi/node-driver-registrar/pull/29))

## Other notable changes

* Use distroless as base image ([#34](https://github.com/kubernetes-csi/node-driver-registrar/pull/34))
* Use GetDriverName from csi-lib-utils ([#33](https://github.com/kubernetes-csi/node-driver-registrar/pull/33))
* Migrate to k8s.io/klog from glog. ([#24](https://github.com/kubernetes-csi/node-driver-registrar/pull/24))
* Update compatibility matrix to only reflect branch head ([#25](https://github.com/kubernetes-csi/node-driver-registrar/pull/25))
* Update documentation and argument descriptions ([#13](https://github.com/kubernetes-csi/node-driver-registrar/pull/13))
* Cleanup vendor ([#11](https://github.com/kubernetes-csi/node-driver-registrar/pull/11))
* Add csi prefix to image name ([#5](https://github.com/kubernetes-csi/node-driver-registrar/pull/5))
