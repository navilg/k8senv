# k8senv

[Kubernetes](https://kubernetes.io) client's version manager inspired by [tfenv](https://github.com/tfutils/tfenv)

Kubernetes clients supported at this moments are `kubectl`, or, `helm`, or, `velero` on Linux servers/machines.

Manytimes we need to manage many Kubernetes clusters from same jumpbox or machine or bastion server. Those many cluster may have different versions of Kubernetes. In that case, It becomes difficult to use multiple versions of Kubernetes clients to communicate with those different versions of clusters.

This small tool is to help manage different versions of Kubernetes clients. Kubernetes clients `kubectl`, `helm` and `velero` are supported at this stage.

## How to install/setup

1. Create a directory .k8senv in home directory `mkdir $HOME/.k8senv`
2. Clone this repository into `.k8senv` directory

```
git clone --depth=1 https://github.com/navilg/k8senv.git ~/.k8senv
```

3. Add ~/.k8senv/bin to your $PATH

```
export PATH="$HOME/.k8senv/bin:$PATH"
echo 'export PATH="$HOME/.k8senv/bin:$PATH"' >> ~/.bashrc
```

Or, You can create a symlink to a directory already in PATH

```
sudo ln -s ~/.k8senv/bin/* /usr/local/bin
which k8senv
```