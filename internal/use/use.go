package use

import (
	"errors"
	"fmt"
	"os"
	"strings"

	"github.com/navilg/k8senv/internal/config"
	"github.com/navilg/k8senv/internal/install"
)

func UseKubectl(version string) error {

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
		fmt.Println("kubectl version", version, "is not installed.")
		fmt.Println("Installing")
		install.InstallKubectl(version, false, 120)
	}

	if _, err := os.Lstat(kubectlBinaryPath); err == nil {
		os.Remove(kubectlBinaryPath)
	}

	err := os.Symlink(binaryFileName, kubectlBinaryPath)
	if err != nil {
		fmt.Println("Failed to setup kubectl", version, "as default.")
		fmt.Println(err)
		return (err)
	}

	fmt.Println("using kubectl", version, "as default.")

	return nil
}
