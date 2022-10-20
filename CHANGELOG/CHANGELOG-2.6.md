# Release notes for v2.6.0

[Documentation](https://kubernetes-csi.github.io)

# Changelog since v2.5.0

## Changes by Kind

### Bug or Regression

- Corrects the handler used to serve http requests, the /healthz and /debug/ endpoints are available now. ([#190](https://github.com/kubernetes-csi/node-driver-registrar/pull/190), [@cnmcavoy](https://github.com/cnmcavoy))

### Other (Cleanup or Flake)

- The directories in `kubeletRegistrationPath` are assumed to be already created and are no longer created by node-driver-registrar ([#214](https://github.com/kubernetes-csi/node-driver-registrar/pull/214), [@mauriciopoppe](https://github.com/mauriciopoppe))
- Update dependencies to v1.25.2 ([#217](https://github.com/kubernetes-csi/node-driver-registrar/pull/217), [#225](https://github.com/kubernetes-csi/node-driver-registrar/pull/225), [#227](https://github.com/kubernetes-csi/node-driver-registrar/pull/227), [@humblec](https://github.com/humblec), [@mauriciopoppe](https://github.com/mauriciopoppe))

## Dependencies

### Added
- github.com/armon/go-socks5: [e753329](https://github.com/armon/go-socks5/tree/e753329)
- github.com/blang/semver/v4: [v4.0.0](https://github.com/blang/semver/v4/tree/v4.0.0)
- github.com/emicklei/go-restful/v3: [v3.8.0](https://github.com/emicklei/go-restful/v3/tree/v3.8.0)
- github.com/go-task/slim-sprig: [348f09d](https://github.com/go-task/slim-sprig/tree/348f09d)
- github.com/golang-jwt/jwt/v4: [v4.2.0](https://github.com/golang-jwt/jwt/v4/tree/v4.2.0)
- github.com/golang/snappy: [v0.0.3](https://github.com/golang/snappy/tree/v0.0.3)
- github.com/google/gnostic: [v0.5.7-v3refs](https://github.com/google/gnostic/tree/v0.5.7-v3refs)
- github.com/josharian/intern: [v1.0.0](https://github.com/josharian/intern/tree/v1.0.0)
- github.com/onsi/ginkgo/v2: [v2.1.6](https://github.com/onsi/ginkgo/v2/tree/v2.1.6)
- google.golang.org/grpc/cmd/protoc-gen-go-grpc: v1.1.0

### Changed
- cloud.google.com/go: v0.81.0 → v0.97.0
- github.com/Azure/go-autorest/autorest/adal: [v0.9.13 → v0.9.20](https://github.com/Azure/go-autorest/autorest/adal/compare/v0.9.13...v0.9.20)
- github.com/Azure/go-autorest/autorest/mocks: [v0.4.1 → v0.4.2](https://github.com/Azure/go-autorest/autorest/mocks/compare/v0.4.1...v0.4.2)
- github.com/Azure/go-autorest/autorest: [v0.11.18 → v0.11.27](https://github.com/Azure/go-autorest/autorest/compare/v0.11.18...v0.11.27)
- github.com/cespare/xxhash/v2: [v2.1.1 → v2.1.2](https://github.com/cespare/xxhash/v2/compare/v2.1.1...v2.1.2)
- github.com/cncf/udpa/go: [5459f2c → 04548b0](https://github.com/cncf/udpa/go/compare/5459f2c...04548b0)
- github.com/cncf/xds/go: [fbca930 → cb28da3](https://github.com/cncf/xds/go/compare/fbca930...cb28da3)
- github.com/container-storage-interface/spec: [v1.5.0 → v1.6.0](https://github.com/container-storage-interface/spec/compare/v1.5.0...v1.6.0)
- github.com/cpuguy83/go-md2man/v2: [v2.0.0 → v2.0.1](https://github.com/cpuguy83/go-md2man/v2/compare/v2.0.0...v2.0.1)
- github.com/envoyproxy/go-control-plane: [63b5d3c → 49ff273](https://github.com/envoyproxy/go-control-plane/compare/63b5d3c...49ff273)
- github.com/go-logr/logr: [v1.2.0 → v1.2.3](https://github.com/go-logr/logr/compare/v1.2.0...v1.2.3)
- github.com/go-logr/zapr: [v1.2.0 → v1.2.3](https://github.com/go-logr/zapr/compare/v1.2.0...v1.2.3)
- github.com/go-openapi/jsonreference: [v0.19.3 → v0.19.5](https://github.com/go-openapi/jsonreference/compare/v0.19.3...v0.19.5)
- github.com/go-openapi/swag: [v0.19.5 → v0.19.14](https://github.com/go-openapi/swag/compare/v0.19.5...v0.19.14)
- github.com/golang/mock: [v1.5.0 → v1.6.0](https://github.com/golang/mock/compare/v1.5.0...v1.6.0)
- github.com/google/go-cmp: [v0.5.5 → v0.5.8](https://github.com/google/go-cmp/compare/v0.5.5...v0.5.8)
- github.com/google/martian/v3: [v3.1.0 → v3.2.1](https://github.com/google/martian/v3/compare/v3.1.0...v3.2.1)
- github.com/google/pprof: [cbba55b → 4bb14d4](https://github.com/google/pprof/compare/cbba55b...4bb14d4)
- github.com/googleapis/gax-go/v2: [v2.0.5 → v2.1.0](https://github.com/googleapis/gax-go/v2/compare/v2.0.5...v2.1.0)
- github.com/imdario/mergo: [v0.3.5 → v0.3.6](https://github.com/imdario/mergo/compare/v0.3.5...v0.3.6)
- github.com/kubernetes-csi/csi-lib-utils: [v0.10.0 → v0.11.0](https://github.com/kubernetes-csi/csi-lib-utils/compare/v0.10.0...v0.11.0)
- github.com/mailru/easyjson: [b2ccc51 → v0.7.6](https://github.com/mailru/easyjson/compare/b2ccc51...v0.7.6)
- github.com/matttproud/golang_protobuf_extensions: [c182aff → v1.0.1](https://github.com/matttproud/golang_protobuf_extensions/compare/c182aff...v1.0.1)
- github.com/mitchellh/mapstructure: [v1.4.1 → v1.1.2](https://github.com/mitchellh/mapstructure/compare/v1.4.1...v1.1.2)
- github.com/moby/term: [9d4ed18 → 3f7ff69](https://github.com/moby/term/compare/9d4ed18...3f7ff69)
- github.com/munnerz/goautoneg: [a547fc6 → a7dc8b6](https://github.com/munnerz/goautoneg/compare/a547fc6...a7dc8b6)
- github.com/nxadm/tail: [v1.4.4 → v1.4.8](https://github.com/nxadm/tail/compare/v1.4.4...v1.4.8)
- github.com/onsi/ginkgo: [v1.14.0 → v1.16.4](https://github.com/onsi/ginkgo/compare/v1.14.0...v1.16.4)
- github.com/onsi/gomega: [v1.10.1 → v1.20.1](https://github.com/onsi/gomega/compare/v1.10.1...v1.20.1)
- github.com/prometheus/client_golang: [v1.11.0 → v1.12.1](https://github.com/prometheus/client_golang/compare/v1.11.0...v1.12.1)
- github.com/prometheus/common: [v0.28.0 → v0.32.1](https://github.com/prometheus/common/compare/v0.28.0...v0.32.1)
- github.com/prometheus/procfs: [v0.6.0 → v0.7.3](https://github.com/prometheus/procfs/compare/v0.6.0...v0.7.3)
- github.com/russross/blackfriday/v2: [v2.0.1 → v2.1.0](https://github.com/russross/blackfriday/v2/compare/v2.0.1...v2.1.0)
- github.com/spf13/afero: [v1.6.0 → v1.2.2](https://github.com/spf13/afero/compare/v1.6.0...v1.2.2)
- github.com/spf13/cobra: [v1.2.1 → v1.4.0](https://github.com/spf13/cobra/compare/v1.2.1...v1.4.0)
- github.com/yuin/goldmark: [v1.4.0 → v1.4.13](https://github.com/yuin/goldmark/compare/v1.4.0...v1.4.13)
- golang.org/x/crypto: 32db794 → 3147a52
- golang.org/x/mod: v0.4.2 → 86c51ed
- golang.org/x/net: e898025 → a158d28
- golang.org/x/oauth2: 2bc19b1 → d3ed0bb
- golang.org/x/sync: 036812b → 886fb93
- golang.org/x/sys: f4d4317 → 8c9f86f
- golang.org/x/term: 6886f2d → 03fcf44
- golang.org/x/time: 1f47c86 → 90d013b
- golang.org/x/tools: d4cc65f → v0.1.12
- google.golang.org/api: v0.44.0 → v0.57.0
- google.golang.org/genproto: fe13028 → c8bf987
- google.golang.org/grpc: v1.40.0 → v1.50.1
- google.golang.org/protobuf: v1.27.1 → v1.28.0
- gopkg.in/yaml.v3: 496545a → v3.0.1
- k8s.io/api: v0.23.0 → v0.25.2
- k8s.io/apimachinery: v0.23.0 → v0.25.2
- k8s.io/client-go: v0.23.0 → v0.25.2
- k8s.io/component-base: v0.23.0 → v0.25.2
- k8s.io/klog/v2: v2.30.0 → v2.80.1
- k8s.io/kube-openapi: e816edb → 67bda5d
- k8s.io/kubelet: v0.23.0 → v0.25.2
- k8s.io/utils: cb0fa31 → ee6ede2
- sigs.k8s.io/json: c049b76 → f223a00
- sigs.k8s.io/structured-merge-diff/v4: v4.1.2 → v4.2.3

### Removed
- cloud.google.com/go/firestore: v1.1.0
- github.com/armon/circbuf: [bbbad09](https://github.com/armon/circbuf/tree/bbbad09)
- github.com/armon/go-metrics: [f0300d1](https://github.com/armon/go-metrics/tree/f0300d1)
- github.com/armon/go-radix: [7fddfc3](https://github.com/armon/go-radix/tree/7fddfc3)
- github.com/bgentry/speakeasy: [v0.1.0](https://github.com/bgentry/speakeasy/tree/v0.1.0)
- github.com/bketelsen/crypt: [v0.0.4](https://github.com/bketelsen/crypt/tree/v0.0.4)
- github.com/blang/semver: [v3.5.1+incompatible](https://github.com/blang/semver/tree/v3.5.1)
- github.com/coreos/go-semver: [v0.3.0](https://github.com/coreos/go-semver/tree/v0.3.0)
- github.com/coreos/go-systemd/v22: [v22.3.2](https://github.com/coreos/go-systemd/v22/tree/v22.3.2)
- github.com/emicklei/go-restful: [ff4f55a](https://github.com/emicklei/go-restful/tree/ff4f55a)
- github.com/fatih/color: [v1.7.0](https://github.com/fatih/color/tree/v1.7.0)
- github.com/form3tech-oss/jwt-go: [v3.2.3+incompatible](https://github.com/form3tech-oss/jwt-go/tree/v3.2.3)
- github.com/godbus/dbus/v5: [v5.0.4](https://github.com/godbus/dbus/v5/tree/v5.0.4)
- github.com/googleapis/gnostic: [v0.5.5](https://github.com/googleapis/gnostic/tree/v0.5.5)
- github.com/gopherjs/gopherjs: [0766667](https://github.com/gopherjs/gopherjs/tree/0766667)
- github.com/hashicorp/consul/api: [v1.1.0](https://github.com/hashicorp/consul/api/tree/v1.1.0)
- github.com/hashicorp/consul/sdk: [v0.1.1](https://github.com/hashicorp/consul/sdk/tree/v0.1.1)
- github.com/hashicorp/errwrap: [v1.0.0](https://github.com/hashicorp/errwrap/tree/v1.0.0)
- github.com/hashicorp/go-cleanhttp: [v0.5.1](https://github.com/hashicorp/go-cleanhttp/tree/v0.5.1)
- github.com/hashicorp/go-immutable-radix: [v1.0.0](https://github.com/hashicorp/go-immutable-radix/tree/v1.0.0)
- github.com/hashicorp/go-msgpack: [v0.5.3](https://github.com/hashicorp/go-msgpack/tree/v0.5.3)
- github.com/hashicorp/go-multierror: [v1.0.0](https://github.com/hashicorp/go-multierror/tree/v1.0.0)
- github.com/hashicorp/go-rootcerts: [v1.0.0](https://github.com/hashicorp/go-rootcerts/tree/v1.0.0)
- github.com/hashicorp/go-sockaddr: [v1.0.0](https://github.com/hashicorp/go-sockaddr/tree/v1.0.0)
- github.com/hashicorp/go-syslog: [v1.0.0](https://github.com/hashicorp/go-syslog/tree/v1.0.0)
- github.com/hashicorp/go-uuid: [v1.0.1](https://github.com/hashicorp/go-uuid/tree/v1.0.1)
- github.com/hashicorp/go.net: [v0.0.1](https://github.com/hashicorp/go.net/tree/v0.0.1)
- github.com/hashicorp/hcl: [v1.0.0](https://github.com/hashicorp/hcl/tree/v1.0.0)
- github.com/hashicorp/logutils: [v1.0.0](https://github.com/hashicorp/logutils/tree/v1.0.0)
- github.com/hashicorp/mdns: [v1.0.0](https://github.com/hashicorp/mdns/tree/v1.0.0)
- github.com/hashicorp/memberlist: [v0.1.3](https://github.com/hashicorp/memberlist/tree/v0.1.3)
- github.com/hashicorp/serf: [v0.8.2](https://github.com/hashicorp/serf/tree/v0.8.2)
- github.com/jtolds/gls: [v4.20.0+incompatible](https://github.com/jtolds/gls/tree/v4.20.0)
- github.com/kr/fs: [v0.1.0](https://github.com/kr/fs/tree/v0.1.0)
- github.com/magiconair/properties: [v1.8.5](https://github.com/magiconair/properties/tree/v1.8.5)
- github.com/mattn/go-colorable: [v0.0.9](https://github.com/mattn/go-colorable/tree/v0.0.9)
- github.com/mattn/go-isatty: [v0.0.3](https://github.com/mattn/go-isatty/tree/v0.0.3)
- github.com/miekg/dns: [v1.0.14](https://github.com/miekg/dns/tree/v1.0.14)
- github.com/mitchellh/cli: [v1.0.0](https://github.com/mitchellh/cli/tree/v1.0.0)
- github.com/mitchellh/go-homedir: [v1.0.0](https://github.com/mitchellh/go-homedir/tree/v1.0.0)
- github.com/mitchellh/go-testing-interface: [v1.0.0](https://github.com/mitchellh/go-testing-interface/tree/v1.0.0)
- github.com/mitchellh/gox: [v0.4.0](https://github.com/mitchellh/gox/tree/v0.4.0)
- github.com/mitchellh/iochan: [v1.0.0](https://github.com/mitchellh/iochan/tree/v1.0.0)
- github.com/pascaldekloe/goe: [57f6aae](https://github.com/pascaldekloe/goe/tree/57f6aae)
- github.com/pelletier/go-toml: [v1.9.3](https://github.com/pelletier/go-toml/tree/v1.9.3)
- github.com/pkg/sftp: [v1.10.1](https://github.com/pkg/sftp/tree/v1.10.1)
- github.com/posener/complete: [v1.1.1](https://github.com/posener/complete/tree/v1.1.1)
- github.com/ryanuber/columnize: [9b3edd6](https://github.com/ryanuber/columnize/tree/9b3edd6)
- github.com/sean-/seed: [e2103e2](https://github.com/sean-/seed/tree/e2103e2)
- github.com/shurcooL/sanitized_anchor_name: [v1.0.0](https://github.com/shurcooL/sanitized_anchor_name/tree/v1.0.0)
- github.com/smartystreets/assertions: [b2de0cb](https://github.com/smartystreets/assertions/tree/b2de0cb)
- github.com/smartystreets/goconvey: [v1.6.4](https://github.com/smartystreets/goconvey/tree/v1.6.4)
- github.com/spf13/cast: [v1.3.1](https://github.com/spf13/cast/tree/v1.3.1)
- github.com/spf13/jwalterweatherman: [v1.1.0](https://github.com/spf13/jwalterweatherman/tree/v1.1.0)
- github.com/spf13/viper: [v1.8.1](https://github.com/spf13/viper/tree/v1.8.1)
- github.com/subosito/gotenv: [v1.2.0](https://github.com/subosito/gotenv/tree/v1.2.0)
- go.etcd.io/etcd/api/v3: v3.5.0
- go.etcd.io/etcd/client/pkg/v3: v3.5.0
- go.etcd.io/etcd/client/v2: v2.305.0
- gopkg.in/ini.v1: v1.62.0
