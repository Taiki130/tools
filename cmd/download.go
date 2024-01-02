/*
Copyright © 2024 NAME HERE <EMAIL ADDRESS>
*/
package cmd

import (
	"io"
	"log"
	"net/http"
	"os"
	"sync"

	"github.com/spf13/cobra"
)

// downloadCmd represents the download command
var downloadCmd = &cobra.Command{
	Use:   "download",
	Short: "指定した複数のURLからデータをダウンロードする。",
	Long:  ``,
	Run: func(cmd *cobra.Command, args []string) {
		urls := []string{
			"https://sreake.com/service-sre/",
			"https://sreake.com/blog/5point-good-postmortem/",
			"https://sreake.com/blog/what-is-sre/",
		}
		var wg sync.WaitGroup
		for _, url := range urls {
			wg.Add(1)
			go func(url string) {
				defer wg.Done()
				download(url)
			}(url)
		}
		wg.Wait()
	},
}

func download(url string) {
	resp, err := http.Get(url)
	if err != nil {
		log.Fatal(err)
	}
	defer resp.Body.Close()
	if resp.StatusCode != http.StatusOK {
		log.Fatal(resp.StatusCode)
	}
	_, err = io.Copy(os.Stdout, resp.Body)
	if err != nil {
		log.Fatal(err)
	}
}
func init() {
	rootCmd.AddCommand(downloadCmd)
}
