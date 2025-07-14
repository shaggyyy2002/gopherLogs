/*
Copyright © 2025 NAME HERE <EMAIL ADDRESS>
*/
package cmd

import (
	"github.com/shaggyyy2002/gopherlogs/functions"
	"github.com/spf13/cobra"
)

// filterCmd represents the filter command
var filterCmd = &cobra.Command{
	Use:   "filter",
	Short: "A brief description of your command",
	Long: `A longer description that spans multiple lines and likely contains examples
and usage of using your command. For example:

Cobra is a CLI library for Go that empowers applications.
This application is a tool to generate the needed files
to quickly create a Cobra application.`,
	Run: func(cmd *cobra.Command, args []string) {
		filepath, _ := cmd.Flags().GetString("FILE")
		keyword, _ := cmd.Flags().GetString("ADD KEYWORD")
		functions.Filter(filepath, keyword)
	},
}

func init() {
	filterCmd.Flags().StringP("file", "f", "", "Path to the log file to filter")
	filterCmd.Flags().StringP("keyword", "k", "", "Keyword to filter in the log file")

	rootCmd.AddCommand(filterCmd)

	// Here you will define your flags and configuration settings.

	// Cobra supports Persistent Flags which will work for this command
	// and all subcommands, e.g.:
	// filterCmd.PersistentFlags().String("foo", "", "A help for foo")

	// Cobra supports local flags which will only run when this command
	// is called directly, e.g.:
	// filterCmd.Flags().BoolP("toggle", "t", false, "Help message for toggle")
}
