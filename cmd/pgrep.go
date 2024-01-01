/*
Copyright © 2024 NAME HERE <EMAIL ADDRESS>
*/
package cmd

import (
	"bufio"
	"bytes"
	"fmt"
	"log"
	"os/exec"
	"strings"

	"github.com/spf13/cobra"
)

// pgrepCmd represents the pgrep command
var pgrepCmd = &cobra.Command{
	Use:   "pgrep",
	Short: "実行中のすべてのプロセスの一覧を表示し、特定のプロセス名をフィルタリングする。",
	Long:  ``,
	Run: func(cmd *cobra.Command, args []string) {
		r, err := exec.Command("ps", "aux").Output()
		if err != nil {
			log.Fatal(err)
		}
		s := bufio.NewScanner(bytes.NewReader(r))
		processName := "docker"
		for s.Scan() {
			if strings.Contains(s.Text(), processName) {
				fmt.Print(s.Text())
			}
		}
	},
}

func init() {
	rootCmd.AddCommand(pgrepCmd)
}
