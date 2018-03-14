FROM alpine
LABEL maintainers="Kubernetes Authors"
LABEL description="CSI Driver registrar"

COPY ./bin/driver-registrar driver-registrar
ENTRYPOINT ["/driver-registrar"]
