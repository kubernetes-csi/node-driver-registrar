FROM gcr.io/distroless/static:latest
LABEL maintainers="Kubernetes Authors"
LABEL description="CSI Node driver registrar"
ARG binary=./bin/csi-node-driver-registrar

COPY ${binary} csi-node-driver-registrar
ENTRYPOINT ["/csi-node-driver-registrar"]
