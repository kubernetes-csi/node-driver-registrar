# Release notes for v2.2.0

[Documentation](https://kubernetes-csi.github.io)
# Changelog since v2.1.0

## Changes by Kind

### Uncategorized

- Updated runtime (Go 1.16) and dependencies ([#136](https://github.com/kubernetes-csi/node-driver-registrar/pull/136), [@pohly](https://github.com/pohly))

## Dependencies

### Added
- github.com/moby/spdystream: [v0.2.0](https://github.com/moby/spdystream/tree/v0.2.0)
- github.com/niemeyer/pretty: [a10e7ca](https://github.com/niemeyer/pretty/tree/a10e7ca)
- github.com/yuin/goldmark: [v1.2.1](https://github.com/yuin/goldmark/tree/v1.2.1)

### Changed
- github.com/Azure/go-autorest/autorest: [v0.11.1 → v0.11.12](https://github.com/Azure/go-autorest/autorest/compare/v0.11.1...v0.11.12)
- github.com/cncf/udpa/go: [efcf912 → 5459f2c](https://github.com/cncf/udpa/go/compare/efcf912...5459f2c)
- github.com/container-storage-interface/spec: [v1.3.0 → v1.4.0](https://github.com/container-storage-interface/spec/compare/v1.3.0...v1.4.0)
- github.com/creack/pty: [v1.1.7 → v1.1.11](https://github.com/creack/pty/compare/v1.1.7...v1.1.11)
- github.com/envoyproxy/go-control-plane: [v0.9.7 → fd9021f](https://github.com/envoyproxy/go-control-plane/compare/v0.9.7...fd9021f)
- github.com/fsnotify/fsnotify: [v1.4.9 → v1.4.7](https://github.com/fsnotify/fsnotify/compare/v1.4.9...v1.4.7)
- github.com/go-logr/logr: [v0.3.0 → v0.4.0](https://github.com/go-logr/logr/compare/v0.3.0...v0.4.0)
- github.com/gogo/protobuf: [v1.3.1 → v1.3.2](https://github.com/gogo/protobuf/compare/v1.3.1...v1.3.2)
- github.com/golang/protobuf: [v1.4.3 → v1.5.1](https://github.com/golang/protobuf/compare/v1.4.3...v1.5.1)
- github.com/google/go-cmp: [v0.5.2 → v0.5.5](https://github.com/google/go-cmp/compare/v0.5.2...v0.5.5)
- github.com/gorilla/websocket: [4201258 → v1.4.2](https://github.com/gorilla/websocket/compare/4201258...v1.4.2)
- github.com/kisielk/errcheck: [v1.2.0 → v1.5.0](https://github.com/kisielk/errcheck/compare/v1.2.0...v1.5.0)
- github.com/kr/text: [v0.1.0 → v0.2.0](https://github.com/kr/text/compare/v0.1.0...v0.2.0)
- github.com/kubernetes-csi/csi-lib-utils: [v0.9.0 → v0.9.1](https://github.com/kubernetes-csi/csi-lib-utils/compare/v0.9.0...v0.9.1)
- github.com/moby/term: [672ec06 → df9cb8a](https://github.com/moby/term/compare/672ec06...df9cb8a)
- github.com/prometheus/client_golang: [v1.8.0 → v1.9.0](https://github.com/prometheus/client_golang/compare/v1.8.0...v1.9.0)
- github.com/prometheus/common: [v0.15.0 → v0.19.0](https://github.com/prometheus/common/compare/v0.15.0...v0.19.0)
- github.com/prometheus/procfs: [v0.2.0 → v0.6.0](https://github.com/prometheus/procfs/compare/v0.2.0...v0.6.0)
- golang.org/x/crypto: 7f63de1 → 5ea612d
- golang.org/x/mod: v0.2.0 → v0.3.0
- golang.org/x/net: ac852fb → d523dce
- golang.org/x/sync: cd5d95a → 09787c9
- golang.org/x/sys: 8ad439b → c4fcb01
- golang.org/x/term: 7de9c90 → 6a3ed07
- golang.org/x/text: v0.3.4 → v0.3.5
- golang.org/x/time: 3af7569 → f8bda1e
- golang.org/x/tools: 95d2e58 → 113979e
- google.golang.org/genproto: 40ec1c2 → 75c7a85
- google.golang.org/grpc: v1.34.0 → v1.36.0
- google.golang.org/protobuf: v1.25.0 → v1.26.0
- gopkg.in/check.v1: 41f04d3 → 8fa4692
- gopkg.in/yaml.v2: v2.3.0 → v2.4.0
- gotest.tools/v3: v3.0.2 → v3.0.3
- k8s.io/api: v0.20.0 → v0.21.0
- k8s.io/apimachinery: v0.20.0 → v0.21.0
- k8s.io/client-go: v0.20.0 → v0.21.0
- k8s.io/component-base: v0.20.0 → v0.21.0
- k8s.io/klog/v2: v2.4.0 → v2.8.0
- k8s.io/kube-openapi: d219536 → 591a79e
- k8s.io/kubelet: v0.20.0 → v0.21.0
- sigs.k8s.io/structured-merge-diff/v4: v4.0.2 → v4.1.0

### Removed
- github.com/docker/spdystream: [449fdfc](https://github.com/docker/spdystream/tree/449fdfc)
- gotest.tools: v2.2.0+incompatible
