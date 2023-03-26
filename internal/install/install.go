package install

import (
	"errors"
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
	"os"
	"strings"

	"github.com/navilg/k8senv/internal/checksum"
	"github.com/navilg/k8senv/internal/config"
)

func InstallKubectl(version string, overwrite bool) error {
	latestVersionUrl := "https://storage.googleapis.com/kubernetes-release/release/stable.txt"
	dotK8sEnvPath := config.GetDotK8senvPath()

	if dotK8sEnvPath == nil {
		fmt.Println(".k8senv/bin directory is not added in PATH environment variable")
		return errors.New(".k8senv/bin is not added in PATH environment variable")
	}

	if version == "latest" {
		// version value is latest

		fmt.Println("Fetching latest stable version")
		client := http.Client{
			CheckRedirect: func(r *http.Request, via []*http.Request) error {
				r.URL.Opaque = r.URL.Path
				return nil
			},
		}

		resp, err := client.Get(latestVersionUrl)
		if err != nil {
			log.Fatal(err)
		}
		defer resp.Body.Close()

		if resp.StatusCode/100 != 2 {
			fmt.Println("Failed to fetch latest kubectl version")
			fmt.Println(err)
			return err
		}

		data, err := ioutil.ReadAll(resp.Body)
		if err != nil {
			fmt.Println("Failed to fetch latest kubectl version")
			fmt.Println(err)
			return err
		}

		version = string(data)
		fmt.Println("Latest available stable version of kubectl is", version)

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

	downloadUrl := "https://dl.k8s.io/release/" + version + "/bin/linux/amd64/kubectl"
	checksumUrl := "https://dl.k8s.io/" + version + "/bin/linux/amd64/kubectl.sha256"
	binaryFileName := *dotK8sEnvPath + "/kubectl." + version

	if _, err := os.Stat(binaryFileName); err == nil && !overwrite {
		fmt.Println("kubectl version", version, "is already installed.")
		fmt.Println("If existing client doesnot work properly or is corrupted, Use --overwrite flag to overwrite/re-install the existing one.")
		return nil
	}

	fmt.Println("Installing kubectl version", version)

	// Create http client
	client := http.Client{
		CheckRedirect: func(r *http.Request, via []*http.Request) error {
			r.URL.Opaque = r.URL.Path
			return nil
		},
	}

	// Perform HTTP GET request
	resp, err := client.Get(downloadUrl)
	if err != nil {
		fmt.Println("Failed to install kubectl version", version)
		fmt.Println("Due to: Failed to make HTTP GET request")
		fmt.Println(err)
		return err
	}

	defer resp.Body.Close()

	if resp.StatusCode/100 != 2 {
		fmt.Println("Failed to install kubectl version", version)
		fmt.Println(resp.Status)
		return err
	}

	data, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		fmt.Println("Failed to install kubectl version", version)
		fmt.Println("Due to: Failed to read received response")
		fmt.Println(err)
		return err
	}

	err = ioutil.WriteFile(binaryFileName, data, 0750)
	if err != nil {
		fmt.Println("Failed to install kubectl version", version)
		fmt.Println(err)
		return err
	}

	fmt.Println("Installed kubectl version", version)
	fmt.Println("Validating checksum")

	// Perform HTTP GET request
	resp, err = client.Get(checksumUrl)
	if err != nil {
		fmt.Println("Failed to validate checksum")
		fmt.Println("Due to: Failed to make HTTP GET request to checksum")
		fmt.Println(err)
		return err
	}

	defer resp.Body.Close()

	if resp.StatusCode/100 != 2 {
		fmt.Println("Failed to validate checksum")
		fmt.Println(resp.Status)
		return err
	}

	checksumdata, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		fmt.Println("Failed to validate checksum")
		fmt.Println(err)
		return err
	}

	if isValid := checksum.ValidateSHA256Sum(strings.TrimRight(string(checksumdata), "\n"), binaryFileName); isValid {
		fmt.Println("Checksum validated.")
	} else {
		fmt.Println("Failed to validate checksum. Deleting the installed client.")
		_ = os.Remove(binaryFileName)
		return errors.New("Failed to validate checksum of installed file")
	}

	return nil

}
