/*
Copyright © 2024 NAME HERE <EMAIL ADDRESS>
*/
package cmd

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"strings"

	"github.com/spf13/cobra"
)

// grepnCmd represents the grepn command
var grepnCmd = &cobra.Command{
	Use:   "grepn",
	Short: "ファイルの中で特定の文字列を検索し、その文字列が現れた行番号と行を表示する。",
	Long:  ``,
	Run: func(cmd *cobra.Command, args []string) {
		f, err := os.Open("examples/example.txt")
		if err != nil {
			log.Fatal(err)
		}
		defer f.Close()
		s := bufio.NewScanner(f)
		count := 0
		for s.Scan() {
			count++
			if strings.Contains(s.Text(), "about.html") {
				fmt.Printf("%d: %s\n", count, s.Text())
			}
		}
		if s.Err() != nil {
			log.Fatal(s.Err())
		}
	},
}

func init() {
	rootCmd.AddCommand(grepnCmd)
}
