#!/usr/bin/env bash

set -e

# Build

bash build.sh
chmod +x ./k8senv-linux-x86_64
echo

mkdir -p $HOME/.k8senv/bin
export PATH=$HOME/.k8senv/bin:$PATH

mv ./k8senv-linux-x86_64 $HOME/.k8senv/bin/k8senv

echo "Testing K8senv version"
k8senv version
echo "✅"
echo

# Test kubectl

echo "Testing kubectl install"
k8senv install kubectl v1.26.2 --timeout 300
k8senv kubectl install 1.23.2 --timeout 300
echo "✅"
echo

echo "Testing kubectl use"
k8senv use kubectl v1.26.2
kubectl version --client
k8senv kubectl use v1.23.2
kubectl version --client
echo "✅"
echo

echo "Testing kubectl list"
k8senv kubectl list
k8senv list kubectl
echo "✅"
echo

echo "Testing kubectl remove"
k8senv kubectl remove v1.23.2
k8senv list kubectl
echo "✅"
echo

# Test velero

echo "Testing velero install"
k8senv install velero v1.10.2 --timeout 300
k8senv velero install 1.8.1 --timeout 300
echo "✅"
echo