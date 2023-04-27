# Release notes for v2.8.0

[Documentation](https://kubernetes-csi.github.io)

# Changelog since v2.7.0

## Dependencies

### Added
- cloud.google.com/go/compute/metadata: v0.2.3
- cloud.google.com/go/compute: v1.15.1
- cloud.google.com/go/errorreporting: v0.3.0
- cloud.google.com/go/firestore: v1.9.0
- cloud.google.com/go/logging: v1.6.1
- cloud.google.com/go/maps: v0.1.0
- cloud.google.com/go/pubsublite: v1.5.0
- cloud.google.com/go/spanner: v1.41.0
- cloud.google.com/go/vmwareengine: v0.1.0
- github.com/go-task/slim-sprig: [348f09d](https://github.com/go-task/slim-sprig/tree/348f09d)

### Changed
- cloud.google.com/go/aiplatform: v1.24.0 → v1.27.0
- cloud.google.com/go/bigquery: v1.43.0 → v1.44.0
- cloud.google.com/go/datastore: v1.1.0 → v1.10.0
- cloud.google.com/go/iam: v0.7.0 → v0.8.0
- cloud.google.com/go/pubsub: v1.3.1 → v1.27.1
- github.com/census-instrumentation/opencensus-proto: [v0.2.1 → v0.4.1](https://github.com/census-instrumentation/opencensus-proto/compare/v0.2.1...v0.4.1)
- github.com/cespare/xxhash/v2: [v2.1.2 → v2.2.0](https://github.com/cespare/xxhash/v2/compare/v2.1.2...v2.2.0)
- github.com/cncf/udpa/go: [04548b0 → c52dc94](https://github.com/cncf/udpa/go/compare/04548b0...c52dc94)
- github.com/cncf/xds/go: [cb28da3 → 06c439d](https://github.com/cncf/xds/go/compare/cb28da3...06c439d)
- github.com/envoyproxy/go-control-plane: [49ff273 → v0.10.3](https://github.com/envoyproxy/go-control-plane/compare/49ff273...v0.10.3)
- github.com/envoyproxy/protoc-gen-validate: [v0.1.0 → v0.9.1](https://github.com/envoyproxy/protoc-gen-validate/compare/v0.1.0...v0.9.1)
- github.com/go-openapi/jsonpointer: [v0.19.5 → v0.19.6](https://github.com/go-openapi/jsonpointer/compare/v0.19.5...v0.19.6)
- github.com/go-openapi/jsonreference: [v0.20.0 → v0.20.1](https://github.com/go-openapi/jsonreference/compare/v0.20.0...v0.20.1)
- github.com/go-openapi/swag: [v0.19.14 → v0.22.3](https://github.com/go-openapi/swag/compare/v0.19.14...v0.22.3)
- github.com/golang/glog: [23def4e → v1.0.0](https://github.com/golang/glog/compare/23def4e...v1.0.0)
- github.com/golang/protobuf: [v1.5.2 → v1.5.3](https://github.com/golang/protobuf/compare/v1.5.2...v1.5.3)
- github.com/google/pprof: [1a94d86 → 4bb14d4](https://github.com/google/pprof/compare/1a94d86...4bb14d4)
- github.com/google/uuid: [v1.1.2 → v1.3.0](https://github.com/google/uuid/compare/v1.1.2...v1.3.0)
- github.com/kr/pretty: [v0.2.0 → v0.3.0](https://github.com/kr/pretty/compare/v0.2.0...v0.3.0)
- github.com/kubernetes-csi/csi-lib-utils: [v0.12.0 → v0.13.0](https://github.com/kubernetes-csi/csi-lib-utils/compare/v0.12.0...v0.13.0)
- github.com/mailru/easyjson: [v0.7.6 → v0.7.7](https://github.com/mailru/easyjson/compare/v0.7.6...v0.7.7)
- github.com/moby/term: [39b0c02 → 1aeaba8](https://github.com/moby/term/compare/39b0c02...1aeaba8)
- github.com/onsi/ginkgo/v2: [v2.4.0 → v2.9.1](https://github.com/onsi/ginkgo/v2/compare/v2.4.0...v2.9.1)
- github.com/onsi/gomega: [v1.23.0 → v1.27.4](https://github.com/onsi/gomega/compare/v1.23.0...v1.27.4)
- github.com/rogpeppe/go-internal: [v1.3.0 → v1.10.0](https://github.com/rogpeppe/go-internal/compare/v1.3.0...v1.10.0)
- github.com/stretchr/objx: [v0.1.1 → v0.5.0](https://github.com/stretchr/objx/compare/v0.1.1...v0.5.0)
- github.com/stretchr/testify: [v1.8.0 → v1.8.1](https://github.com/stretchr/testify/compare/v1.8.0...v1.8.1)
- go.opentelemetry.io/contrib/instrumentation/net/http/otelhttp: v0.35.0 → v0.35.1
- go.uber.org/goleak: v1.2.0 → v1.2.1
- golang.org/x/mod: 86c51ed → v0.8.0
- golang.org/x/net: v0.4.0 → v0.8.0
- golang.org/x/oauth2: ee48083 → v0.4.0
- golang.org/x/sys: v0.3.0 → v0.7.0
- golang.org/x/term: v0.3.0 → v0.6.0
- golang.org/x/text: v0.5.0 → v0.8.0
- golang.org/x/tools: v0.1.12 → v0.7.0
- golang.org/x/xerrors: 5ec99f8 → 04be3eb
- google.golang.org/genproto: 142d8a6 → 76db087
- google.golang.org/grpc: v1.51.0 → v1.54.0
- gopkg.in/check.v1: 8fa4692 → 10cb982
- k8s.io/api: v0.26.0 → v0.27.1
- k8s.io/apimachinery: v0.26.0 → v0.27.1
- k8s.io/client-go: v0.26.0 → v0.27.1
- k8s.io/component-base: v0.26.0 → v0.27.1
- k8s.io/klog/v2: v2.80.1 → v2.90.1
- k8s.io/kube-openapi: 172d655 → 15aac26
- k8s.io/kubelet: v0.26.0 → v0.27.1
- k8s.io/utils: 1a15be2 → a36077c
- sigs.k8s.io/json: f223a00 → bc3834c

### Removed
- github.com/PuerkitoBio/purell: [v1.1.1](https://github.com/PuerkitoBio/purell/tree/v1.1.1)
- github.com/PuerkitoBio/urlesc: [de5bf2a](https://github.com/PuerkitoBio/urlesc/tree/de5bf2a)
- github.com/elazarl/goproxy: [947c36d](https://github.com/elazarl/goproxy/tree/947c36d)
- github.com/niemeyer/pretty: [a10e7ca](https://github.com/niemeyer/pretty/tree/a10e7ca)
- gotest.tools/v3: v3.0.3
com/go-logr/stdr: [v1.2.2](https://github.com/go-logr/stdr/tree/v1.2.2)
- github.com/grpc-ecosystem/grpc-gateway/v2: [v2.7.0](https://github.com/grpc-ecosystem/grpc-gateway/v2/tree/v2.7.0)
- go.opentelemetry.io/otel/exporters/otlp/internal/retry: v1.10.0
- go.opentelemetry.io/otel/exporters/otlp/otlptrace/otlptracegrpc: v1.10.0
- go.opentelemetry.io/otel/exporters/otlp/otlptrace: v1.10.0

### Changed
- cloud.google.com/go/bigquery: v1.8.0 → v1.43.0
- cloud.google.com/go: v0.97.0 → v0.105.0
- github.com/container-storage-interface/spec: [v1.6.0 → v1.7.0](https://github.com/container-storage-interface/spec/compare/v1.6.0...v1.7.0)
- github.com/creack/pty: [v1.1.11 → v1.1.9](https://github.com/creack/pty/compare/v1.1.11...v1.1.9)
- github.com/emicklei/go-restful/v3: [v3.8.0 → v3.9.0](https://github.com/emicklei/go-restful/v3/compare/v3.8.0...v3.9.0)
- github.com/felixge/httpsnoop: [v1.0.1 → v1.0.3](https://github.com/felixge/httpsnoop/compare/v1.0.1...v1.0.3)
- github.com/go-kit/log: [v0.1.0 → v0.2.0](https://github.com/go-kit/log/compare/v0.1.0...v0.2.0)
- github.com/go-logfmt/logfmt: [v0.5.0 → v0.5.1](https://github.com/go-logfmt/logfmt/compare/v0.5.0...v0.5.1)
- github.com/go-openapi/jsonreference: [v0.19.5 → v0.20.0](https://github.com/go-openapi/jsonreference/compare/v0.19.5...v0.20.0)
- github.com/golang/mock: [v1.6.0 → v1.4.4](https://github.com/golang/mock/compare/v1.6.0...v1.4.4)
- github.com/google/go-cmp: [v0.5.8 → v0.5.9](https://github.com/google/go-cmp/compare/v0.5.8...v0.5.9)
- github.com/google/martian/v3: [v3.2.1 → v3.0.0](https://github.com/google/martian/v3/compare/v3.2.1...v3.0.0)
- github.com/google/pprof: [4bb14d4 → 1a94d86](https://github.com/google/pprof/compare/4bb14d4...1a94d86)
- github.com/googleapis/gax-go/v2: [v2.1.0 → v2.0.5](https://github.com/googleapis/gax-go/v2/compare/v2.1.0...v2.0.5)
- github.com/ianlancetaylor/demangle: [28f6c0f → 5e5cf60](https://github.com/ianlancetaylor/demangle/compare/28f6c0f...5e5cf60)
- github.com/inconshreveable/mousetrap: [v1.0.0 → v1.0.1](https://github.com/inconshreveable/mousetrap/compare/v1.0.0...v1.0.1)
- github.com/kubernetes-csi/csi-lib-utils: [v0.11.0 → v0.12.0](https://github.com/kubernetes-csi/csi-lib-utils/compare/v0.11.0...v0.12.0)
- github.com/matttproud/golang_protobuf_extensions: [v1.0.1 → v1.0.4](https://github.com/matttproud/golang_protobuf_extensions/compare/v1.0.1...v1.0.4)
- github.com/moby/term: [3f7ff69 → 39b0c02](https://github.com/moby/term/compare/3f7ff69...39b0c02)
- github.com/onsi/ginkgo/v2: [v2.1.6 → v2.4.0](https://github.com/onsi/ginkgo/v2/compare/v2.1.6...v2.4.0)
- github.com/onsi/gomega: [v1.20.1 → v1.23.0](https://github.com/onsi/gomega/compare/v1.20.1...v1.23.0)
- github.com/prometheus/client_golang: [v1.12.1 → v1.14.0](https://github.com/prometheus/client_golang/compare/v1.12.1...v1.14.0)
- github.com/prometheus/client_model: [v0.2.0 → v0.3.0](https://github.com/prometheus/client_model/compare/v0.2.0...v0.3.0)
- github.com/prometheus/common: [v0.32.1 → v0.37.0](https://github.com/prometheus/common/compare/v0.32.1...v0.37.0)
- github.com/prometheus/procfs: [v0.7.3 → v0.8.0](https://github.com/prometheus/procfs/compare/v0.7.3...v0.8.0)
- github.com/spf13/cobra: [v1.4.0 → v1.6.0](https://github.com/spf13/cobra/compare/v1.4.0...v1.6.0)
- github.com/stretchr/testify: [v1.7.0 → v1.8.0](https://github.com/stretchr/testify/compare/v1.7.0...v1.8.0)
- github.com/yuin/goldmark: [v1.4.13 → v1.2.1](https://github.com/yuin/goldmark/compare/v1.4.13...v1.2.1)
- go.opencensus.io: v0.23.0 → v0.22.4
- go.opentelemetry.io/contrib/instrumentation/net/http/otelhttp: v0.20.0 → v0.35.0
- go.opentelemetry.io/otel/metric: v0.20.0 → v0.31.0
- go.opentelemetry.io/otel/sdk: v0.20.0 → v1.10.0
- go.opentelemetry.io/otel/trace: v0.20.0 → v1.10.0
- go.opentelemetry.io/otel: v0.20.0 → v1.10.0
- go.opentelemetry.io/proto/otlp: v0.7.0 → v0.19.0
- go.uber.org/goleak: v1.1.10 → v1.2.0
- golang.org/x/crypto: 3147a52 → 75b2880
- golang.org/x/lint: 6edffad → 738671d
- golang.org/x/net: a158d28 → v0.4.0
- golang.org/x/oauth2: d3ed0bb → ee48083
- golang.org/x/sync: 886fb93 → 0de741c
- golang.org/x/sys: 8c9f86f → v0.3.0
- golang.org/x/term: 03fcf44 → v0.3.0
- golang.org/x/text: v0.3.7 → v0.5.0
- google.golang.org/api: v0.57.0 → v0.30.0
- google.golang.org/genproto: c8bf987 → 142d8a6
- google.golang.org/grpc: v1.50.1 → v1.51.0
- google.golang.org/protobuf: v1.28.0 → v1.28.1
- k8s.io/api: v0.25.2 → v0.26.0
- k8s.io/apimachinery: v0.25.2 → v0.26.0
- k8s.io/client-go: v0.25.2 → v0.26.0
- k8s.io/component-base: v0.25.2 → v0.26.0
- k8s.io/kube-openapi: 67bda5d → 172d655
- k8s.io/kubelet: v0.25.2 → v0.26.0
- k8s.io/utils: ee6ede2 → 1a15be2
- sigs.k8s.io/yaml: v1.2.0 → v1.3.0

### Removed
- github.com/Azure/go-autorest/autorest/adal: [v0.9.20](https://github.com/Azure/go-autorest/autorest/adal/tree/v0.9.20)
- github.com/Azure/go-autorest/autorest/date: [v0.3.0](https://github.com/Azure/go-autorest/autorest/date/tree/v0.3.0)
- github.com/Azure/go-autorest/autorest/mocks: [v0.4.2](https://github.com/Azure/go-autorest/autorest/mocks/tree/v0.4.2)
- github.com/Azure/go-autorest/autorest: [v0.11.27](https://github.com/Azure/go-autorest/autorest/tree/v0.11.27)
- github.com/Azure/go-autorest/logger: [v0.2.1](https://github.com/Azure/go-autorest/logger/tree/v0.2.1)
- github.com/Azure/go-autorest/tracing: [v0.6.0](https://github.com/Azure/go-autorest/tracing/tree/v0.6.0)
- github.com/Azure/go-autorest: [v14.2.0+incompatible](https://github.com/Azure/go-autorest/tree/v14.2.0)
- github.com/OneOfOne/xxhash: [v1.2.2](https://github.com/OneOfOne/xxhash/tree/v1.2.2)
- github.com/antihax/optional: [v1.0.0](https://github.com/antihax/optional/tree/v1.0.0)
- github.com/benbjohnson/clock: [v1.1.0](https://github.com/benbjohnson/clock/tree/v1.1.0)
- github.com/cespare/xxhash: [v1.1.0](https://github.com/cespare/xxhash/tree/v1.1.0)
- github.com/cpuguy83/go-md2man/v2: [v2.0.1](https://github.com/cpuguy83/go-md2man/v2/tree/v2.0.1)
- github.com/fsnotify/fsnotify: [v1.4.9](https://github.com/fsnotify/fsnotify/tree/v1.4.9)
- github.com/getkin/kin-openapi: [v0.76.0](https://github.com/getkin/kin-openapi/tree/v0.76.0)
- github.com/ghodss/yaml: [v1.0.0](https://github.com/ghodss/yaml/tree/v1.0.0)
- github.com/go-task/slim-sprig: [348f09d](https://github.com/go-task/slim-sprig/tree/348f09d)
- github.com/golang-jwt/jwt/v4: [v4.2.0](https://github.com/golang-jwt/jwt/v4/tree/v4.2.0)
- github.com/golang/snappy: [v0.0.3](https://github.com/golang/snappy/tree/v0.0.3)
- github.com/gorilla/mux: [v1.8.0](https://github.com/gorilla/mux/tree/v1.8.0)
- github.com/gorilla/websocket: [v1.4.2](https://github.com/gorilla/websocket/tree/v1.4.2)
- github.com/grpc-ecosystem/grpc-gateway: [v1.16.0](https://github.com/grpc-ecosystem/grpc-gateway/tree/v1.16.0)
- github.com/hpcloud/tail: [v1.0.0](https://github.com/hpcloud/tail/tree/v1.0.0)
- github.com/nxadm/tail: [v1.4.8](https://github.com/nxadm/tail/tree/v1.4.8)
- github.com/onsi/ginkgo: [v1.16.4](https://github.com/onsi/ginkgo/tree/v1.16.4)
- github.com/rogpeppe/fastuuid: [v1.2.0](https://github.com/rogpeppe/fastuuid/tree/v1.2.0)
- github.com/russross/blackfriday/v2: [v2.1.0](https://github.com/russross/blackfriday/v2/tree/v2.1.0)
- github.com/spaolacci/murmur3: [f09979e](https://github.com/spaolacci/murmur3/tree/f09979e)
- github.com/spf13/afero: [v1.2.2](https://github.com/spf13/afero/tree/v1.2.2)
- go.opentelemetry.io/contrib: v0.20.0
- go.opentelemetry.io/otel/exporters/otlp: v0.20.0
- go.opentelemetry.io/otel/oteltest: v0.20.0
- go.opentelemetry.io/otel/sdk/export/metric: v0.20.0
- go.opentelemetry.io/otel/sdk/metric: v0.20.0
- google.golang.org/grpc/cmd/protoc-gen-go-grpc: v1.1.0
- gopkg.in/fsnotify.v1: v1.4.7
- gopkg.in/tomb.v1: dd63297
