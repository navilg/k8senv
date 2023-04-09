#!/usr/bin/env bash

go version || exit 1

# Linux
echo "Building for Linux OS with AMD64 Arch"
CGO_ENABLED=0  GOOS=linux GOARCH=amd64 go build -o k8senv-linux-x86_64 main.go && echo "✅ DONE" || echo "❌ FAILED"
echo "Building for Linux OS with ARM64 Arch"
CGO_ENABLED=0  GOOS=linux GOARCH=arm64 go build -o k8senv-linux-arm64  main.go && echo "✅ DONE" || echo "❌ FAILED"

# macOS
# echo "Building for MacOS with AMD64 Arch"
# CGO_ENABLED=0  GOOS=darwin GOARCH=amd64 go build -o k8senv-macos-x86_64 main.go && echo "✅ DONE" || echo "❌ FAILED"
# echo "Building for MacOS with ARM64 Arch"
# CGO_ENABLED=0  GOOS=darwin GOARCH=arm64 go build -o k8senv-macos-arm64  main.go && echo "✅ DONE" || echo "❌ FAILED"