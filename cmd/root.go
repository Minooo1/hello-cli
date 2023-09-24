/*
Copyright © 2023 NAME HERE <EMAIL ADDRESS>

*/
package cmd

import (
	"fmt"
	"os"

	"github.com/spf13/cobra"
	"golang.org/x/term"
)

// rootCmd represents the base command when called without any subcommands
var rootCmd = &cobra.Command{
	Use:   "ascii-art",
	Short: "This is a command to display various Ascii art.",
}

func GetWindowSize(fd int) (width, height int, err error) {
	width, height, err = term.GetSize(fd)
	return width, height, err
}

// Execute adds all child commands to the root command and sets flags appropriately.
// This is called by main.main(). It only needs to happen once to the rootCmd.
func Execute() {
	err := rootCmd.Execute()

	// width, _, err := GetWindowSize(int(os.Stdin.Fd()))

	argsCount := len(os.Args)

	if argsCount == 1 {
		// main.go の時だけ実行する
		b, err := os.ReadFile("asciiArt/welcome.txt")
		fmt.Println(string(b))

		if err != nil {
			os.Exit(1)
		}
	}

	if err != nil {
		os.Exit(1)
	}
}

func init() {
	rootCmd.Flags().BoolP("toggle", "t", false, "Help message for toggle")
}
