package list

import (
	"errors"
	"fmt"
	"io/ioutil"
	"os"
	"path/filepath"
	"strings"

	"github.com/navilg/k8senv/internal/config"
)

func ListVersions(toolname string) error {
	dotK8sEnvPath := config.GetDotK8senvPath()
	if dotK8sEnvPath == nil {
		fmt.Println(".k8senv/bin directory is not added in PATH environment variable")
		return errors.New(".k8senv/bin is not added in PATH environment variable")
	}

	fileinfo, err := ioutil.ReadDir(*dotK8sEnvPath)
	if err != nil {
		fmt.Println("Failed to list installed versions")
		fmt.Println(err)
		return (err)
	}

	binaryPath := *dotK8sEnvPath + "/" + toolname
	var versionInUse string

	if _, err := os.Lstat(binaryPath); err == nil {
		currentToolPathInUse, err := os.Readlink(binaryPath)
		if err != nil {
			fmt.Println("Failed to list installed versions")
			fmt.Println(err)
			return (err)
		}

		versionInUse = strings.TrimPrefix(filepath.Base(currentToolPathInUse), toolname+".")
	}

	count := 0

	for _, file := range fileinfo {
		if !file.IsDir() {
			if strings.HasPrefix(file.Name(), toolname+".") {
				version := strings.TrimPrefix(file.Name(), toolname+".")
				if versionInUse == version {
					fmt.Println("*", version)
				} else {
					fmt.Println(" ", version)
				}
				count = count + 1
			}
		}
	}

	if count == 0 {
		fmt.Println("No version of", toolname, "is installed by k8senv.")
	}

	return nil
}
