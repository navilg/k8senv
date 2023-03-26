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

var kubectlCmd = &cobra.Command{
	Use:   "kubectl",
	Short: "A brief description of your command",
	Long: `A longer description that spans multiple lines and likely contains examples
and usage of using your command. For example:

Cobra is a CLI library for Go that empowers applications.
This application is a tool to generate the needed files
to quickly create a Cobra application.`,

	Run: func(cmd *cobra.Command, args []string) {
		fmt.Println("kubectl called")
	},
}

// kubectlCmd represents the kubectl command
var installKubectlCmd = &cobra.Command{
	Use:   "kubectl",
	Short: "A brief description of your command",
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

var useKubectlCmd = &cobra.Command{
	Use:   "kubectl",
	Short: "A brief description of your command",
	Long: `A longer description that spans multiple lines and likely contains examples
and usage of using your command. For example:

Cobra is a CLI library for Go that empowers applications.
This application is a tool to generate the needed files
to quickly create a Cobra application.`,
	Run: func(cmd *cobra.Command, args []string) {
		fmt.Println("use kubectl called")
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
