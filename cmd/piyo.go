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

// piyoCmd represents the piyo command
var piyoCmd = &cobra.Command{
	Use:   "piyo",
	Short: "Print the funny ascii art of piyo",
	RunE: func(cmd *cobra.Command, args []string) error {
		file, err := asciiArt.AsciiArt.Open("piyo.txt")

		if err != nil {
			fmt.Println(err)
			return err
		}

		buf := new(bytes.Buffer)
		buf.ReadFrom(file)
		fmt.Print(buf.String())

		defer file.Close()

		return nil
	},
}

func init() {
	rootCmd.AddCommand(piyoCmd)

}
