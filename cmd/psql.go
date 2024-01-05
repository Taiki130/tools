/*
Copyright © 2024 NAME HERE <EMAIL ADDRESS>
*/
package cmd

import (
	"database/sql"
	"fmt"
	"log"
	"os"

	_ "github.com/lib/pq"
	"github.com/spf13/cobra"
	yaml "gopkg.in/yaml.v3"
)

type Config struct {
	Host     string `yaml:"host"`
	User     string `yaml:"user"`
	Password string `yaml:"password"`
	DBName   string `yaml:"dbname"`
}

// psqlCmd represents the psql command
var psqlCmd = &cobra.Command{
	Use:   "psql",
	Short: "指定したYAMLファイルからPostgreSQLサーバーの接続情報を取得し、接続が成功するかどうかを確認する",
	Long:  ``,
	Run: func(cmd *cobra.Command, args []string) {
		configFile, err := os.ReadFile("db/config.yml")
		if err != nil {
			log.Fatal(err)
		}

		var cfg Config
		err = yaml.Unmarshal(configFile, &cfg)

		dbDriver := "postgres"
		dsn := fmt.Sprintf("host=%s port=5432 user=%s password=%s dbname=%s sslmode=disable", cfg.Host, cfg.User, cfg.Password, cfg.DBName)
		db, err := sql.Open(dbDriver, dsn)
		if err != nil {
			log.Fatal(err)
		}
		defer db.Close()

		fmt.Printf("Connection to PostgreSQL server at %s is successful.\n", cfg.Host)
	},
}

func init() {
	rootCmd.AddCommand(psqlCmd)

	// Here you will define your flags and configuration settings.

	// Cobra supports Persistent Flags which will work for this command
	// and all subcommands, e.g.:
	// psqlCmd.PersistentFlags().String("foo", "", "A help for foo")

	// Cobra supports local flags which will only run when this command
	// is called directly, e.g.:
	// psqlCmd.Flags().BoolP("toggle", "t", false, "Help message for toggle")
}
