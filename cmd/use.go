/*
Copyright © 2023 NAME HERE <EMAIL ADDRESS>
*/
package cmd

import (
	"fmt"
	"os"

	"github.com/navilg/k8senv/internal/use"
	"github.com/spf13/cobra"
)

// useCmd represents the use command
var useCmd = &cobra.Command{
	Use:   "use",
	Short: "Switch to a version of kubectl, helm or velero",
	// 	Long: `A longer description that spans multiple lines and likely contains examples
	// and usage of using your command. For example:

	// Cobra is a CLI library for Go that empowers applications.
	// This application is a tool to generate the needed files
	// to quickly create a Cobra application.`,
	// Run: func(cmd *cobra.Command, args []string) {
	// 	fmt.Println("use called")
	// },
}

var kubectUseCmd = &cobra.Command{
	Use:   "use VERSION",
	Short: "Switch to a version of kubectl client",
	Long: `Switch to a version of kubectl client
	
Examples:
	# Switch to kubectl version 1.26.2
	k8senv kubectl use v1.26.2
	k8senv kubectl use 1.26.2
	
Supported version formats:
	v1.20.3
	1.20.3	# Defaults to v1.20.3
	1.20 	# Defaults to v1.20.0
	1 	# Defaults to v1.0.0`,
	Run: func(cmd *cobra.Command, args []string) {
		if len(args) != 1 {
			fmt.Println("Exactly one argumanet is required. Provide kubectl version to use e.g. v1.20.3")
			os.Exit(1)
		}
		err := use.UseVersion("kubectl", args[0])
		if err != nil {
			os.Exit(1)
		}
	},
}

var veleroUseCmd = &cobra.Command{
	Use:   "use VERSION",
	Short: "Switch to a version of velero client",
	Long: `Switch to a version of velero client
	
Examples:
	# Switch to velero version 1.10.2
	k8senv velero use v1.10.2
	k8senv velero use 1.10.2
	
Supported version formats:
	v1.10.2
	1.10.2	# Defaults to v1.10.2
	1.10 	# Defaults to v1.10.0
	1 	# Defaults to v1.0.0`,
	Run: func(cmd *cobra.Command, args []string) {
		if len(args) != 1 {
			fmt.Println("Exactly one argumanet is required. Provide velero version to use e.g. v1.10.2")
			os.Exit(1)
		}
		err := use.UseVersion("velero", args[0])
		if err != nil {
			os.Exit(1)
		}
	},
}

var helmUseCmd = &cobra.Command{
	Use:   "use VERSION",
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

func init() {
	rootCmd.AddCommand(useCmd)
	kubectlCmd.AddCommand(kubectUseCmd)
	veleroCmd.AddCommand(veleroUseCmd)
	helmCmd.AddCommand(helmUseCmd)

	// Here you will define your flags and configuration settings.

	// Cobra supports Persistent Flags which will work for this command
	// and all subcommands, e.g.:
	// useCmd.PersistentFlags().String("foo", "", "A help for foo")

	// Cobra supports local flags which will only run when this command
	// is called directly, e.g.:
	// useCmd.Flags().BoolP("toggle", "t", false, "Help message for toggle")
}
