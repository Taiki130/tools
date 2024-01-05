/*
Copyright © 2024 NAME HERE <EMAIL ADDRESS>
*/
package cmd

import (
	"fmt"
	"log"
	"os"
	"path/filepath"

	"github.com/spf13/cobra"
)

type FileSize struct {
	Path string
	Size int64
}

// fileSizeCmd represents the fileSize command
var fileSizeCmd = &cobra.Command{
	Use:   "fileSize",
	Short: "指定したディレクトリ内のファイルを再帰的にリストアップし、各ファイルのサイズを表示する",
	Long:  ``,
	Run: func(cmd *cobra.Command, args []string) {
		dir := os.Args[2]
		err := filepath.Walk(dir, func(path string, info os.FileInfo, err error) error {
			if err != nil {
				return err
			}
			if info.IsDir() {
				return nil
			}
			fileSize := FileSize{
				Path: path,
				Size: info.Size(),
			}
			fmt.Printf("path: %s, size: %d bytes\n", fileSize.Path, fileSize.Size)
			return nil
		})
		if err != nil {
			log.Fatal(err)
		}
	},
}

func init() {
	rootCmd.AddCommand(fileSizeCmd)

	// Here you will define your flags and configuration settings.

	// Cobra supports Persistent Flags which will work for this command
	// and all subcommands, e.g.:
	// fileSizeCmd.PersistentFlags().String("foo", "", "A help for foo")

	// Cobra supports local flags which will only run when this command
	// is called directly, e.g.:
	// fileSizeCmd.Flags().BoolP("toggle", "t", false, "Help message for toggle")
}
