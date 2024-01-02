/*
Copyright © 2024 NAME HERE <EMAIL ADDRESS>
*/
package cmd

import (
	"log"
	"os"
	"strconv"
	"syscall"

	"github.com/spf13/cobra"
)

// pkillCmd represents the pkill command
var pkillCmd = &cobra.Command{
	Use:   "pkill",
	Short: "特定のプロセスIDを持つプロセスを終了させる。",
	Long:  ``,
	Run: func(cmd *cobra.Command, args []string) {
		pid, err := strconv.Atoi(os.Args[2])
		if err != nil {
			log.Fatal(err)
		}
		proc, err := os.FindProcess(pid)
		if err != nil {
			log.Fatal(err)
		}
		err = proc.Signal(syscall.SIGTERM)
		if err != nil {
			log.Fatal(err)
		}
	},
}

func init() {
	rootCmd.AddCommand(pkillCmd)
}
