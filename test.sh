#!/usr/bin/env bash

set -e

# Build

bash build.sh
chmod +x ./k8senv-linux-amd64
echo

mkdir -p $HOME/.k8senv/bin
export PATH=$HOME/.k8senv/bin:$PATH

cp -f ./k8senv-linux-amd64 $HOME/.k8senv/bin/k8senv

echo "Testing 'k8senv version'"
k8senv version
echo "✅"
echo

# Test kubectl

echo "Testing 'k8senv install kubectl'"
k8senv install kubectl v1.26.2 --timeout 300 --overwrite
echo "Testing 'k8senv kubectl install'"
k8senv kubectl install 1.23.2 --timeout 300 --overwrite
echo "✅"
echo

echo "Testing 'k8senv use kubectl'"
k8senv use kubectl v1.26.2
kubectl version --client
echo "Testing 'k8senv kubectl use'"
k8senv kubectl use 1.23.2
kubectl version --client
echo "✅"
echo

echo "Testing 'k8senv kubectl list'"
k8senv kubectl list
echo "Testing 'k8senv list kubectl'"
k8senv list kubectl
echo "✅"
echo

echo "Testing 'k8senv kubectl remove'"
k8senv kubectl remove v1.23.2
k8senv list kubectl
echo "Testing 'k8senv remove kubectl'"
k8senv remove kubectl 1.26.2
k8senv list kubectl
echo "✅"
echo

# Test velero

echo "Testing 'k8senv install velero'"
k8senv install velero v1.10.2 --timeout 300 --overwrite
echo "Testing 'k8senv velero install'"
k8senv velero install 1.8.1 --timeout 300 --overwrite
echo "✅"
echo

echo "Testing 'k8senv use velero'"
k8senv use velero v1.8.1
velero version --client-only
echo "Testing 'k8senv velero use'"
k8senv velero use 1.10.2
velero version --client-only
echo "✅"
echo

echo "Testing 'k8senv velero list'"
k8senv velero list
echo "Testing 'k8senv list velero'"
k8senv list velero
echo "✅"
echo

echo "Testing 'k8senv velero remove'"
k8senv velero remove v1.10.2
k8senv list velero
echo "Testing 'k8senv remove velero'"
k8senv remove velero 1.8.1
k8senv list velero
echo "✅"
echo

# Test helm

echo "Testing 'k8senv install helm'"
k8senv install helm v3.10.2 --timeout 300 --overwrite
echo "Testing 'k8senv helm install'"
k8senv helm install 3.8.1 --timeout 300 --overwrite
echo "✅"
echo

echo "Testing 'k8senv use helm'"
k8senv use helm v3.8.1
helm version
echo "Testing 'k8senv helm use'"
k8senv helm use 3.10.2
helm version
echo "✅"
echo

echo "Testing 'k8senv helm list'"
k8senv helm list
echo "Testing 'k8senv list helm'"
k8senv list helm
echo "✅"
echo

echo "Testing 'k8senv helm remove'"
k8senv helm remove v3.10.2
k8senv list helm
echo "Testing 'k8senv remove helm'"
k8senv remove helm 3.8.1
k8senv list helm
echo "✅"
echo