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

type AvailableBook struct {
	BookTitle  string
	AutherName string
}

type LoansBook struct {
	BookTitle  string
	AutherName string
}

type ExpiredBook struct {
	UserName  string
	BookTitle string
}

func DBExec(db *sql.DB, file string) error {
	query, err := os.ReadFile(file)
	if err != nil {
		return err
	}
	queryString := string(query)
	if _, err = db.Exec(queryString); err != nil {
		return err
	}
	return nil
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
		if err != nil {
			log.Fatal(err)
		}

		dbDriver := "postgres"
		dsn := fmt.Sprintf("host=%s port=5432 user=%s password=%s dbname=%s sslmode=disable", cfg.Host, cfg.User, cfg.Password, cfg.DBName)
		db, err := sql.Open(dbDriver, dsn)
		if err != nil {
			log.Fatal(err)
		}
		defer db.Close()

		if err = db.Ping(); err != nil {
			log.Fatal(err)
		}

		fmt.Printf("Connection to PostgreSQL server at %s is successful.\n", cfg.Host)

		files := []string{"db/ddl.sql", "db/dml.sql"}

		for _, file := range files {
			if err := DBExec(db, file); err != nil {
				log.Fatal(err)
			}
		}

		queryString := []string{
			"SELECT books.title, authors.name FROM books JOIN authors ON books.author_id = authors.id;",
			"SELECT books.title, authors.name FROM books JOIN authors ON books.author_id = authors.id JOIN loans ON books.id = loans.book_id WHERE loans.user_id = 1;",
			"SELECT users.name, books.title FROM users JOIN loans ON users.id = loans.user_id JOIN books ON loans.book_id = books.id WHERE loans.due_date < CURRENT_DATE;",
		}

		rows, err := db.Query(queryString[0])
		if err != nil {
			log.Fatal(err)
		}
		defer rows.Close()
		var ab AvailableBook
		fmt.Println("全ての利用可能な書籍のタイトルと著者名は、")
		for rows.Next() {
			err := rows.Scan(&ab.BookTitle, &ab.AutherName)
			if err != nil {
				log.Fatal(err)
			}
			fmt.Printf("BookTitle: %s, AutherName: %s\n", ab.BookTitle, ab.AutherName)
		}

		rows, err = db.Query(queryString[1])
		if err != nil {
			log.Fatal(err)
		}
		defer rows.Close()
		var lb LoansBook
		fmt.Println("指定されたユーザーが借りた書籍は、")
		for rows.Next() {
			err := rows.Scan(&lb.BookTitle, &lb.AutherName)
			if err != nil {
				log.Fatal(err)
			}
			fmt.Printf("BookTitle: %s, AutherName: %s\n", lb.BookTitle, lb.AutherName)
		}

		rows, err = db.Query(queryString[2])
		if err != nil {
			log.Fatal(err)
		}
		defer rows.Close()
		var eb ExpiredBook
		fmt.Println("期限切れの貸出書籍のユーザー名とタイトルは、")
		for rows.Next() {
			err := rows.Scan(&eb.UserName, &eb.BookTitle)
			if err != nil {
				log.Fatal(err)
			}
			fmt.Printf("UserName: %s, BookTitle: %s\n", eb.UserName, eb.BookTitle)
		}
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
