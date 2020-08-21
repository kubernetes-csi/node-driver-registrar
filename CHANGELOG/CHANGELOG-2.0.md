# Release notes for v2.0

[Documentation](https://kubernetes-csi.github.io/docs/)
# Changelog since v1.3.0

## Urgent Upgrade Notes 

### (No, really, you MUST read this before you upgrade)

- Use v1 API for plugin registration with kubelet. This change requires at least Kubernetes v1.13. ([#96](https://github.com/kubernetes-csi/node-driver-registrar/pull/96), [@gnufied](https://github.com/gnufied))
 
## Changes by Kind

### Feature

- Windows and multiarch Linux images are now available. The Dockerfile must accept a "binary" build argument. ([#84](https://github.com/kubernetes-csi/node-driver-registrar/pull/84), [@pohly](https://github.com/pohly))

### Uncategorized

- Make image, tag and registry configurable in docker file for windows build of node driver registrar. ([#77](https://github.com/kubernetes-csi/node-driver-registrar/pull/77), [@jingxu97](https://github.com/jingxu97))
- Use v1 API for plugin registration with kubelet ([#96](https://github.com/kubernetes-csi/node-driver-registrar/pull/96), [@gnufied](https://github.com/gnufied))
- Registration socket is now removed on node-driver-registrar shutdown, de-registering the CSI driver from kubelet. ([#61](https://github.com/kubernetes-csi/node-driver-registrar/pull/61), [@Madhu-1](https://github.com/Madhu-1))

## Dependencies

### Added
- cloud.google.com/go/bigquery: v1.0.1
- cloud.google.com/go/datastore: v1.0.0
- cloud.google.com/go/pubsub: v1.0.1
- cloud.google.com/go/storage: v1.0.0
- cloud.google.com/go: v0.51.0
- dmitri.shuralyov.com/gpu/mtl: 666a987
- github.com/Azure/go-ansiterm: [d6e3b33](https://github.com/Azure/go-ansiterm/tree/d6e3b33)
- github.com/Azure/go-autorest/autorest/adal: [v0.8.2](https://github.com/Azure/go-autorest/autorest/adal/tree/v0.8.2)
- github.com/Azure/go-autorest/autorest/date: [v0.2.0](https://github.com/Azure/go-autorest/autorest/date/tree/v0.2.0)
- github.com/Azure/go-autorest/autorest/mocks: [v0.3.0](https://github.com/Azure/go-autorest/autorest/mocks/tree/v0.3.0)
- github.com/Azure/go-autorest/autorest: [v0.9.6](https://github.com/Azure/go-autorest/autorest/tree/v0.9.6)
- github.com/Azure/go-autorest/logger: [v0.1.0](https://github.com/Azure/go-autorest/logger/tree/v0.1.0)
- github.com/Azure/go-autorest/tracing: [v0.5.0](https://github.com/Azure/go-autorest/tracing/tree/v0.5.0)
- github.com/BurntSushi/toml: [v0.3.1](https://github.com/BurntSushi/toml/tree/v0.3.1)
- github.com/BurntSushi/xgb: [27f1227](https://github.com/BurntSushi/xgb/tree/27f1227)
- github.com/NYTimes/gziphandler: [56545f4](https://github.com/NYTimes/gziphandler/tree/56545f4)
- github.com/PuerkitoBio/purell: [v1.0.0](https://github.com/PuerkitoBio/purell/tree/v1.0.0)
- github.com/PuerkitoBio/urlesc: [5bd2802](https://github.com/PuerkitoBio/urlesc/tree/5bd2802)
- github.com/alecthomas/template: [fb15b89](https://github.com/alecthomas/template/tree/fb15b89)
- github.com/alecthomas/units: [c3de453](https://github.com/alecthomas/units/tree/c3de453)
- github.com/beorn7/perks: [v1.0.1](https://github.com/beorn7/perks/tree/v1.0.1)
- github.com/blang/semver: [v3.5.0+incompatible](https://github.com/blang/semver/tree/v3.5.0)
- github.com/census-instrumentation/opencensus-proto: [v0.2.1](https://github.com/census-instrumentation/opencensus-proto/tree/v0.2.1)
- github.com/cespare/xxhash/v2: [v2.1.1](https://github.com/cespare/xxhash/v2/tree/v2.1.1)
- github.com/chzyer/logex: [v1.1.10](https://github.com/chzyer/logex/tree/v1.1.10)
- github.com/chzyer/readline: [2972be2](https://github.com/chzyer/readline/tree/2972be2)
- github.com/chzyer/test: [a1ea475](https://github.com/chzyer/test/tree/a1ea475)
- github.com/client9/misspell: [v0.3.4](https://github.com/client9/misspell/tree/v0.3.4)
- github.com/dgrijalva/jwt-go: [v3.2.0+incompatible](https://github.com/dgrijalva/jwt-go/tree/v3.2.0)
- github.com/docker/spdystream: [449fdfc](https://github.com/docker/spdystream/tree/449fdfc)
- github.com/docopt/docopt-go: [ee0de3b](https://github.com/docopt/docopt-go/tree/ee0de3b)
- github.com/elazarl/goproxy: [947c36d](https://github.com/elazarl/goproxy/tree/947c36d)
- github.com/emicklei/go-restful: [ff4f55a](https://github.com/emicklei/go-restful/tree/ff4f55a)
- github.com/envoyproxy/go-control-plane: [5f8ba28](https://github.com/envoyproxy/go-control-plane/tree/5f8ba28)
- github.com/envoyproxy/protoc-gen-validate: [v0.1.0](https://github.com/envoyproxy/protoc-gen-validate/tree/v0.1.0)
- github.com/evanphx/json-patch: [e83c0a1](https://github.com/evanphx/json-patch/tree/e83c0a1)
- github.com/fsnotify/fsnotify: [v1.4.9](https://github.com/fsnotify/fsnotify/tree/v1.4.9)
- github.com/ghodss/yaml: [73d445a](https://github.com/ghodss/yaml/tree/73d445a)
- github.com/go-gl/glfw/v3.3/glfw: [12ad95a](https://github.com/go-gl/glfw/v3.3/glfw/tree/12ad95a)
- github.com/go-kit/kit: [v0.9.0](https://github.com/go-kit/kit/tree/v0.9.0)
- github.com/go-logfmt/logfmt: [v0.4.0](https://github.com/go-logfmt/logfmt/tree/v0.4.0)
- github.com/go-logr/logr: [v0.2.0](https://github.com/go-logr/logr/tree/v0.2.0)
- github.com/go-openapi/jsonpointer: [46af16f](https://github.com/go-openapi/jsonpointer/tree/46af16f)
- github.com/go-openapi/jsonreference: [13c6e35](https://github.com/go-openapi/jsonreference/tree/13c6e35)
- github.com/go-openapi/spec: [6aced65](https://github.com/go-openapi/spec/tree/6aced65)
- github.com/go-openapi/swag: [1d0bd11](https://github.com/go-openapi/swag/tree/1d0bd11)
- github.com/go-stack/stack: [v1.8.0](https://github.com/go-stack/stack/tree/v1.8.0)
- github.com/golang/groupcache: [215e871](https://github.com/golang/groupcache/tree/215e871)
- github.com/golang/mock: [v1.3.1](https://github.com/golang/mock/tree/v1.3.1)
- github.com/google/btree: [v1.0.0](https://github.com/google/btree/tree/v1.0.0)
- github.com/google/go-cmp: [v0.4.0](https://github.com/google/go-cmp/tree/v0.4.0)
- github.com/google/gofuzz: [v1.1.0](https://github.com/google/gofuzz/tree/v1.1.0)
- github.com/google/martian: [v2.1.0+incompatible](https://github.com/google/martian/tree/v2.1.0)
- github.com/google/pprof: [d4f498a](https://github.com/google/pprof/tree/d4f498a)
- github.com/google/renameio: [v0.1.0](https://github.com/google/renameio/tree/v0.1.0)
- github.com/google/uuid: [v1.1.1](https://github.com/google/uuid/tree/v1.1.1)
- github.com/googleapis/gax-go/v2: [v2.0.5](https://github.com/googleapis/gax-go/v2/tree/v2.0.5)
- github.com/googleapis/gnostic: [v0.4.1](https://github.com/googleapis/gnostic/tree/v0.4.1)
- github.com/gregjones/httpcache: [9cad4c3](https://github.com/gregjones/httpcache/tree/9cad4c3)
- github.com/hashicorp/golang-lru: [v0.5.1](https://github.com/hashicorp/golang-lru/tree/v0.5.1)
- github.com/hpcloud/tail: [v1.0.0](https://github.com/hpcloud/tail/tree/v1.0.0)
- github.com/ianlancetaylor/demangle: [5e5cf60](https://github.com/ianlancetaylor/demangle/tree/5e5cf60)
- github.com/imdario/mergo: [v0.3.5](https://github.com/imdario/mergo/tree/v0.3.5)
- github.com/json-iterator/go: [v1.1.10](https://github.com/json-iterator/go/tree/v1.1.10)
- github.com/jstemmer/go-junit-report: [v0.9.1](https://github.com/jstemmer/go-junit-report/tree/v0.9.1)
- github.com/julienschmidt/httprouter: [v1.2.0](https://github.com/julienschmidt/httprouter/tree/v1.2.0)
- github.com/kisielk/errcheck: [v1.2.0](https://github.com/kisielk/errcheck/tree/v1.2.0)
- github.com/kisielk/gotool: [v1.0.0](https://github.com/kisielk/gotool/tree/v1.0.0)
- github.com/konsorten/go-windows-terminal-sequences: [v1.0.3](https://github.com/konsorten/go-windows-terminal-sequences/tree/v1.0.3)
- github.com/kr/logfmt: [b84e30a](https://github.com/kr/logfmt/tree/b84e30a)
- github.com/kr/pretty: [v0.2.0](https://github.com/kr/pretty/tree/v0.2.0)
- github.com/kr/pty: [v1.1.1](https://github.com/kr/pty/tree/v1.1.1)
- github.com/kr/text: [v0.1.0](https://github.com/kr/text/tree/v0.1.0)
- github.com/mailru/easyjson: [d5b7844](https://github.com/mailru/easyjson/tree/d5b7844)
- github.com/matttproud/golang_protobuf_extensions: [c182aff](https://github.com/matttproud/golang_protobuf_extensions/tree/c182aff)
- github.com/moby/term: [672ec06](https://github.com/moby/term/tree/672ec06)
- github.com/modern-go/concurrent: [bacd9c7](https://github.com/modern-go/concurrent/tree/bacd9c7)
- github.com/modern-go/reflect2: [v1.0.1](https://github.com/modern-go/reflect2/tree/v1.0.1)
- github.com/munnerz/goautoneg: [a547fc6](https://github.com/munnerz/goautoneg/tree/a547fc6)
- github.com/mwitkow/go-conntrack: [cc309e4](https://github.com/mwitkow/go-conntrack/tree/cc309e4)
- github.com/mxk/go-flowrate: [cca7078](https://github.com/mxk/go-flowrate/tree/cca7078)
- github.com/onsi/ginkgo: [v1.11.0](https://github.com/onsi/ginkgo/tree/v1.11.0)
- github.com/onsi/gomega: [v1.7.0](https://github.com/onsi/gomega/tree/v1.7.0)
- github.com/peterbourgon/diskv: [v2.0.1+incompatible](https://github.com/peterbourgon/diskv/tree/v2.0.1)
- github.com/pkg/errors: [v0.9.1](https://github.com/pkg/errors/tree/v0.9.1)
- github.com/prometheus/client_golang: [v1.7.1](https://github.com/prometheus/client_golang/tree/v1.7.1)
- github.com/prometheus/client_model: [v0.2.0](https://github.com/prometheus/client_model/tree/v0.2.0)
- github.com/prometheus/common: [v0.10.0](https://github.com/prometheus/common/tree/v0.10.0)
- github.com/prometheus/procfs: [v0.1.3](https://github.com/prometheus/procfs/tree/v0.1.3)
- github.com/rogpeppe/go-internal: [v1.3.0](https://github.com/rogpeppe/go-internal/tree/v1.3.0)
- github.com/sirupsen/logrus: [v1.6.0](https://github.com/sirupsen/logrus/tree/v1.6.0)
- github.com/spf13/afero: [v1.2.2](https://github.com/spf13/afero/tree/v1.2.2)
- github.com/spf13/pflag: [v1.0.5](https://github.com/spf13/pflag/tree/v1.0.5)
- go.opencensus.io: v0.22.2
- go.uber.org/atomic: v1.4.0
- go.uber.org/multierr: v1.1.0
- go.uber.org/zap: v1.10.0
- golang.org/x/crypto: bac4c82
- golang.org/x/exp: da58074
- golang.org/x/image: cff245a
- golang.org/x/lint: fdd1cda
- golang.org/x/mobile: d2bd2a2
- golang.org/x/mod: c90efee
- golang.org/x/oauth2: 858c2ad
- golang.org/x/time: 555d28b
- golang.org/x/tools: 7b8e75d
- golang.org/x/xerrors: 9bdfabe
- google.golang.org/api: v0.15.0
- google.golang.org/appengine: v1.6.5
- google.golang.org/protobuf: v1.24.0
- gopkg.in/alecthomas/kingpin.v2: v2.2.6
- gopkg.in/errgo.v2: v2.1.0
- gopkg.in/fsnotify.v1: v1.4.7
- gopkg.in/inf.v0: v0.9.1
- gopkg.in/tomb.v1: dd63297
- gotest.tools/v3: v3.0.2
- gotest.tools: v2.2.0+incompatible
- honnef.co/go/tools: v0.0.1-2019.2.3
- k8s.io/api: v0.19.0-rc.2
- k8s.io/apimachinery: v0.19.0-rc.2
- k8s.io/component-base: v0.19.0-rc.2
- k8s.io/gengo: 3a45101
- k8s.io/klog/v2: v2.2.0
- k8s.io/kube-openapi: 656914f
- k8s.io/kubelet: v0.19.0-rc.2
- k8s.io/utils: 0bdb4ca
- rsc.io/binaryregexp: v0.2.0
- sigs.k8s.io/structured-merge-diff/v3: 43c19bb
- sigs.k8s.io/yaml: v1.2.0

### Changed
- github.com/davecgh/go-spew: [v1.1.0 → v1.1.1](https://github.com/davecgh/go-spew/compare/v1.1.0...v1.1.1)
- github.com/gogo/protobuf: [v1.0.0 → v1.3.1](https://github.com/gogo/protobuf/compare/v1.0.0...v1.3.1)
- github.com/golang/protobuf: [v1.1.0 → v1.4.2](https://github.com/golang/protobuf/compare/v1.1.0...v1.4.2)
- github.com/stretchr/objx: [v0.1.0 → v0.1.1](https://github.com/stretchr/objx/compare/v0.1.0...v0.1.1)
- golang.org/x/net: 22ae77b → d3edc99
- golang.org/x/sys: dd2ff4a → ed371f2
- golang.org/x/text: v0.3.0 → v0.3.3
- google.golang.org/genproto: 2c5e7ac → cb27e3a
- google.golang.org/grpc: v1.10.0 → v1.27.0
- gopkg.in/check.v1: 20d25e2 → 41f04d3
- gopkg.in/yaml.v2: v2.2.2 → v2.2.8
- k8s.io/klog: v0.1.0 → v1.0.0

### Removed
- k8s.io/kubernetes: v1.11.0-beta.2
