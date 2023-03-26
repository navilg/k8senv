package remove

import (
	"errors"
	"fmt"
	"os"
	"strings"

	"github.com/navilg/k8senv/internal/config"
)

func RemoveKubectl(version string) error {
	dotK8sEnvPath := config.GetDotK8senvPath()
	if dotK8sEnvPath == nil {
		fmt.Println(".k8senv/bin directory is not added in PATH environment variable")
		return errors.New(".k8senv/bin is not added in PATH environment variable")
	}

	major_minor_patch_vers := strings.Split(version, ".")

	if !strings.HasPrefix(major_minor_patch_vers[0], "v") {
		version = "v" + version
	}

	if len(major_minor_patch_vers) == 2 {
		version = version + ".0"
	} else if len(major_minor_patch_vers) == 1 {
		version = version + ".0.0"
	}

	binaryFileName := *dotK8sEnvPath + "/kubectl." + version
	kubectlBinaryPath := *dotK8sEnvPath + "/kubectl"

	if _, err := os.Stat(binaryFileName); os.IsNotExist(err) {
		fmt.Println("Kubectl version", version, "is not installed.")
		return nil
	}

	if _, err := os.Lstat(kubectlBinaryPath); err == nil {
		currentKubectlPathInUse, err := os.Readlink(kubectlBinaryPath)
		if err != nil {
			fmt.Println("Failed unset kubectl version", version, "as default")
			fmt.Println(err)
		}

		if currentKubectlPathInUse == binaryFileName {
			err = os.Remove(kubectlBinaryPath)
			if err != nil {
				fmt.Println("Failed unset kubectl version", version, "as default")
				fmt.Println("Due to,", err)
			}
		}
	}

	err := os.Remove(binaryFileName)
	if err != nil {
		fmt.Println("Failed to kubectl version", version)
		fmt.Println("Due to,", err)
		return err
	}

	fmt.Println("Successfully removed kubectl version", version)

	return nil
}
