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

// veleroCmd represents the velero command
var veleroCmd = &cobra.Command{
	Use:   "velero",
	Short: "Install, Use or List versions of velero",
	// 	Long: `A longer description that spans multiple lines and likely contains examples
	// and usage of using your command. For example:

	// Cobra is a CLI library for Go that empowers applications.
	// This application is a tool to generate the needed files
	// to quickly create a Cobra application.`,
	// Run: func(cmd *cobra.Command, args []string) {
	// 	fmt.Println("velero called")
	// },
}

var installVeleroCmd = &cobra.Command{
	Use:   "velero VERSION",
	Short: "Install a version of velero client",
	Long: `Install a version of velero client
	
Examples:
	# Install velero client version 1.10.2
	k8senv install velero v1.10.2

	# Install latest available stable version of velero client
	k8senv install velero latest

	# Install velero client version 1.8.0 and overwrite it if it already exists
	k8senv install velero 1.8.0 --overwrite

	# Install velero client version 1.8.0 aand set timeout to 300 seconds (If internet is slow), Default: 120 seconds
	k8senv install velero 1.8.0 --timeout=300

Supported version formats:
	v1.10.2
	1.10.2	# Defaults to v1.10.2
	1.10 	# Defaults to v1.10.0
	1 	# Defaults to v1.0.0`,

	Run: func(cmd *cobra.Command, args []string) {
		if len(args) != 1 {
			fmt.Println("Exactly one argumanet is required. Provide velero client version to install e.g. v1.10.2")
			os.Exit(1)
		}
		err := install.InstallVelero(args[0], overwriteInstall, timeout, proxy)
		if err != nil {
			os.Exit(1)
		}

	},
}

var useVeleroCmd = &cobra.Command{
	Use:   "velero VERSION",
	Short: "Switch to a version of velero client",
	Long: `Switch to a version of velero client
	
Examples:
	# Switch to velero client version 1.10.2
	k8senv use velero v1.10.2
	k8senv use velero 1.10.2

Supported version formats:
	v1.10.2
	1.10.2	# Defaults to v1.10.2
	1.10 	# Defaults to v1.10.0
	1 	# Defaults to v1.0.0`,
	Run: func(cmd *cobra.Command, args []string) {
		if len(args) != 1 {
			fmt.Println("Exactly one argumanet is required. Provide velero client version to use e.g. v1.10.2")
			os.Exit(1)
		}
		// err := use.UseVelero(args[0])
		// if err != nil {
		// 	os.Exit(1)
		// }
	},
}

var listVeleroCmd = &cobra.Command{
	Use:   "velero",
	Short: "List all installed versions of velero client",
	Long: `List all installed versions of velero client
	
Examples:
	k8senv list velero`,
	Run: func(cmd *cobra.Command, args []string) {
		if len(args) != 0 {
			fmt.Println("No argument is required for listing versions")
			os.Exit(1)
		}
		// err := list.ListVelero()
		// if err != nil {
		// 	os.Exit(1)
		// }
	},
}

var removeVeleroCmd = &cobra.Command{
	Use:   "remove VERSION",
	Short: "Remove an installed versions of velero client",
	Long: `Remove an installed versions of velero client
	
Examples:
	k8senv remove velero v1.10.2`,
	Run: func(cmd *cobra.Command, args []string) {
		if len(args) != 1 {
			fmt.Println("Exactly one argumanet is required. Provide velero client version to remove e.g. v1.10.2")
			os.Exit(1)
		}
		// err := remove.RemoveVelero(args[0])
		// if err != nil {
		// 	os.Exit(1)
		// }
	},
}

func init() {
	rootCmd.AddCommand(veleroCmd)
	installCmd.AddCommand(installVeleroCmd)
	useCmd.AddCommand(useVeleroCmd)
	listCmd.AddCommand(listVeleroCmd)
	removeCmd.AddCommand(removeVeleroCmd)

	// Here you will define your flags and configuration settings.

	// Cobra supports Persistent Flags which will work for this command
	// and all subcommands, e.g.:
	// veleroCmd.PersistentFlags().String("foo", "", "A help for foo")
	veleroInstallCmd.PersistentFlags().BoolVarP(&overwriteInstall, "overwrite", "f", false, "Overwrite or re-install existing version")
	veleroInstallCmd.PersistentFlags().IntVarP(&timeout, "timeout", "t", 120, "Timeout in seconds [DEFAULT: 120 seconds]")
	veleroInstallCmd.PersistentFlags().StringVarP(&proxy, "proxy", "p", "", "HTTP/HTTPS proxy to use for downloading clients from its source")

	// Cobra supports local flags which will only run when this command
	// is called directly, e.g.:
	// veleroCmd.Flags().BoolP("toggle", "t", false, "Help message for toggle")
}
