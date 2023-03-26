/*
Copyright © 2023 NAME HERE <EMAIL ADDRESS>
*/
package cmd

import (
	"fmt"
	"os"

	"github.com/navilg/k8senv/internal/install"
	"github.com/navilg/k8senv/internal/use"
	"github.com/spf13/cobra"
)

var kubectlCmd = &cobra.Command{
	Use:   "kubectl",
	Short: "Install, Use or List versions of kubectl",
	// Long:  ``,

	// Run: func(cmd *cobra.Command, args []string) {
	// 	fmt.Println("kubectl called")
	// },
}

// kubectlCmd represents the kubectl command
var installKubectlCmd = &cobra.Command{
	Use:   "kubectl VERSION",
	Short: "Install a version of kubectl",
	Long: `Install a version of kubectl client
	
Examples:
	# Install kubectl version 1.26.2
	k8senv install kubectl v1.26.2

	# Install latest available stable version of kubectl
	k8senv install kubectl latest

	# Install kubectl version 1.20.0 and overwrite it if it already exists
	k8senv install kubectl 1.20.0 --overwrite

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
		_ = install.InstallKubectl(args[0], overwriteInstall)

	},
}

var useKubectlCmd = &cobra.Command{
	Use:   "kubectl VERSION",
	Short: "Switch to a version of kubectl",
	Long: `Switch to a version of kubectl client
	
Examples:
	# Switch to kubectl version 1.26.2
	k8senv use kubectl v1.26.2
	k8senv use kubectl 1.26.2

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
		_ = use.UseKubectl(args[0])
	},
}

func init() {
	rootCmd.AddCommand(kubectlCmd)
	installCmd.AddCommand(installKubectlCmd)
	useCmd.AddCommand(useKubectlCmd)

	// Here you will define your flags and configuration settings.

	// Cobra supports Persistent Flags which will work for this command
	// and all subcommands, e.g.:
	// kubectlCmd.PersistentFlags().String("foo", "", "A help for foo")
	kubectlInstallCmd.PersistentFlags().BoolVarP(&overwriteInstall, "overwrite", "f", false, "Overwrite or re-install existing version")

	// Cobra supports local flags which will only run when this command
	// is called directly, e.g.:
	// kubectlCmd.Flags().BoolP("toggle", "t", false, "Help message for toggle")
}
