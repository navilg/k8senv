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

// installCmd represents the install command
var installCmd = &cobra.Command{
	Use:   "install",
	Short: "Install a version of kubectl, helm or velero",
	Long: `A longer description that spans multiple lines and likely contains examples
and usage of using your command. For example:

Cobra is a CLI library for Go that empowers applications.
This application is a tool to generate the needed files
to quickly create a Cobra application.`,
	Run: func(cmd *cobra.Command, args []string) {
		fmt.Println("install called")
	},
}

// installCmd represents the install command
var kubectlInstallCmd = &cobra.Command{
	Use:   "install",
	Short: "Install a version of kubectl, helm or velero",
	Long: `A longer description that spans multiple lines and likely contains examples
and usage of using your command. For example:

Cobra is a CLI library for Go that empowers applications.
This application is a tool to generate the needed files
to quickly create a Cobra application.`,
	Run: func(cmd *cobra.Command, args []string) {
		if len(args) != 1 {
			fmt.Println("Exactly one argumanet is required. Provide kubectl version to install e.g. v1.20.3")
			os.Exit(1)
		}
		_ = install.InstallKubectl(args[0], overwriteInstall)

	},
}

func init() {
	rootCmd.AddCommand(installCmd)
	kubectlCmd.AddCommand(kubectlInstallCmd)

	// Here you will define your flags and configuration settings.

	// Cobra supports Persistent Flags which will work for this command
	// and all subcommands, e.g.:
	// installCmd.PersistentFlags().String("foo", "", "A help for foo")

	// installKubectlCmd.PersistentFlags().BoolP("overwrite", "f", false, "Overwrite or re-install existing version")
	installKubectlCmd.PersistentFlags().BoolVarP(&overwriteInstall, "overwrite", "f", false, "Overwrite or re-install existing version")

	// Cobra supports local flags which will only run when this command
	// is called directly, e.g.:
	// installCmd.Flags().BoolP("toggle", "t", false, "Help message for toggle")
}
