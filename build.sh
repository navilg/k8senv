#!/usr/bin/env bash

go version || exit 1

commitId=$(git log --format="%H" -n 1)
echo $commitId

sed -i "s|###GitCommitPlaceholder###|${commitId}|g" internal/config/config.go

# Linux
echo "Building for Linux OS with AMD64 Arch"
CGO_ENABLED=0  GOOS=linux GOARCH=amd64 go build -o k8senv-linux-amd64 main.go && echo "✅ DONE" || echo "❌ FAILED"
echo "Building for Linux OS with ARM64 Arch"
CGO_ENABLED=0  GOOS=linux GOARCH=arm64 go build -o k8senv-linux-arm64  main.go && echo "✅ DONE" || echo "❌ FAILED"

sed -i "s|${commitId}|###GitCommitPlaceholder###|g" internal/config/config.go

# macOS
# echo "Building for MacOS with AMD64 Arch"
# CGO_ENABLED=0  GOOS=darwin GOARCH=amd64 go build -o k8senv-macos-x86_64 main.go && echo "✅ DONE" || echo "❌ FAILED"
# echo "Building for MacOS with ARM64 Arch"
# CGO_ENABLED=0  GOOS=darwin GOARCH=arm64 go build -o k8senv-macos-arm64  main.go && echo "✅ DONE" || echo "❌ FAILED"