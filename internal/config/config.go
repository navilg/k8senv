package config

import (
	"os"
	"runtime"
	"strings"
)

type VersionInfo struct {
	GoVersion string `json:"GoVersion"`
	K8senv    string `json:"K8senv"`
}

var Version = VersionInfo{
	GoVersion: runtime.Version(),
	K8senv:    "v0.1.2",
}

func GetDotK8senvPath() *string {
	pathenv := os.Getenv("PATH")
	paths := strings.Split(pathenv, ":")

	for _, path := range paths {
		if strings.Contains(path, ".k8senv/bin") {
			return &path
		}
	}

	return nil
}
