module github.com/kubernetes-csi/node-driver-registrar

go 1.16

require (
	github.com/container-storage-interface/spec v1.5.0 // indirect
	github.com/kubernetes-csi/csi-lib-utils v0.9.1
	golang.org/x/sys v0.0.0-20210616094352-59db8d763f22
	google.golang.org/grpc v1.38.0
	k8s.io/client-go v0.22.0-rc.0
	k8s.io/klog/v2 v2.9.0
	k8s.io/kubelet v0.22.0-rc.0
)
