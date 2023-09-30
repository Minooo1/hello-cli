/*
Copyright © 2023 NAME HERE <EMAIL ADDRESS>

*/
package cmd

import (
	"bytes"
	"fmt"

	"github.com/Minooo1/ascii-art/asciiArt"
	"github.com/spf13/cobra"
)

// squidCmd represents the stegen command
var squidCmd = &cobra.Command{
	Use:   "squid",
	Short: "Print the funny ascii art of squid",
	RunE: func(cmd *cobra.Command, args []string) error {
		file, err := asciiArt.AsciiArt.Open("squid.txt")

		if err != nil {
			fmt.Println(err)
			return err
		}

		defer file.Close()

		buf := new(bytes.Buffer)
		buf.ReadFrom(file)

		fmt.Print(buf.String())

		return nil
	},
}

func init() {
	rootCmd.AddCommand(squidCmd)
}
