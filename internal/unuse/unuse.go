package unuse

import (
	"errors"
	"fmt"
	"os"

	"github.com/navilg/k8senv/internal/config"
)

func UnuseVersions(toolname string) error {
	dotK8sEnvPath := config.GetDotK8senvPath()
	if dotK8sEnvPath == nil {
		fmt.Println(".k8senv/bin directory is not added in PATH environment variable")
		return errors.New(".k8senv/bin is not added in PATH environment variable")
	}

	binaryPath := *dotK8sEnvPath + "/" + toolname

	if _, err := os.Lstat(binaryPath); err == nil {
		err = os.Remove(binaryPath)
		if err != nil {
			fmt.Println("Failed to unuse", toolname)
			fmt.Println("Due to,", err)
		}
	}

	return nil
}
