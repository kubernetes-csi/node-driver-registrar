FROM alpine
LABEL maintainers="Kubernetes Authors"
LABEL description="CSI Node driver registrar"

COPY ./bin/node-driver-registrar node-driver-registrar
ENTRYPOINT ["/node-driver-registrar"]
