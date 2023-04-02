/*
Copyright © 2023 NAME HERE <EMAIL ADDRESS>
*/
package cmd

import (
	"fmt"
	"os"

	"github.com/navilg/k8senv/internal/list"
	"github.com/spf13/cobra"
)

// listCmd represents the list command
var listCmd = &cobra.Command{
	Use:   "list",
	Short: "List all installed versions of kubectl, helm or velero",
	// 	Long: `A longer description that spans multiple lines and likely contains examples
	// and usage of using your command. For example:

	// Cobra is a CLI library for Go that empowers applications.
	// This application is a tool to generate the needed files
	// to quickly create a Cobra application.`,
	//
	//	Run: func(cmd *cobra.Command, args []string) {
	//		fmt.Println("list called")
	//	},
}

var kubectlListCmd = &cobra.Command{
	Use:   "list",
	Short: "List all installed versions of kubectl",
	Long: `List all installed versions of kubectl
	
Examples:
	k8senv kubectl list`,
	Run: func(cmd *cobra.Command, args []string) {
		if len(args) != 0 {
			fmt.Println("No argument is required for listing versions")
			os.Exit(1)
		}
		err := list.ListKubectl()
		if err != nil {
			os.Exit(1)
		}
	},
}

func init() {
	rootCmd.AddCommand(listCmd)
	kubectlCmd.AddCommand(kubectlListCmd)

	// Here you will define your flags and configuration settings.

	// Cobra supports Persistent Flags which will work for this command
	// and all subcommands, e.g.:
	// listCmd.PersistentFlags().String("foo", "", "A help for foo")

	// Cobra supports local flags which will only run when this command
	// is called directly, e.g.:
	// listCmd.Flags().BoolP("toggle", "t", false, "Help message for toggle")
}
