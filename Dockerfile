FROM --platform=$BUILDPLATFORM golang:1.13.3 AS builder
ARG TARGETPLATFORM
ARG BUILDPLATFORM
WORKDIR /code
ADD . /code/

RUN cd /code/ && GOARCH=$(echo $TARGETPLATFORM | cut -f2 -d '/') make build

FROM gcr.io/distroless/static:latest
LABEL maintainers="Kubernetes Authors"
LABEL description="CSI Node driver registrar"

COPY --from=builder /code/bin/csi-node-driver-registrar csi-node-driver-registrar
