/*
Copyright © 2023 NAME HERE <EMAIL ADDRESS>
*/
package cmd

import (
	"bufio"
	"fmt"
	"log"
	"os"

	"github.com/spf13/cobra"
)

// catCmd represents the cat command
var catCmd = &cobra.Command{
	Use:   "cat",
	Short: "cat – concatenate and print files",
	Long:  `The cat utility reads files sequentially, writing them to the standard output.`,
	Run: func(cmd *cobra.Command, args []string) {
		if len(args) != 1 {
			log.Fatal("The argument must be one.")
			return
		}
		file, err := os.Open(args[0])
		if err != nil {
			log.Fatal(err)
		}
		defer file.Close()
		scanner := bufio.NewScanner(file)
		var lines []string
		for scanner.Scan() {
			lines = append(lines, scanner.Text())
		}
		for i := 0; i < len(lines); i++ {
			fmt.Println(lines[i])
		}
	},
}

func init() {
	rootCmd.AddCommand(catCmd)
}
