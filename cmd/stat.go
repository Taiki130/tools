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

type FileInfo struct {
	Path string
	Size int64
}

// statCmd represents the stat command
var statCmd = &cobra.Command{
	Use:   "stat",
	Short: "指定したディレクトリ内のファイルをリストアップし、各ファイルのサイズを表示する。",
	Long:  ``,
	Run: func(cmd *cobra.Command, args []string) {
		dir := os.Args[2]
		files, err := os.ReadDir(dir)
		if err != nil {
			log.Fatal(err)
		}
		for _, file := range files {
			if !file.IsDir() {
				info, err := file.Info()
				if err != nil {
					log.Fatal(err)
				}
				fileInfo := FileInfo{
					Path: filepath.Join(dir, info.Name()),
					Size: info.Size(),
				}
				fmt.Printf("Path: %s, Size: %d bytes\n", fileInfo.Path, fileInfo.Size)
			}
		}
	},
}

func init() {
	rootCmd.AddCommand(statCmd)

	// Here you will define your flags and configuration settings.

	// Cobra supports Persistent Flags which will work for this command
	// and all subcommands, e.g.:
	// statCmd.PersistentFlags().String("foo", "", "A help for foo")

	// Cobra supports local flags which will only run when this command
	// is called directly, e.g.:
	// statCmd.Flags().BoolP("toggle", "t", false, "Help message for toggle")
}
