package config

import (
	"os"
	"strings"
)

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
