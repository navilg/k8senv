# k8senv

[Kubernetes](https://kubernetes.io) client's version manager inspired by [tfenv](https://github.com/tfutils/tfenv)

Kubernetes clients supported at this moments are `kubectl`, or, `helm`, or, `velero` on Linux servers/machines.

Manytimes we need to manage many Kubernetes clusters from same jumpbox or machine or bastion server. Those many cluster may have different versions of Kubernetes. In that case, It becomes difficult to use multiple versions of Kubernetes clients to communicate with those different versions of clusters.

This small tool is to help manage different versions of Kubernetes clients. Kubernetes clients `kubectl`, `helm` and `velero` are supported at this stage.

## How to install/setup

1. Create a directory .k8senv/bin in home directory `mkdir -p ~/.k8senv/bin`
2. Download `k8senv` tool into `.k8senv/bin` directory

```
cd ~/.k8senv/bin
curl -LO https://raw.githubusercontent.com/navilg/k8senv/main/bin/k8senv
chmod +x k8senv
```

3. Add `~/.k8senv/bin` directory to your `PATH` environment variable

```
export PATH="$HOME/.k8senv/bin:$PATH"
echo 'export PATH="$HOME/.k8senv/bin:$PATH"' >> ~/.bashrc
```

```
which k8senv
```

**Under development**