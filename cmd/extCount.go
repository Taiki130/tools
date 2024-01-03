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

type Info struct {
	Path string
	Ext  string
}

// extCountCmd represents the extCount command
var extCountCmd = &cobra.Command{
	Use:   "extCount",
	Short: "指定したディレクトリ内のファイルをリストアップし、各ファイルの拡張子ごとにファイル数を表示する。",
	Long:  ``,
	Run: func(cmd *cobra.Command, args []string) {
		targetDir := os.Args[2]
		fileInfos, err := listFileInfos(targetDir)
		if err != nil {
			log.Fatal(err)
		}

		extCount := make(map[string]int)

		for _, fileInfo := range fileInfos {
			extCount[fileInfo.Ext]++
		}

		for ext, count := range extCount {
			fmt.Printf("Ext: %s, Count: %d\n", ext, count)
		}
	},
}

func listFileInfos(d string) ([]Info, error) {
	var fileInfos []Info
	err := filepath.Walk(d, func(path string, info os.FileInfo, err error) error {
		if err != nil {
			return err
		}
		if !info.IsDir() {
			ext := filepath.Ext(info.Name())
			if ext == "" {
				ext = "other"
			}
			fileInfos = append(fileInfos, Info{
				Path: path,
				Ext:  ext,
			})
		}
		return nil
	})
	return fileInfos, err
}

func init() {
	rootCmd.AddCommand(extCountCmd)

	// Here you will define your flags and configuration settings.

	// Cobra supports Persistent Flags which will work for this command
	// and all subcommands, e.g.:
	// extCountCmd.PersistentFlags().String("foo", "", "A help for foo")

	// Cobra supports local flags which will only run when this command
	// is called directly, e.g.:
	// extCountCmd.Flags().BoolP("toggle", "t", false, "Help message for toggle")
}
