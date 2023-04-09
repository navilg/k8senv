/*
Copyright © 2023 NAME HERE <EMAIL ADDRESS>
*/
package cmd

import (
	"fmt"
	"os"

	"github.com/navilg/k8senv/internal/remove"
	"github.com/spf13/cobra"
)

// removeCmd represents the remove command
var removeCmd = &cobra.Command{
	Use:   "remove",
	Short: "Remove a version of kubectl, helm or velero",
	// 	Long: `A longer description that spans multiple lines and likely contains examples
	// and usage of using your command. For example:

	// Cobra is a CLI library for Go that empowers applications.
	// This application is a tool to generate the needed files
	// to quickly create a Cobra application.`,
	//
	//	Run: func(cmd *cobra.Command, args []string) {
	//		fmt.Println("remove called")
	//	},
}

var kubectlRemoveCmd = &cobra.Command{
	Use:   "remove VERSION",
	Short: "Remove an installed versions of kubectl",
	Long: `Remove an installed versions of kubectl
	
Examples:
	k8senv kubectl remove v1.19.2`,
	Run: func(cmd *cobra.Command, args []string) {
		if len(args) != 1 {
			fmt.Println("Exactly one argumanet is required. Provide kubectl version to remove e.g. v1.20.3")
			os.Exit(1)
		}
		err := remove.RemoveKubectl(args[0])
		if err != nil {
			os.Exit(1)
		}
	},
}

var veleroRemoveCmd = &cobra.Command{
	Use:   "remove VERSION",
	Short: "Remove an installed versions of velero",
	Long: `Remove an installed versions of velero
	
Examples:
	k8senv velero remove v1.10.2`,
	Run: func(cmd *cobra.Command, args []string) {
		if len(args) != 1 {
			fmt.Println("Exactly one argumanet is required. Provide velero version to remove e.g. v1.10.2")
			os.Exit(1)
		}
		err := remove.RemoveVelero(args[0])
		if err != nil {
			os.Exit(1)
		}
	},
}

func init() {
	rootCmd.AddCommand(removeCmd)
	kubectlCmd.AddCommand(kubectlRemoveCmd)
	veleroCmd.AddCommand(veleroRemoveCmd)

	// Here you will define your flags and configuration settings.

	// Cobra supports Persistent Flags which will work for this command
	// and all subcommands, e.g.:
	// removeCmd.PersistentFlags().String("foo", "", "A help for foo")

	// Cobra supports local flags which will only run when this command
	// is called directly, e.g.:
	// removeCmd.Flags().BoolP("toggle", "t", false, "Help message for toggle")
}
