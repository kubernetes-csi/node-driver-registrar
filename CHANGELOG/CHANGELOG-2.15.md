# Release notes for v2.15.0

[Documentation](https:/kubernetes-csi.github.io)

# Changelog since v2.14.0

## Changes by Kind

### Other (Cleanup or Flake)

- Update kubernetes dependencies to v1.34.0 ([#529](https://github.com/kubernetes-csi/node-driver-registrar/pull/529), [@dobsonj](https://github.com/dobsonj))

## Dependencies

### Added
- go.yaml.in/yaml/v2: v2.4.2
- go.yaml.in/yaml/v3: v3.0.4
- sigs.k8s.io/structured-merge-diff/v6: v6.3.0

### Changed
- github.com/fxamacker/cbor/v2: [v2.8.0 → v2.9.0](https://github.com/fxamacker/cbor/compare/v2.8.0...v2.9.0)
- github.com/google/gnostic-models: [v0.6.9 → v0.7.0](https://github.com/google/gnostic-models/compare/v0.6.9...v0.7.0)
- github.com/grpc-ecosystem/grpc-gateway/v2: [v2.24.0 → v2.26.3](https://github.com/grpc-ecosystem/grpc-gateway/compare/v2.24.0...v2.26.3)
- github.com/modern-go/reflect2: [v1.0.2 → 35a7c28](https://github.com/modern-go/reflect2/compare/v1.0.2...35a7c28)
- go.opentelemetry.io/otel/exporters/otlp/otlptrace/otlptracegrpc: v1.33.0 → v1.34.0
- go.opentelemetry.io/otel/exporters/otlp/otlptrace: v1.33.0 → v1.34.0
- go.opentelemetry.io/proto/otlp: v1.4.0 → v1.5.0
- google.golang.org/genproto/googleapis/api: 56aae31 → a0af3ef
- google.golang.org/genproto/googleapis/rpc: 56aae31 → a0af3ef
- k8s.io/api: v0.33.0 → v0.34.0
- k8s.io/apimachinery: v0.33.0 → v0.34.0
- k8s.io/apiserver: v0.33.0 → v0.34.0
- k8s.io/client-go: v0.33.0 → v0.34.0
- k8s.io/component-base: v0.33.0 → v0.34.0
- k8s.io/cri-api: v0.33.0 → v0.34.0
- k8s.io/gengo/v2: a7b603a → 85fd79d
- k8s.io/kube-openapi: c8a335a → f3f2b99
- k8s.io/kubelet: v0.33.0 → v0.34.0
- k8s.io/utils: 24370be → 4c0f3b2
- sigs.k8s.io/structured-merge-diff/v4: v4.7.0 → v4.6.0
- sigs.k8s.io/yaml: v1.4.0 → v1.6.0

### Removed
_Nothing has changed._
