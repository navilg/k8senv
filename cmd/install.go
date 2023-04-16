/*
Copyright © 2023 NAME HERE <EMAIL ADDRESS>
*/
package cmd

import (
	"fmt"
	"os"

	"github.com/navilg/k8senv/internal/install"
	"github.com/spf13/cobra"
)

var overwriteInstall bool
var timeout int
var proxy string

// installCmd represents the install command
var installCmd = &cobra.Command{
	Use:   "install",
	Short: "Install a version of kubectl, helm or velero client",
	// Long:  ``,
	// Run: func(cmd *cobra.Command, args []string) {
	// 	fmt.Println("install called")
	// },
}

// installCmd represents the install command
var kubectlInstallCmd = &cobra.Command{
	Use:   "install VERSION",
	Short: "Install a version of kubectl",
	Long: `Install a version of kubectl client
	
Examples:
	# Install kubectl version 1.26.2
	k8senv kubectl install v1.26.2

	# Install latest available stable version of kubectl
	k8senv kubectl install latest

	# Install kubectl version 1.20.0 and overwrite it if it already exists
	k8senv kubectl install 1.20.0 --overwrite

	# Install kubectl version 1.20.0 aand set timeout to 300 seconds (If internet is slow), Default: 120 seconds
	k8senv kubectl install 1.20.0 --timeout=300

Supported version formats:
	v1.20.3
	1.20.3	# Defaults to v1.20.3
	1.20 	# Defaults to v1.20.0
	1 	# Defaults to v1.0.0`,
	Run: func(cmd *cobra.Command, args []string) {
		if len(args) != 1 {
			fmt.Println("Exactly one argumanet is required. Provide kubectl version to install e.g. v1.20.3")
			os.Exit(1)
		}
		err := install.InstallVersion("kubectl", args[0], overwriteInstall, timeout, proxy)
		if err != nil {
			os.Exit(1)
		}

	},
}

var veleroInstallCmd = &cobra.Command{
	Use:   "install VERSION",
	Short: "Install a version of velero",
	Long: `Install a version of velero client
	
Examples:
	# Install velero client version 1.10.2
	k8senv velero install v1.10.2

	# Install latest available stable version of velero client
	k8senv velero install latest

	# Install velero client version 1.8.0 and overwrite it if it already exists
	k8senv velero install 1.8.0 --overwrite

	# Install velero client version 1.8.0 aand set timeout to 300 seconds (If internet is slow), Default: 120 seconds
	k8senv velero install 1.8.0 --timeout=300

Supported version formats:
	v1.10.2
	1.10.2	# Defaults to v1.10.2
	1.10 	# Defaults to v1.10.0
	1 	# Defaults to v1.0.0`,
	Run: func(cmd *cobra.Command, args []string) {
		if len(args) != 1 {
			fmt.Println("Exactly one argumanet is required. Provide velero version to install e.g. v1.20.3")
			os.Exit(1)
		}
		err := install.InstallVersion("velero", args[0], overwriteInstall, timeout, proxy)
		if err != nil {
			os.Exit(1)
		}

	},
}

var helmInstallCmd = &cobra.Command{
	Use:   "install VERSION",
	Short: "Install a version of helm",
	Long: `Install a version of helm
	
Examples:
	# Install helm version 3.10.2
	k8senv helm install v3.10.2

	# Install latest available stable version of helm
	k8senv helm install latest

	# Install helm version 3.8.0 and overwrite it if it already exists
	k8senv helm install 3.8.0 --overwrite

	# Install helm version 3.8.0 aand set timeout to 300 seconds (If internet is slow), Default: 120 seconds
	k8senv helm install 3.8.0 --timeout=300

Supported version formats:
	v3.10.2
	3.10.2	# Defaults to v3.10.2
	1.10 	# Defaults to v1.10.0
	1 	# Defaults to v1.0.0`,
	Run: func(cmd *cobra.Command, args []string) {
		if len(args) != 1 {
			fmt.Println("Exactly one argumanet is required. Provide helm version to install e.g. v3.10.2")
			os.Exit(1)
		}
		err := install.InstallVersion("helm", args[0], overwriteInstall, timeout, proxy)
		if err != nil {
			os.Exit(1)
		}

	},
}

func init() {
	rootCmd.AddCommand(installCmd)
	kubectlCmd.AddCommand(kubectlInstallCmd)
	veleroCmd.AddCommand(veleroInstallCmd)
	helmCmd.AddCommand(helmInstallCmd)

	// Here you will define your flags and configuration settings.

	// Cobra supports Persistent Flags which will work for this command
	// and all subcommands, e.g.:
	// installCmd.PersistentFlags().String("foo", "", "A help for foo")

	// installKubectlCmd.PersistentFlags().BoolP("overwrite", "f", false, "Overwrite or re-install existing version")
	installKubectlCmd.PersistentFlags().BoolVarP(&overwriteInstall, "overwrite", "f", false, "Overwrite or re-install existing version")
	installKubectlCmd.PersistentFlags().IntVarP(&timeout, "timeout", "t", 120, "Timeout in seconds [DEFAULT: 120 seconds]")
	installKubectlCmd.PersistentFlags().StringVarP(&proxy, "proxy", "p", "", "HTTP/HTTPS proxy to use for downloading clients from its source")

	installVeleroCmd.PersistentFlags().BoolVarP(&overwriteInstall, "overwrite", "f", false, "Overwrite or re-install existing version")
	installVeleroCmd.PersistentFlags().IntVarP(&timeout, "timeout", "t", 120, "Timeout in seconds [DEFAULT: 120 seconds]")
	installVeleroCmd.PersistentFlags().StringVarP(&proxy, "proxy", "p", "", "HTTP/HTTPS proxy to use for downloading clients from its source")

	installHelmCmd.PersistentFlags().BoolVarP(&overwriteInstall, "overwrite", "f", false, "Overwrite or re-install existing version")
	installHelmCmd.PersistentFlags().IntVarP(&timeout, "timeout", "t", 120, "Timeout in seconds [DEFAULT: 120 seconds]")
	installHelmCmd.PersistentFlags().StringVarP(&proxy, "proxy", "p", "", "HTTP/HTTPS proxy to use for downloading clients from its source")

	// Cobra supports local flags which will only run when this command
	// is called directly, e.g.:
	// installCmd.Flags().BoolP("toggle", "t", false, "Help message for toggle")
}
