package install

import (
	"archive/tar"
	"compress/gzip"
	"encoding/json"
	"errors"
	"fmt"
	"io"
	"io/ioutil"
	"os"
	"path/filepath"
	"strings"
	"time"

	"github.com/navilg/k8senv/internal/checksum"
	"github.com/navilg/k8senv/internal/config"
	"github.com/navilg/k8senv/internal/download"
)

func InstallVersion(toolname, version string, overwrite bool, timeout int, proxy string) error {
	if toolname == "kubectl" {
		err := InstallKubectl(version, overwrite, timeout, proxy)
		if err != nil {
			return err
		}
	} else if toolname == "velero" {
		err := InstallVelero(version, overwrite, timeout, proxy)
		if err != nil {
			return err
		}
	} else if toolname == "helm" {
		err := InstallHelm(version, overwrite, timeout, proxy)
		if err != nil {
			return err
		}
	} else {
		fmt.Println(toolname, "is not a valid tool supported by k8senv.")
		return errors.New(toolname + " is not a valid tool supported by k8senv.")
	}

	return nil
}

func InstallKubectl(version string, overwrite bool, timeout int, proxy string) error {
	latestVersionUrl := "https://storage.googleapis.com/kubernetes-release/release/stable.txt"
	dotK8sEnvPath := config.GetDotK8senvPath()

	if dotK8sEnvPath == nil {
		fmt.Println(".k8senv/bin directory is not added in PATH environment variable")
		return errors.New(".k8senv/bin is not added in PATH environment variable")
	}

	if version == "latest" {
		// version value is latest

		fmt.Println("Fetching latest stable version")
		data, err := download.Download(latestVersionUrl, 30, proxy)
		if err != nil {
			fmt.Println("Failed to fetch latest kubectl version")
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

	goos := config.Version.OS
	goarch := config.Version.Arch

	downloadUrl := "https://dl.k8s.io/release/" + version + "/bin/" + goos + "/" + goarch + "/kubectl"
	checksumUrl := "https://dl.k8s.io/" + version + "/bin/" + goos + "/" + goarch + "/kubectl.sha256"
	binaryFileName := *dotK8sEnvPath + "/kubectl." + version

	if _, err := os.Stat(binaryFileName); err == nil && !overwrite {
		fmt.Println("kubectl version", version, "is already installed. Use command `k8senv use kubectl", version+"` to use it.")
		fmt.Println("If existing client doesnot work properly or is corrupted, Use --overwrite flag to overwrite/re-install the existing one.")
		return nil
	}

	fmt.Println("Downloading kubectl version", version)
	fmt.Println("Download in progress... It may take time depending on internet speed. Default timeout:", timeout, "seconds.")

	data, err := download.Download(downloadUrl, time.Duration(timeout), proxy)
	if err != nil {
		fmt.Println("Failed to install kubectl version", version)
		return err
	}

	err = ioutil.WriteFile(binaryFileName, data, 0750)
	if err != nil {
		fmt.Println("Failed to install kubectl version", version)
		fmt.Println(err)
		return err
	}

	fmt.Println("Downloaded kubectl version", version)

	fmt.Println("Validating checksum")

	checksumdata, err := download.Download(checksumUrl, 30, proxy)
	if err != nil {
		fmt.Println("Failed to validate checksum")
		return err
	}

	if isValid := checksum.ValidateSHA256Sum(strings.TrimSuffix(string(checksumdata), "\n"), binaryFileName); isValid {
		fmt.Println("Checksum validated.")
	} else {
		fmt.Println("Failed to validate checksum. Deleting the installed client.")
		_ = os.Remove(binaryFileName)
		return errors.New("Failed to validate checksum of installed file")
	}

	return nil
}

func InstallVelero(version string, overwrite bool, timeout int, proxy string) error {
	latestVersionUrl := "https://api.github.com/repos/vmware-tanzu/velero/releases/latest"
	dotK8sEnvPath := config.GetDotK8senvPath()

	if dotK8sEnvPath == nil {
		fmt.Println(".k8senv/bin directory is not added in PATH environment variable")
		return errors.New(".k8senv/bin is not added in PATH environment variable")
	}

	if version == "latest" {
		// version value is latest

		fmt.Println("Fetching latest stable version")
		data, err := download.Download(latestVersionUrl, 30, proxy)
		if err != nil {
			fmt.Println("Failed to fetch latest velero client version")
			return err
		}

		type latestVeleroVersionInfo struct {
			TagName string `json:"tag_name"`
		}

		var latestVeleroVersion latestVeleroVersionInfo

		err = json.Unmarshal(data, &latestVeleroVersion)
		if err != nil {
			fmt.Println("Failed to fetch latest velero client version")
			return err
		}

		version = string(latestVeleroVersion.TagName)
		fmt.Println("Latest available stable version of velero client is", version)

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

	fullVersion := version + "-" + config.Version.OS + "-" + config.Version.Arch

	downloadUrl := "https://github.com/vmware-tanzu/velero/releases/download/" + version + "/velero-" + fullVersion + ".tar.gz"
	checksumUrl := "https://github.com/vmware-tanzu/velero/releases/download/" + version + "/CHECKSUM"
	binaryFileName := *dotK8sEnvPath + "/velero." + version

	if _, err := os.Stat(binaryFileName); err == nil && !overwrite {
		fmt.Println("velero client version", version, "is already installed. Use command `k8senv use velero", version+"` to use it.")
		fmt.Println("If existing client doesnot work properly or is corrupted, Use --overwrite flag to overwrite/re-install the existing one.")
		return nil
	}

	fmt.Println("Downloading velero package version", version)
	tempDir, err := ioutil.TempDir("/tmp", "velero"+version+"*")
	if err != nil {
		fmt.Println("Failed to create temporary directory")
		fmt.Println(err)
	}

	fmt.Println("Download in progress... It may take time depending on internet speed. Default timeout:", timeout, "seconds.")

	data, err := download.Download(downloadUrl, time.Duration(timeout), proxy)
	if err != nil {
		fmt.Println("Failed to install velero client version", version)
		return err
	}

	err = ioutil.WriteFile(tempDir+"/velero-"+fullVersion+".tar.gz", data, 0750)
	if err != nil {
		fmt.Println("Failed to install velero client version", version)
		fmt.Println(err)
		return err
	}

	fmt.Println("Installation package downloaded for velero client version", version)

	fmt.Println("Validating checksum")

	checksumdata, err := download.Download(checksumUrl, 30, proxy)
	if err != nil {
		fmt.Println("Failed to validate checksum")
		return err
	}

	var isChecksumValidated bool = false

	for _, line := range strings.Split(string(checksumdata), "\n") {
		words := strings.Fields(line)
		if len(words) < 2 {
			continue
		}
		if words[1] == "velero-"+fullVersion+".tar.gz" {
			if checksum.ValidateSHA256Sum(strings.TrimSuffix(string(words[0]), "\n"), tempDir+"/velero-"+fullVersion+".tar.gz") {
				isChecksumValidated = true
			}
			break
		}
	}

	if isChecksumValidated {
		fmt.Println("Checksum validated.")
	} else {
		fmt.Println("Failed to validate checksum. Deleting the downloaded package.")
		_ = os.Remove(tempDir + "/velero-" + fullVersion + ".tar.gz")
		return errors.New("Failed to validate checksum of downloaded file")
	}

	// Gun-Unzipping

	fmt.Println("Unzipping the package")
	reader, err := os.Open(tempDir + "/velero-" + fullVersion + ".tar.gz")
	if err != nil {
		fmt.Println("Failed to unzip the package")
		fmt.Println(err)
		return err
	}
	defer reader.Close()

	archive, err := gzip.NewReader(reader)
	if err != nil {
		fmt.Println("Failed to unzip the package")
		fmt.Println(err)
		return err
	}
	defer archive.Close()

	target := filepath.Join(tempDir+"/velero-"+fullVersion+".tar", archive.Name)
	writer, err := os.Create(target)
	if err != nil {
		fmt.Println("Failed to unzip the package")
		fmt.Println(err)
		return err
	}
	defer writer.Close()

	_, err = io.Copy(writer, archive)
	if err != nil {
		fmt.Println("Failed to unzip the package")
		fmt.Println(err)
		return err
	}

	// Untaring file
	fmt.Println("Getting the velero client")
	reader, err = os.Open(tempDir + "/velero-" + fullVersion + ".tar")
	if err != nil {
		fmt.Println(err)
		return err
	}
	defer reader.Close()
	tarReader := tar.NewReader(reader)

	for {
		header, err := tarReader.Next()
		if err == io.EOF {
			fmt.Println("Failed to get velero client")
			fmt.Println("Binary file not present in package")
			break
		} else if err != nil {
			fmt.Println("Failed to get velero client")
			fmt.Println(err)
			return err
		}
		if header.FileInfo().IsDir() {
			continue
		}

		if filepath.Base(header.Name) != "velero" {
			continue
		}

		file, err := os.OpenFile(binaryFileName, os.O_CREATE|os.O_TRUNC|os.O_WRONLY, 0750)
		if err != nil {
			fmt.Println("Failed to install velero client")
			fmt.Println(err)
			return err
		}
		defer file.Close()

		_, err = io.Copy(file, tarReader)
		if err != nil {
			fmt.Println("Failed to install velero client")
			fmt.Println(err)
			return err
		}

		break
	}

	fmt.Println("Installed velero client version", version)

	return nil
}

func InstallHelm(version string, overwrite bool, timeout int, proxy string) error {

	latestVersionUrl := "https://api.github.com/repos/helm/helm/releases/latest"
	dotK8sEnvPath := config.GetDotK8senvPath()

	if dotK8sEnvPath == nil {
		fmt.Println(".k8senv/bin directory is not added in PATH environment variable")
		return errors.New(".k8senv/bin is not added in PATH environment variable")
	}

	if version == "latest" {
		// version value is latest

		fmt.Println("Fetching latest stable version")
		data, err := download.Download(latestVersionUrl, 30, proxy)
		if err != nil {
			fmt.Println("Failed to fetch latest helm version")
			return err
		}

		type latestHelmVersionInfo struct {
			TagName string `json:"tag_name"`
		}

		var latestHelmVersion latestHelmVersionInfo

		err = json.Unmarshal(data, &latestHelmVersion)
		if err != nil {
			fmt.Println("Failed to fetch latest helm version")
			return err
		}

		version = string(latestHelmVersion.TagName)
		fmt.Println("Latest available stable version of helm is", version)

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

	fullVersion := version + "-" + config.Version.OS + "-" + config.Version.Arch // Full version vM.m.p-os-arch e.g. v1.10.2-linux-amd64

	downloadUrl := "https://get.helm.sh/helm-" + fullVersion + ".tar.gz"
	checksumUrl := "https://get.helm.sh/helm-" + fullVersion + ".tar.gz.sha256sum"
	binaryFileName := *dotK8sEnvPath + "/helm." + version

	if _, err := os.Stat(binaryFileName); err == nil && !overwrite {
		fmt.Println("helm version", version, "is already installed. Use command `k8senv use helm", version+"` to use it.")
		fmt.Println("If existing client doesnot work properly or is corrupted, Use --overwrite flag to overwrite/re-install the existing one.")
		return nil
	}

	fmt.Println("Downloading helm package version", version)
	tempDir, err := ioutil.TempDir("/tmp", "helm"+version+"*")
	if err != nil {
		fmt.Println("Failed to create temporary directory")
		fmt.Println(err)
	}

	fmt.Println("Download in progress... It may take time depending on internet speed. Default timeout:", timeout, "seconds.")

	data, err := download.Download(downloadUrl, time.Duration(timeout), proxy)
	if err != nil {
		fmt.Println("Failed to install helm version", version)
		return err
	}

	err = ioutil.WriteFile(tempDir+"/helm-"+fullVersion+".tar.gz", data, 0750)
	if err != nil {
		fmt.Println("Failed to install helm version", version)
		fmt.Println(err)
		return err
	}

	fmt.Println("Installation package downloaded for helm version", version)

	fmt.Println("Validating checksum")

	checksumdata, err := download.Download(checksumUrl, 30, proxy)
	if err != nil {
		fmt.Println("Failed to validate checksum")
		return err
	}

	var isChecksumValidated bool = false

	for _, line := range strings.Split(string(checksumdata), "\n") {
		words := strings.Fields(line)
		if len(words) < 2 {
			continue
		}
		if words[1] == "helm-"+fullVersion+".tar.gz" {
			if checksum.ValidateSHA256Sum(strings.TrimSuffix(string(words[0]), "\n"), tempDir+"/helm-"+fullVersion+".tar.gz") {
				isChecksumValidated = true
			}
			break
		}
	}

	if isChecksumValidated {
		fmt.Println("Checksum validated.")
	} else {
		fmt.Println("Failed to validate checksum. Deleting the downloaded package.")
		_ = os.Remove(tempDir + "/helm-" + fullVersion + ".tar.gz")
		return errors.New("Failed to validate checksum of downloaded file")
	}

	// Gun-Unzipping

	fmt.Println("Unzipping the package")
	reader, err := os.Open(tempDir + "/helm-" + fullVersion + ".tar.gz")
	if err != nil {
		fmt.Println("Failed to unzip the package")
		fmt.Println(err)
		return err
	}
	defer reader.Close()

	archive, err := gzip.NewReader(reader)
	if err != nil {
		fmt.Println("Failed to unzip the package")
		fmt.Println(err)
		return err
	}
	defer archive.Close()

	target := filepath.Join(tempDir+"/helm-"+fullVersion+".tar", archive.Name)
	writer, err := os.Create(target)
	if err != nil {
		fmt.Println("Failed to unzip the package")
		fmt.Println(err)
		return err
	}
	defer writer.Close()

	_, err = io.Copy(writer, archive)
	if err != nil {
		fmt.Println("Failed to unzip the package")
		fmt.Println(err)
		return err
	}

	// Untaring file
	fmt.Println("Getting the helm")
	reader, err = os.Open(tempDir + "/helm-" + fullVersion + ".tar")
	if err != nil {
		fmt.Println(err)
		return err
	}
	defer reader.Close()
	tarReader := tar.NewReader(reader)

	for {
		header, err := tarReader.Next()
		if err == io.EOF {
			fmt.Println("Failed to get helm")
			fmt.Println("Binary file not present in package")
			break
		} else if err != nil {
			fmt.Println("Failed to get helm")
			fmt.Println(err)
			return err
		}
		if header.FileInfo().IsDir() {
			continue
		}

		if filepath.Base(header.Name) != "helm" {
			continue
		}

		file, err := os.OpenFile(binaryFileName, os.O_CREATE|os.O_TRUNC|os.O_WRONLY, 0750)
		if err != nil {
			fmt.Println("Failed to install helm")
			fmt.Println(err)
			return err
		}
		defer file.Close()

		_, err = io.Copy(file, tarReader)
		if err != nil {
			fmt.Println("Failed to install helm")
			fmt.Println(err)
			return err
		}

		break
	}

	fmt.Println("Installed helm version", version)

	return nil
}
