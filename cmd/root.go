/*
Copyright © 2023 NAME HERE <EMAIL ADDRESS>

*/
package cmd

import (
	"fmt"
	"os"

	"github.com/spf13/cobra"
)

// rootCmd represents the base command when called without any subcommands
var rootCmd = &cobra.Command{
	Use:   "ascii-art",
	Short: "A brief description of your application",
}

// Execute adds all child commands to the root command and sets flags appropriately.
// This is called by main.main(). It only needs to happen once to the rootCmd.
func Execute() {
	err := rootCmd.Execute()

	argsCount := len(os.Args)

	if argsCount == 1 {
		// main.go の時だけ実行する
		b, err := os.ReadFile("askiiArt/welcome.txt")

		if err != nil {
			fmt.Println(err)
		}

		fmt.Println(string(b))
	}

	if err != nil {
		os.Exit(1)
	}

}

func init() {
	rootCmd.Flags().BoolP("toggle", "t", false, "Help message for toggle")
}
