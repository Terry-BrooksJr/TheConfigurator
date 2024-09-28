/*
Copyright © 2024 NAME HERE <EMAIL ADDRESS>
*/
package cmd

import (
	"fmt"

	"github.com/spf13/cobra"
)

// resetCmd represents the reset command
var resetCmd = &cobra.Command{
	Use:   "reset",
	Short: "Erase all Past SetUps",
	Long:  `"This command will reverse all of the previously completed Server Configuration Tasks"`,
	Run: func(cmd *cobra.Command, args []string) {
		fmt.Println("reset called")
	},
}

func init() {
	rootCmd.AddCommand(resetCmd)

	// Here you will define your flags and configuration settings.

	// Cobra supports Persistent Flags which will work for this command
	// and all subcommands, e.g.:
	// resetCmd.PersistentFlags().String("foo", "", "A help for foo")

	// Cobra supports local flags which will only run when this command
	// is called directly, e.g.:
	// resetCmd.Flags().BoolP("toggle", "t", false, "Help message for toggle")
}
