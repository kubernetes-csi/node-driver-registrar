FROM golang:alpine
LABEL maintainers="Kubernetes Authors"
LABEL description="CSI Driver Registrar"

WORKDIR /go/src/github.com/kubernetes-csi/driver-registrar
COPY . .
RUN cd cmd/driver-registrar && \
    go install
