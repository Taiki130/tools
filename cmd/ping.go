/*
Copyright © 2024 NAME HERE <EMAIL ADDRESS>
*/
package cmd

import (
	"fmt"
	"log"
	"os"

	"github.com/go-ping/ping"
	"github.com/spf13/cobra"
)

// pingCmd represents the ping command
var pingCmd = &cobra.Command{
	Use:   "ping",
	Short: "指定したホスト名またはIPアドレスへのpingが成功するかどうかを確認する。",
	Long:  ``,
	Run: func(cmd *cobra.Command, args []string) {
		host := os.Args[2]
		pinger, err := ping.NewPinger(host)
		if err != nil {
			log.Fatal(err)
		}
		pinger.Count = 1
		err = pinger.Run()
		if err != nil {
			log.Fatal(err)
		}
		stats := pinger.Statistics()
		fmt.Print(stats)
	},
}

func init() {
	rootCmd.AddCommand(pingCmd)

	// Here you will define your flags and configuration settings.

	// Cobra supports Persistent Flags which will work for this command
	// and all subcommands, e.g.:
	// pingCmd.PersistentFlags().String("foo", "", "A help for foo")

	// Cobra supports local flags which will only run when this command
	// is called directly, e.g.:
	// pingCmd.Flags().BoolP("toggle", "t", false, "Help message for toggle")
}
