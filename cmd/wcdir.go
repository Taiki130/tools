/*
Copyright © 2023 NAME HERE <EMAIL ADDRESS>
*/
package cmd

import (
	"fmt"
	"log"
	"os"
	"path/filepath"

	"github.com/spf13/cobra"
)

// wcdirCmd represents the wcdir command
var wcdirCmd = &cobra.Command{
	Use:   "wcdir",
	Short: "指定したディレクトリ内のファイルとディレクトリの数をカウントし、それぞれの総数を表示する。",
	Long:  ``,
	Run: func(cmd *cobra.Command, args []string) {
		var fCount int
		var dCount int
		err := filepath.Walk("../", func(path string, info os.FileInfo, err error) error {
			if err != nil {
				return err
			}
			if info.IsDir() {
				dCount++
			} else {
				fCount++
			}
			return nil
		})
		if err != nil {
			log.Fatal(err)
		}
		fmt.Printf("Directory: %d, File: %d", dCount, fCount)
	},
}

func init() {
	rootCmd.AddCommand(wcdirCmd)

	// Here you will define your flags and configuration settings.

	// Cobra supports Persistent Flags which will work for this command
	// and all subcommands, e.g.:
	// wcdirCmd.PersistentFlags().String("foo", "", "A help for foo")

	// Cobra supports local flags which will only run when this command
	// is called directly, e.g.:
	// wcdirCmd.Flags().BoolP("toggle", "t", false, "Help message for toggle")
}
