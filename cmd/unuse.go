/*
Copyright © 2023 NAME HERE <EMAIL ADDRESS>
*/
package cmd

import (
	"fmt"
	"os"

	"github.com/navilg/k8senv/internal/unuse"
	"github.com/spf13/cobra"
)

// unuseCmd represents the unuse command
var unuseCmd = &cobra.Command{
	Use:   "unuse",
	Short: "Stop using k8senv managed kubectl, helm or velero",
	// 	Long: `A longer description that spans multiple lines and likely contains examples
	// and usage of using your command. For example:

	// Cobra is a CLI library for Go that empowers applications.
	// This application is a tool to generate the needed files
	// to quickly create a Cobra application.`,
	//
	//	Run: func(cmd *cobra.Command, args []string) {
	//		fmt.Println("unuse called")
	//	},
}

var kubectlUnuseCmd = &cobra.Command{
	Use:   "unuse",
	Short: "Stop using k8senv managed kubectl",
	Long: `Stop using k8senv managed kubectl. 
This will reset your system to use system installed client of kubectl if present.
	
Examples:
	k8senv kubectl unuse`,
	Run: func(cmd *cobra.Command, args []string) {
		if len(args) != 0 {
			fmt.Println("No argument is required for unusing kubectl")
			os.Exit(1)
		}
		err := unuse.UnuseVersions("kubectl")
		if err != nil {
			os.Exit(1)
		}
	},
}

var veleroUnuseCmd = &cobra.Command{
	Use:   "unuse",
	Short: "Stop using k8senv managed velero",
	Long: `Stop using k8senv managed velero. 
This will reset your system to use system installed client of velero if present.
	
Examples:
	k8senv velero unuse`,
	Run: func(cmd *cobra.Command, args []string) {
		if len(args) != 0 {
			fmt.Println("No argument is required for unusing velero")
			os.Exit(1)
		}
		err := unuse.UnuseVersions("velero")
		if err != nil {
			os.Exit(1)
		}
	},
}

var helmUnuseCmd = &cobra.Command{
	Use:   "unuse",
	Short: "Stop using k8senv managed helm",
	Long: `Stop using k8senv managed helm. 
This will reset your system to use system installed client of helm if present.
	
Examples:
	k8senv helm unuse`,
	Run: func(cmd *cobra.Command, args []string) {
		if len(args) != 0 {
			fmt.Println("No argument is required for unusing helm")
			os.Exit(1)
		}
		err := unuse.UnuseVersions("helm")
		if err != nil {
			os.Exit(1)
		}
	},
}

func init() {
	rootCmd.AddCommand(unuseCmd)
	kubectlCmd.AddCommand(kubectlUnuseCmd)
	veleroCmd.AddCommand(veleroUnuseCmd)
	helmCmd.AddCommand(helmUnuseCmd)

	// Here you will define your flags and configuration settings.

	// Cobra supports Persistent Flags which will work for this command
	// and all subcommands, e.g.:
	// unuseCmd.PersistentFlags().String("foo", "", "A help for foo")

	// Cobra supports local flags which will only run when this command
	// is called directly, e.g.:
	// unuseCmd.Flags().BoolP("toggle", "t", false, "Help message for toggle")
}
