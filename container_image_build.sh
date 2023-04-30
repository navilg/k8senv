#!/usr/bin/env bash

# Docker image

version=$(go run main.go version | cut -d " " -f 2 | jq .K8senv | tr -d "\"")
docker buildx create --name dockerxbuilder --use --bootstrap

if [ "$1" == "push" ]; then
    docker buildx build --push --platform linux/amd64,linux/arm64,linux/arm/v7 --tag linuxshots/k8senv:$version .
    docker buildx build --push --platform linux/amd64,linux/arm64,linux/arm/v7 --tag linuxshots/k8senv .
else
    docker buildx build --platform linux/amd64,linux/arm64,linux/arm/v7 --tag linuxshots/k8senv:$version .
    docker buildx build --platform linux/amd64,linux/arm64,linux/arm/v7 --tag linuxshots/k8senv .
fi
