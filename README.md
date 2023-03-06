# Node Driver Registrar

The node-driver-registrar is a sidecar container that registers the CSI driver
with Kubelet using the
[kubelet plugin registration mechanism](https://kubernetes.io/docs/concepts/extend-kubernetes/compute-storage-net/device-plugins/#device-plugin-registration).

This is necessary because Kubelet is responsible for issuing CSI `NodeGetInfo`,
`NodeStageVolume`, `NodePublishVolume` calls. The `node-driver-registrar` registers
your CSI driver with Kubelet so that it knows which Unix domain socket to issue
the CSI calls on.

## Compatibility

This information reflects the head of this branch.

| Compatible with CSI Version                                                                | Container Image                                  | [Min K8s Version](https://kubernetes-csi.github.io/docs/kubernetes-compatibility.html#minimum-version) |
| ------------------------------------------------------------------------------------------ | ------------------------------------------------ | ------------------------------------------------------------------------------------------------------ |
| [CSI Spec v1.5.0](https://github.com/container-storage-interface/spec/releases/tag/v1.5.0) | registry.k8s.io/sig-storage/csi-node-driver-registrar | 1.13                                                                                                   |

For release-0.4 and below, please refer to the [driver-registrar
repository](https://github.com/kubernetes-csi/driver-registrar).

## Usage

There are two UNIX domain sockets used by the node-driver-registrar:

* Registration socket:
  * Registers the driver with kubelet.
  * Created by the `node-driver-registrar`.
  * Exposed on a Kubernetes node via hostpath in the Kubelet plugin registry.
    (typically `/var/lib/kubelet/plugins_registry/<drivername.example.com>-reg.sock`).
    The hostpath volume must be mounted at `/registration`.

* CSI driver socket:
  * Used by kubelet to interact with the CSI driver.
  * Created by the CSI driver.
  * Exposed on a Kubernetes node via hostpath somewhere other than the Kubelet plugin registry. (typically `/var/lib/kubelet/plugins/<drivername.example.com>/csi.sock`).
  * This is the socket referenced by the `--csi-address` and `--kubelet-registration-path` arguments.
  * Note that before Kubernetes v1.17, if the csi socket is in the `/var/lib/kubelet/plugins/` path, kubelet may log a lot of harmless errors regarding grpc `GetInfo` call not implemented (fix in kubernetes/kubernetes#84533). The `/var/lib/kubelet/csi-plugins/` path is preferred in Kubernetes versions prior to v1.17.

### Required arguments

* `--csi-address`: This is the path to the CSI driver socket (defined above) inside the
  pod that the `node-driver-registrar` container will use to issue CSI
  operations (e.g. `/csi/csi.sock`).
* `--kubelet-registration-path`: This is the path to the CSI driver socket on
  the host node that kubelet will use to issue CSI operations (e.g.
  `/var/lib/kubelet/plugins/<drivername.example.com>/csi.sock`). Note this is NOT
  the path to the registration socket.

### Optional arguments

* `--http-endpoint`: The TCP network address where the HTTP server for diagnostics, including
  the health check indicating whether the registration socket exists, will listen (example:
  `:8080`). The default is empty string, which means the server is disabled.

* `--health-port`: (deprecated) This is the port of the health check server for the
  node-driver-registrar, which checks if the registration socket exists. A value &lt;= 0 disables
  the server. Server is disabled by default.

* `--timeout <duration>`: Timeout of all calls to CSI driver. It should be set to a value that accommodates the `GetDriverName` calls. 1 second is used by default.

* `--mode <mode>` (default: `--mode=registration`): The running mode of node-driver-registrar. `registration` runs node-driver-registrar as a long running process to register the driver with kubelet. `kubelet-registration-probe` runs as a health check and returns a status code of 0 if the driver was registered successfully. In the probe definition make sure that the value of `--kubelet-registration-path` is the same as in the container.

* `--enable-pprof`: Enable pprof profiling on the TCP network address specified by `--http-endpoint`.

### Required permissions

The node-driver-registrar does not interact with the Kubernetes API, so no RBAC
rules are needed.

It does, however, need to be able to mount hostPath volumes and have the file
permissions to:

* Access the CSI driver socket (typically in `/var/lib/kubelet/plugins/<drivername.example.com>/`).
  * Used by the `node-driver-registrar` to fetch the driver name from the driver
    contain (via the CSI `GetPluginInfo()` call).
* Access the registration socket (typically in `/var/lib/kubelet/plugins_registry/`).
  * Used by the `node-driver-registrar` to register the driver with kubelet.

### Health Check with the http server

If `--http-endpoint` is set, the node-driver-registrar exposes a health check endpoint at the
specified address and the path `/healthz`, indicating whether the registration socket exists.

### Health Check with an exec probe

If `--mode=kubelet-registration-probe` is set, node-driver-registrar can act as a probe checking if kubelet plugin registration succeeded. This is useful to detect if the registration got stuck as seen in issue [#143](https://github.com/kubernetes-csi/node-driver-registrar/issues/143)

The value of `--kubelet-registration-path` must be the same as the one set in the container args, `--csi-address` is not required in this mode, for example:

**Linux**

```yaml
  containers:
    - name: csi-driver-registrar
      image: k8s.gcr.io/sig-storage/csi-node-driver-registrar:v2.5.0
      args:
        - "--v=5"
        - "--csi-address=/csi/csi.sock"
        - "--kubelet-registration-path=/var/lib/kubelet/plugins/<drivername.example.com>/csi.sock"
      livenessProbe:
        exec:
          command:
          - /csi-node-driver-registrar
          - --kubelet-registration-path=/var/lib/kubelet/plugins/<drivername.example.com>/csi.sock
          - --mode=kubelet-registration-probe
        initialDelaySeconds: 30
        timeoutSeconds: 15
```

**Windows**
```yaml
  containers:
    - name: csi-driver-registrar
      image: k8s.gcr.io/sig-storage/csi-node-driver-registrar:v2.5.0
      args:
        - --v=5
        - --csi-address=unix://C:\\csi\\csi.sock
        - --kubelet-registration-path=C:\\var\\lib\\kubelet\\plugins\\<drivername.example.com>\\csi.sock
      livenessProbe:
        exec:
          command:
          - /csi-node-driver-registrar.exe
          - --kubelet-registration-path=C:\\var\\lib\\kubelet\\plugins\\<drivername.example.com>\\csi.sock
          - --mode=kubelet-registration-probe
        initialDelaySeconds: 30
        timeoutSeconds: 15
```

Related issue [#143](https://github.com/kubernetes-csi/node-driver-registrar/issues/143)

### Example

Here is an example sidecar spec in the driver DaemonSet. `<drivername.example.com>` should be replaced by
the actual driver's name.

```bash
      containers:
        - name: csi-driver-registrar
          image: k8s.gcr.io/sig-storage/csi-node-driver-registrar:v2.5.0
          args:
            - "--csi-address=/csi/csi.sock"
            - "--kubelet-registration-path=/var/lib/kubelet/plugins/<drivername.example.com>/csi.sock"
            - "--health-port=9809"
          volumeMounts:
            - name: plugin-dir
              mountPath: /csi
            - name: registration-dir
              mountPath: /registration
          ports:
            - containerPort: 9809
              name: healthz
          livenessProbe:
            httpGet:
              path: /healthz
              port: healthz
            initialDelaySeconds: 5
            timeoutSeconds: 5
      volumes:
        - name: registration-dir
          hostPath:
            path: /var/lib/kubelet/plugins_registry/
            type: Directory
        - name: plugin-dir
          hostPath:
            path: /var/lib/kubelet/plugins/<drivername.example.com>/
            type: DirectoryOrCreate
```

## Community, discussion, contribution, and support

Learn how to engage with the Kubernetes community on the [community page](http://kubernetes.io/community/).

You can reach the maintainers of this project at:

* Slack channels
  * [#wg-csi](https://kubernetes.slack.com/messages/wg-csi)
  * [#sig-storage](https://kubernetes.slack.com/messages/sig-storage)
* [Mailing list](https://groups.google.com/forum/#!forum/kubernetes-sig-storage)

### Code of conduct

Participation in the Kubernetes community is governed by the [Kubernetes Code of Conduct](code-of-conduct.md).
