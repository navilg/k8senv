package config

import (
	"os"
	"runtime"
	"strings"
)

var version string = "v0.3.0"
var gitCommit string = "###GitCommitPlaceholder###"

type VersionInfo struct {
	GoVersion string `json:"GoVersion"`
	K8senv    string `json:"K8senv"`
	OS        string `json:"OS"`
	Arch      string `json:"ARCH"`
	GitCommit string `json:"GitCommit"`
}

var Version = VersionInfo{
	GoVersion: runtime.Version(),
	K8senv:    version,
	OS:        runtime.GOOS,
	Arch:      runtime.GOARCH,
	GitCommit: gitCommit,
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
