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

func ListKubectl() error {
	dotK8sEnvPath := config.GetDotK8senvPath()
	if dotK8sEnvPath == nil {
		fmt.Println(".k8senv/bin directory is not added in PATH environment variable")
		return errors.New(".k8senv/bin is not added in PATH environment variable")
	}

	fileinfo, err := ioutil.ReadDir(*dotK8sEnvPath)
	if err != nil {
		fmt.Println("Failed to list installed kubectl versions")
		fmt.Println(err)
		return (err)
	}

	kubectlBinaryPath := *dotK8sEnvPath + "/kubectl"
	var kubectlVersionInUse string

	if _, err := os.Lstat(kubectlBinaryPath); err == nil {
		currentKubectlPathInUse, err := os.Readlink(kubectlBinaryPath)
		if err != nil {
			fmt.Println("Failed to list installed kubectl versions")
			fmt.Println(err)
			return (err)
		}

		kubectlVersionInUse = strings.TrimPrefix(filepath.Base(currentKubectlPathInUse), "kubectl.")
	}

	count := 0

	for _, file := range fileinfo {
		if !file.IsDir() {
			if strings.HasPrefix(file.Name(), "kubectl.") {
				version := strings.TrimPrefix(file.Name(), "kubectl.")
				if kubectlVersionInUse == version {
					fmt.Println("*", version)
				} else {
					fmt.Println(" ", version)
				}
				count = count + 1
			}
		}
	}

	if count == 0 {
		fmt.Println("No version of kubectl client is installed by k8senv.")
	}

	return nil
}

func ListVelero() error {
	dotK8sEnvPath := config.GetDotK8senvPath()
	if dotK8sEnvPath == nil {
		fmt.Println(".k8senv/bin directory is not added in PATH environment variable")
		return errors.New(".k8senv/bin is not added in PATH environment variable")
	}

	fileinfo, err := ioutil.ReadDir(*dotK8sEnvPath)
	if err != nil {
		fmt.Println("Failed to list installed velero versions")
		fmt.Println(err)
		return (err)
	}

	veleroBinaryPath := *dotK8sEnvPath + "/velero"
	var veleroVersionInUse string

	if _, err := os.Lstat(veleroBinaryPath); err == nil {
		currentVeleroPathInUse, err := os.Readlink(veleroBinaryPath)
		if err != nil {
			fmt.Println("Failed to list installed velero versions")
			fmt.Println(err)
			return (err)
		}

		veleroVersionInUse = strings.TrimPrefix(filepath.Base(currentVeleroPathInUse), "velero.")
	}

	count := 0

	for _, file := range fileinfo {
		if !file.IsDir() {
			if strings.HasPrefix(file.Name(), "velero.") {
				version := strings.TrimPrefix(file.Name(), "velero.")
				if veleroVersionInUse == version {
					fmt.Println("*", version)
				} else {
					fmt.Println(" ", version)
				}
				count = count + 1
			}
		}
	}

	if count == 0 {
		fmt.Println("No version of velero client is installed by k8senv.")
	}

	return nil
}
