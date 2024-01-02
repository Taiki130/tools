/*
Copyright © 2024 NAME HERE <EMAIL ADDRESS>
*/
package cmd

import (
	"fmt"
	"log"

	"github.com/robfig/cron/v3"
	"github.com/spf13/cobra"
)

// cronCmd represents the cron command
var cronCmd = &cobra.Command{
	Use:   "cron",
	Short: "毎日午前3時にバックアップスクリプトを実行するcronジョブを作成する。",
	Long:  ``,
	Run: func(cmd *cobra.Command, args []string) {
		c := cron.New()
		_, err := c.AddFunc("0 0 3 * * *", backupScript)
		if err != nil {
			log.Fatal(err)
		}
		c.Start()
		defer c.Stop()
		select {}
	},
}

func backupScript() {
	fmt.Println("バックアップを実行します。")
}

func init() {
	rootCmd.AddCommand(cronCmd)
}
