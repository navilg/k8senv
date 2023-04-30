# k8senv

[Kubernetes](https://kubernetes.io) client's version manager inspired by [tfenv](https://github.com/tfutils/tfenv)

Kubernetes clients supported at this moments are `kubectl`, or, `helm`, or, `velero` on Linux servers/machines.

Manytimes we need to manage many Kubernetes clusters from same jumpbox or machine or bastion server. Those many cluster may have different versions of Kubernetes. In that case, It becomes difficult to use multiple versions of Kubernetes clients to communicate with those different versions of clusters.

This small tool is to help manage different versions of Kubernetes clients. Kubernetes clients `kubectl`, `helm` and `velero` are supported at this stage.

**NOTE: Currently only `kubectl` is supported. Support for `helm` and `velero` is under development.**

## How to install/setup

1. Create a directory .k8senv/bin in home directory

```
mkdir -p ~/.k8senv/bin
```

2. Download `k8senv` tool into `.k8senv/bin` directory

[amd64](https://github.com/navilg/k8senv/releases/latest/download/k8senv-linux-amd64) Intel or AMD 64-Bit CPU

Download latest `k8senv` tool:
```
cd ~/.k8senv/bin
curl -L https://github.com/navilg/k8senv/releases/latest/download/k8senv-linux-amd64 -o k8senv
chmod +x k8senv
```

Download specific version of `k8senv`. For e.g. To download version `v1.0.0`:

```
cd ~/.k8senv/bin
curl -L https://github.com/navilg/k8senv/releases/download/v1.0.0/k8senv-linux-amd64 -o k8senv
chmod +x k8senv
```

[arm64](https://github.com/navilg/k8senv/releases/latest/download/k8senv-linux-arm64) Arm-based 64-Bit CPU (i.e. in Raspberry Pi)

Download latest `k8senv` tool:
```
cd ~/.k8senv/bin
curl -L https://github.com/navilg/k8senv/releases/latest/download/k8senv-linux-arm64 -o k8senv
chmod +x k8senv
```

Download specific version of `k8senv`. For e.g. To download version `v1.0.0`:

```
cd ~/.k8senv/bin
curl -L https://github.com/navilg/k8senv/releases/download/v1.0.0/k8senv-linux-arm64 -o k8senv
chmod +x k8senv
```

3. Add `~/.k8senv/bin` directory to your `PATH` environment variable

```
export PATH="$HOME/.k8senv/bin:$PATH"
echo 'export PATH="$HOME/.k8senv/bin:$PATH"' >> ~/.bashrc
hash -r
```

```
which k8senv
# OR
type k8senv
```

## Usage

```
k8senv [VERB] [CLIENT] [VERSION]

# [VERB] can be install, use, list or remove
# [CLIENT] can be kubectl, velero or helm
# [VERSION] can be any client's version in format v*.*.*

examples:
k8senv use kubectl v1.23.2
k8senv install velero v1.8.1
```

Examples:

**Install a kubectl version**

Any of below commands can be used to download 1.26.2 version of kubectl

```
k8senv install kubectl v1.26.2
k8senv install kubectl 1.26.2
k8senv kubectl install v1.26.2
k8senv install kubectl 1.26.2 --overwrite   # Installs even if it already exists
k8senv install kubectl latest               # Installs latest stable version of kubectl
k8s install kubectl v1.19.2 --timeout=300   # Install 1.19.2 with timeout of 300 seconds. Default timeout is 120 seconds.
```

**List all installed version of kubectl**

Any of below commandds can be used to list all kubectl client's version installed by k8senv.

```
k8senv list kubectl
k8senv kubectl list
```

**Switch to a version of kubectl**

Any of below commands can be used to switch kubectl version to 1.26.2. If version is not available, It will install it automatically.

```
k8senv use kubectl v1.26.2
k8senv use kubectl 1.26.2
k8senv kubectl use v1.26.2
```

**Unuse/Disable k8senv managed kubectl**

Any of the below command can be used to unuse or disable k8senv managed kubectl without uninstalling it. This will reset the system to use system installed kubectl, if present.

```
k8senv unuse kubectl
k8senv kubectl unuse
```

**Remove an existing version of kubectl**

Any of below commands can be used to remove kubectl version 1.26.2.

```
k8senv remove kubectl v1.26.2
k8senv remove kubectl 1.26.2
k8senv kubectl remove v1.26.2
```

**Similar command can be used to install, use, list and remove velero and helm clients.**

For e.g.

```
k8senv install helm v3.10.2
k8senv remove velero 1.10.1
```