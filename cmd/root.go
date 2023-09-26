/*
Copyright © 2023 NAME HERE <EMAIL ADDRESS>

*/
package cmd

import (
	"bytes"
	"fmt"
	"os"

	"github.com/Minooo1/ascii-art/asciiArt"
	"github.com/spf13/cobra"
	"golang.org/x/term"
)

// rootCmd represents the base command when called without any subcommands
var rootCmd = &cobra.Command{
	Use:   "ascii-art",
	Short: "This is a command to display various Ascii art.",
	Run: func(cmd *cobra.Command, args []string) {
		argsCount := len(os.Args)

		if argsCount == 1 {
			// main.go の時だけ実行する
			file, err := asciiArt.AsciiArt.Open("welcome.txt")

			if err != nil {
				fmt.Println(err)
			}

			defer file.Close()

			buf := new(bytes.Buffer)
			buf.ReadFrom(file)

			fmt.Print(buf.String())
		}
	},
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

	if err != nil {
		os.Exit(1)
	}
}

func init() {
	rootCmd.Flags().BoolP("toggle", "t", false, "Help message for toggle")
}
