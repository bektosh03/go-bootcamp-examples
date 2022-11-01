/*
Copyright © 2022 NAME HERE <EMAIL ADDRESS>
*/
package cmd

import (
	"strconv"

	"github.com/spf13/cobra"
)

// doubleCmd represents the double command
var doubleCmd = &cobra.Command{
	Use:   "double",
	Short: "A brief description of your command",
	Long: `A longer description that spans multiple lines and likely contains examples
and usage of using your command. For example:

Cobra is a CLI library for Go that empowers applications.
This application is a tool to generate the needed files
to quickly create a Cobra application.`,
	Run: func(cmd *cobra.Command, args []string) {
		if len(args) != 1 {
			cmd.Println("expected exactly one argument")
			return
		}

		n, err := strconv.Atoi(args[0])
		if err != nil {
			cmd.Println("expected a number")
			return
		}

		cmd.Println("result:", n*2)
	},
}

func init() {
	rootCmd.AddCommand(doubleCmd)

	// Here you will define your flags and configuration settings.

	// Cobra supports Persistent Flags which will work for this command
	// and all subcommands, e.g.:
	// doubleCmd.PersistentFlags().String("foo", "", "A help for foo")

	// Cobra supports local flags which will only run when this command
	// is called directly, e.g.:
	// doubleCmd.Flags().BoolP("toggle", "t", false, "Help message for toggle")
}
