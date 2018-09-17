[![Build Status](https://travis-ci.org/kubernetes-csi/driver-registrar.svg?branch=master)](https://travis-ci.org/kubernetes-csi/driver-registrar)
# Driver Registrar

A sidecar container that

1. Registers the containerized CSI driver with kubelet (in the future).
2. Adds the drivers custom `NodeId` (retrieved via `GetNodeID` call) to an annotation on the Kubernetes Node API Object.

## Community, discussion, contribution, and support

Learn how to engage with the Kubernetes community on the [community page](http://kubernetes.io/community/).

You can reach the maintainers of this project at:

- Slack channels
  - [#wg-csi](https://kubernetes.slack.com/messages/wg-csi)
  - [#sig-storage](https://kubernetes.slack.com/messages/sig-storage)
- [Mailing list](https://groups.google.com/forum/#!forum/kubernetes-sig-storage)

### Code of conduct

Participation in the Kubernetes community is governed by the [Kubernetes Code of Conduct](code-of-conduct.md).
