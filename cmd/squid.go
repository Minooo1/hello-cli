/*
Copyright © 2023 NAME HERE <EMAIL ADDRESS>

*/
package cmd

import (
	"fmt"
	"os"

	"github.com/spf13/cobra"
)

// squidCmd represents the stegen command
var squidCmd = &cobra.Command{
	Use:   "squid",
	Short: "Print the funny ascii art of squid",
	Run: func(cmd *cobra.Command, args []string) {
		b, err := os.ReadFile("asciiArt/squid.txt")

		if err != nil {
			fmt.Println(err)
		}

		fmt.Print(string(b))
	},
}

func init() {
	rootCmd.AddCommand(squidCmd)
}
