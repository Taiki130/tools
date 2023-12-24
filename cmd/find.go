/*
Copyright © 2023 NAME HERE <EMAIL ADDRESS>
*/
package cmd

import (
	"fmt"
	"os"
	"log"
	"path/filepath"

	"github.com/spf13/cobra"
)

// findCmd represents the find command
var findCmd = &cobra.Command{
	Use:   "find",
	Short: "特定のディレクトリ内の全てのファイルとディレクトリを再帰的にリスト表示する。",
	Long: `A longer description that spans multiple lines and likely contains examples
and usage of using your command. For example:

Cobra is a CLI library for Go that empowers applications.
This application is a tool to generate the needed files
to quickly create a Cobra application.`,
	Run: func(cmd *cobra.Command, args []string) {
		err := filepath.Walk("../", func(path string, info os.FileInfo, err error) error {
			if err != nil {
				return err
			}
			fmt.Println(path)
			return nil
		})
		if err != nil {
			log.Fatal(err)
		}
	},
}

func init() {
	rootCmd.AddCommand(findCmd)
}
