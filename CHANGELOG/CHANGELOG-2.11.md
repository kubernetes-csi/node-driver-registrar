# Release notes for v2.11.0

[Documentation](https://kubernetes-csi.github.io)

# Changelog since v2.10.0

## Changes by Kind

### Feature

- Added support for contextual logging. ([#410](https://github.com/kubernetes-csi/node-driver-registrar/pull/410), [@bells17](https://github.com/bells17))
- Added support for structured logging (the log messages have been changed due to the activation of structured logging) ([#349](https://github.com/kubernetes-csi/node-driver-registrar/pull/349), [@rlia](https://github.com/rlia))
- Updated Kubernetes deps to v1.30 ([#406](https://github.com/kubernetes-csi/node-driver-registrar/pull/406), [@jsafrane](https://github.com/jsafrane))

### Uncategorized

- Update google.golang.org/protobuf to v1.33.0 to resolve CVE-2024-24786 ([#387](https://github.com/kubernetes-csi/node-driver-registrar/pull/387), [@dobsonj](https://github.com/dobsonj))

## Dependencies

### Added
- github.com/cpuguy83/go-md2man/v2: [v2.0.2](https://github.com/cpuguy83/go-md2man/v2/tree/v2.0.2)
- github.com/fxamacker/cbor/v2: [v2.6.0](https://github.com/fxamacker/cbor/v2/tree/v2.6.0)
- github.com/go-task/slim-sprig/v3: [v3.0.0](https://github.com/go-task/slim-sprig/v3/tree/v3.0.0)
- github.com/russross/blackfriday/v2: [v2.1.0](https://github.com/russross/blackfriday/v2/tree/v2.1.0)
- github.com/x448/float16: [v0.8.4](https://github.com/x448/float16/tree/v0.8.4)
- go.uber.org/goleak: v1.3.0
- k8s.io/gengo/v2: 51d4e06

### Changed
- cloud.google.com/go/compute: v1.23.0 → v1.24.0
- github.com/alecthomas/kingpin/v2: [v2.3.2 → v2.4.0](https://github.com/alecthomas/kingpin/v2/compare/v2.3.2...v2.4.0)
- github.com/cespare/xxhash/v2: [v2.2.0 → v2.3.0](https://github.com/cespare/xxhash/v2/compare/v2.2.0...v2.3.0)
- github.com/cncf/xds/go: [e9ce688 → 0fa0005](https://github.com/cncf/xds/go/compare/e9ce688...0fa0005)
- github.com/emicklei/go-restful/v3: [v3.11.0 → v3.12.0](https://github.com/emicklei/go-restful/v3/compare/v3.11.0...v3.12.0)
- github.com/envoyproxy/go-control-plane: [v0.11.1 → v0.12.0](https://github.com/envoyproxy/go-control-plane/compare/v0.11.1...v0.12.0)
- github.com/envoyproxy/protoc-gen-validate: [v1.0.2 → v1.0.4](https://github.com/envoyproxy/protoc-gen-validate/compare/v1.0.2...v1.0.4)
- github.com/go-logr/logr: [v1.3.0 → v1.4.1](https://github.com/go-logr/logr/compare/v1.3.0...v1.4.1)
- github.com/go-logr/zapr: [v1.2.3 → v1.3.0](https://github.com/go-logr/zapr/compare/v1.2.3...v1.3.0)
- github.com/go-openapi/jsonpointer: [v0.19.6 → v0.21.0](https://github.com/go-openapi/jsonpointer/compare/v0.19.6...v0.21.0)
- github.com/go-openapi/jsonreference: [v0.20.2 → v0.21.0](https://github.com/go-openapi/jsonreference/compare/v0.20.2...v0.21.0)
- github.com/go-openapi/swag: [v0.22.3 → v0.23.0](https://github.com/go-openapi/swag/compare/v0.22.3...v0.23.0)
- github.com/golang/glog: [v1.1.2 → v1.2.0](https://github.com/golang/glog/compare/v1.1.2...v1.2.0)
- github.com/golang/protobuf: [v1.5.3 → v1.5.4](https://github.com/golang/protobuf/compare/v1.5.3...v1.5.4)
- github.com/google/pprof: [4bb14d4 → a892ee0](https://github.com/google/pprof/compare/4bb14d4...a892ee0)
- github.com/google/uuid: [v1.3.1 → v1.6.0](https://github.com/google/uuid/compare/v1.3.1...v1.6.0)
- github.com/kubernetes-csi/csi-lib-utils: [v0.17.0 → v0.18.0](https://github.com/kubernetes-csi/csi-lib-utils/compare/v0.17.0...v0.18.0)
- github.com/onsi/ginkgo/v2: [v2.13.0 → v2.17.2](https://github.com/onsi/ginkgo/v2/compare/v2.13.0...v2.17.2)
- github.com/onsi/gomega: [v1.29.0 → v1.33.1](https://github.com/onsi/gomega/compare/v1.29.0...v1.33.1)
- github.com/prometheus/client_golang: [v1.16.0 → v1.19.1](https://github.com/prometheus/client_golang/compare/v1.16.0...v1.19.1)
- github.com/prometheus/client_model: [v0.4.0 → v0.6.1](https://github.com/prometheus/client_model/compare/v0.4.0...v0.6.1)
- github.com/prometheus/common: [v0.44.0 → v0.53.0](https://github.com/prometheus/common/compare/v0.44.0...v0.53.0)
- github.com/prometheus/procfs: [v0.10.1 → v0.14.0](https://github.com/prometheus/procfs/compare/v0.10.1...v0.14.0)
- github.com/rogpeppe/go-internal: [v1.10.0 → v1.11.0](https://github.com/rogpeppe/go-internal/compare/v1.10.0...v1.11.0)
- github.com/stretchr/objx: [v0.5.0 → v0.1.0](https://github.com/stretchr/objx/compare/v0.5.0...v0.1.0)
- github.com/stretchr/testify: [v1.8.4 → v1.9.0](https://github.com/stretchr/testify/compare/v1.8.4...v1.9.0)
- go.opentelemetry.io/contrib/instrumentation/google.golang.org/grpc/otelgrpc: v0.46.0 → v0.51.0
- go.opentelemetry.io/otel/metric: v1.20.0 → v1.26.0
- go.opentelemetry.io/otel/trace: v1.20.0 → v1.26.0
- go.opentelemetry.io/otel: v1.20.0 → v1.26.0
- go.uber.org/zap: v1.19.0 → v1.26.0
- golang.org/x/crypto: v0.15.0 → v0.23.0
- golang.org/x/mod: v0.8.0 → v0.17.0
- golang.org/x/net: v0.18.0 → v0.25.0
- golang.org/x/oauth2: v0.13.0 → v0.18.0
- golang.org/x/sync: v0.4.0 → v0.7.0
- golang.org/x/sys: v0.15.0 → v0.20.0
- golang.org/x/term: v0.14.0 → v0.20.0
- golang.org/x/text: v0.14.0 → v0.15.0
- golang.org/x/tools: v0.12.0 → v0.20.0
- golang.org/x/xerrors: 04be3eb → 5ec99f8
- google.golang.org/genproto/googleapis/api: d307bd8 → 6ceb2ff
- google.golang.org/genproto/googleapis/rpc: bbf56f3 → 6275950
- google.golang.org/genproto: d783a09 → 6ceb2ff
- google.golang.org/grpc: v1.60.1 → v1.63.2
- google.golang.org/protobuf: v1.31.0 → v1.34.1
- k8s.io/api: v0.29.0 → v0.30.0
- k8s.io/apimachinery: v0.29.0 → v0.30.0
- k8s.io/apiserver: v0.29.0 → v0.30.0
- k8s.io/client-go: v0.29.0 → v0.30.0
- k8s.io/component-base: v0.29.0 → v0.30.0
- k8s.io/cri-api: v0.29.0 → v0.30.0
- k8s.io/klog/v2: v2.110.1 → v2.120.1
- k8s.io/kube-openapi: 2dd684a → f0e62f9
- k8s.io/kubelet: v0.29.0 → v0.30.0

### Removed
- github.com/cncf/udpa/go: [c52dc94](https://github.com/cncf/udpa/go/tree/c52dc94)
- github.com/creack/pty: [v1.1.9](https://github.com/creack/pty/tree/v1.1.9)
- github.com/kr/pty: [v1.1.1](https://github.com/kr/pty/tree/v1.1.1)
- go.uber.org/atomic: v1.10.0
- k8s.io/gengo: 9cce18d
