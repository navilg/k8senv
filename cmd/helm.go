/*
Copyright © 2023 NAME HERE <EMAIL ADDRESS>
*/
package cmd

import (
	"fmt"
	"os"

	"github.com/navilg/k8senv/internal/install"
	"github.com/navilg/k8senv/internal/list"
	"github.com/navilg/k8senv/internal/remove"
	"github.com/navilg/k8senv/internal/use"
	"github.com/spf13/cobra"
)

// helmCmd represents the helm command
var helmCmd = &cobra.Command{
	Use:   "helm",
	Short: "Install, Use or List versions of helm",
	// 	Long: `A longer description that spans multiple lines and likely contains examples
	// and usage of using your command. For example:

	// Cobra is a CLI library for Go that empowers applications.
	// This application is a tool to generate the needed files
	// to quickly create a Cobra application.`,
	//
	//	Run: func(cmd *cobra.Command, args []string) {
	//		fmt.Println("helm called")
	//	},
}

var installHelmCmd = &cobra.Command{
	Use:   "helm VERSION",
	Short: "Install a version of helm",
	Long: `Install a version of helm
	
Examples:
	# Install helm version 3.10.2
	k8senv install helm v3.10.2

	# Install latest available stable version of helm
	k8senv install helm latest

	# Install helm version 3.8.0 and overwrite it if it already exists
	k8senv install helm 3.8.0 --overwrite

	# Install helm version 1.8.0 aand set timeout to 300 seconds (If internet is slow), Default: 120 seconds
	k8senv install helm 3.8.0 --timeout=300

Supported version formats:
	v3.10.2
	3.10.2	# Defaults to v3.10.2
	3.10 	# Defaults to v3.10.0
	3 	# Defaults to v3.0.0`,

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

var useHelmCmd = &cobra.Command{
	Use:   "helm VERSION",
	Short: "Switch to a version of helm",
	Long: `Switch to a version of helm
	
Examples:
	# Switch to helm version 3.10.2
	k8senv use helm v3.10.2
	k8senv use helm 3.10.2
	
Supported version formats:
	v3.10.2
	3.10.2	# Defaults to v3.10.2
	3.10 	# Defaults to v3.10.0
	3 	# Defaults to v3.0.0`,
	Run: func(cmd *cobra.Command, args []string) {
		if len(args) != 1 {
			fmt.Println("Exactly one argumanet is required. Provide helm version to use e.g. v3.10.2")
			os.Exit(1)
		}
		err := use.UseVersion("helm", args[0])
		if err != nil {
			os.Exit(1)
		}
	},
}

var listHelmCmd = &cobra.Command{
	Use:   "helm",
	Short: "List all installed versions of helm",
	Long: `List all installed versions of helm
	
Examples:
	k8senv list helm`,
	Run: func(cmd *cobra.Command, args []string) {
		if len(args) != 0 {
			fmt.Println("No argument is required for listing versions")
			os.Exit(1)
		}
		err := list.ListVersions("helm")
		if err != nil {
			os.Exit(1)
		}
	},
}

var removeHelmCmd = &cobra.Command{
	Use:   "helm VERSION",
	Short: "Remove an installed versions of helm",
	Long: `Remove an installed versions of helm
	
Examples:
	k8senv remove helm v3.10.2`,
	Run: func(cmd *cobra.Command, args []string) {
		if len(args) != 1 {
			fmt.Println("Exactly one argumanet is required. Provide helm version to remove e.g. v3.10.2")
			os.Exit(1)
		}
		err := remove.RemoveHelm(args[0])
		if err != nil {
			os.Exit(1)
		}
	},
}

func init() {
	rootCmd.AddCommand(helmCmd)
	installCmd.AddCommand(installHelmCmd)
	useCmd.AddCommand(useHelmCmd)
	listCmd.AddCommand(listHelmCmd)
	removeCmd.AddCommand(removeHelmCmd)

	// Here you will define your flags and configuration settings.

	// Cobra supports Persistent Flags which will work for this command
	// and all subcommands, e.g.:
	// helmCmd.PersistentFlags().String("foo", "", "A help for foo")
	helmInstallCmd.PersistentFlags().BoolVarP(&overwriteInstall, "overwrite", "f", false, "Overwrite or re-install existing version")
	helmInstallCmd.PersistentFlags().IntVarP(&timeout, "timeout", "t", 120, "Timeout in seconds [DEFAULT: 120 seconds]")
	helmInstallCmd.PersistentFlags().StringVarP(&proxy, "proxy", "p", "", "HTTP/HTTPS proxy to use for downloading clients from its source")

	// Cobra supports local flags which will only run when this command
	// is called directly, e.g.:
	// helmCmd.Flags().BoolP("toggle", "t", false, "Help message for toggle")
}
