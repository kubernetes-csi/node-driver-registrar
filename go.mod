module github.com/kubernetes-csi/node-driver-registrar

go 1.13

require (
	github.com/container-storage-interface/spec v1.0.0 // indirect
	github.com/kubernetes-csi/csi-lib-utils v0.4.0-rc1
	golang.org/x/sys v0.0.0-20200622214017-ed371f2e16b4
	google.golang.org/grpc v1.27.0
	k8s.io/client-go v0.19.0-rc.2
	k8s.io/klog v1.0.0
	k8s.io/kubelet v0.19.0-rc.2
)
