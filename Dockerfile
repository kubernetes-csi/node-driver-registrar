#export DOCKER_CLI_EXPERIMENTAL=enabled
#docker run --rm --privileged docker/binfmt:66f9012c56a8316f9244ffd7622d7c21c1f6f28d
#docker buildx create --use --name mybuilder
#docker buildx build -t colek42/csi-node-driver-registrar --platform=linux/arm,linux/arm64,linux/amd64 . --push

FROM golang:1.13 as builder
WORKDIR /app
ADD . /app/
RUN make

FROM gcr.io/distroless/static:latest
LABEL maintainers="Kubernetes Authors"
LABEL description="CSI Node driver registrar"

COPY --from=builder /app/bin/csi-node-driver-registrar csi-node-driver-registrar
ENTRYPOINT ["/csi-node-driver-registrar"]
