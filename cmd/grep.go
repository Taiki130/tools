/*
Copyright © 2023 NAME HERE <EMAIL ADDRESS>
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

// grepCmd represents the grep command
var grepCmd = &cobra.Command{
	Use:   "grep",
	Short: "テキストファイルから特定のパターンに一致する行だけを抽出し、出力する",
	Long: `テキストファイルから特定のパターンに一致する行だけを抽出し、出力する。
	usage:
		grep <pattern> <file>`,
	Run: func(cmd *cobra.Command, args []string) {
		file := ""
		pattern := ""

		if len(args) == 2 {
			file = args[0]
			pattern = args[1]
		} else if len(args) == 1 {
			log.Fatal("ファイル名を指定してください")
		} else {
			log.Fatal("引数の数が正しくありません")
		}

		fp, err := os.Open(file)
		if err != nil {
			log.Fatal(err)
		}
		defer fp.Close()

		scanner := bufio.NewScanner(fp)
		for scanner.Scan() {
			line := scanner.Text()
			if strings.Contains(line, pattern) {
				fmt.Println(line)
			}
		}
	},
}

func init() {
	rootCmd.AddCommand(grepCmd)
}
