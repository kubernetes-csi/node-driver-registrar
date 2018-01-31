FROM alpine
LABEL maintainers="Kubernetes Authors"
LABEL description="CSI Driver registrar"

COPY driver-registrar driver-registrar
ENTRYPOINT ["/driver-registrar"]
