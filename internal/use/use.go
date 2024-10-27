package use

import (
	"errors"
	"fmt"
	"os"
	"runtime"
	"strings"

	"github.com/navilg/k8senv/internal/config"
	"github.com/navilg/k8senv/internal/ikubernetes"
	"github.com/navilg/k8senv/internal/install"
)

func UseVersion(toolname, version string) error {
	dotK8sEnvPath := config.GetDotK8senvPath()

	if dotK8sEnvPath == nil {
		fmt.Println(".k8senv/bin directory is not added in PATH environment variable")
		return errors.New(".k8senv/bin is not added in PATH environment variable")
	}

	if toolname == "kubectl" && version == "auto" {
		fmt.Println("Fetching Kubernetes server version")
		k8sVersion, err := ikubernetes.GetK8sVersion()
		if err != nil {
			fmt.Println("Error getting Kubernetes server version")
			return err
		}
		version = *k8sVersion
		fmt.Println("Kubernetes server version is", version)
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

	binaryFileName := *dotK8sEnvPath + "/" + toolname + "." + version
	binaryPath := *dotK8sEnvPath + "/" + toolname

	if _, err := os.Stat(binaryFileName); os.IsNotExist(err) {
		fmt.Println(toolname, "version", version, "is not installed.")
		fmt.Println("Installing")
		install.InstallVersion(toolname, version, false, 120, "")
	}

	if _, err := os.Lstat(binaryPath); err == nil {
		os.Remove(binaryPath)
	}

	err := os.Symlink(binaryFileName, binaryPath)
	if err != nil {
		fmt.Println("Failed to setup", toolname, version, "as default.")
		fmt.Println(err)
		return (err)
	}

	fmt.Println("Using", toolname, version, "as default.")

	if runtime.GOOS == "linux" {

		fmt.Println("Bash shell caches the program location. You might need to run 'hash -d " + toolname + "' to clear it for this to work.")
	}

	return nil
}
