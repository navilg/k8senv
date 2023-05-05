package unuse

import (
	"errors"
	"fmt"
	"os"
	"runtime"

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
			return err
		}
		if runtime.GOOS == "linux" {
			fmt.Println(toolname, "successfully set to unuse.")
			fmt.Println("Bash shell caches the program location. You might need to run 'hash -d " + toolname + "' to clear it for this to work.")
		}
	}

	return nil
}
