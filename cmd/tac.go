package cmd

import (
	"bufio"
	"fmt"
	"log"
	"os"

	"github.com/spf13/cobra"
)

// tacCmd represents the tac command
var tacCmd = &cobra.Command{
	Use:   "tac",
	Short: "tac - concatenate and print files in reverse",
	Long:  `Write each FILE to standard output, last line first.`,
	Run: func(cmd *cobra.Command, args []string) {
		if len(args) != 1 {
			log.Fatal("The argument must be one.")
			return
		}
		file, err := os.Open(args[0])
		if err != nil {
			log.Fatal(err)
		}
		defer file.Close()
		scanner := bufio.NewScanner(file)
		var lines []string
		for scanner.Scan() {
			lines = append(lines, scanner.Text())
		}
		for i := len(lines) - 1; i >= 0; i-- {
			fmt.Println(lines[i])
		}
	},
}

func init() {
	rootCmd.AddCommand(tacCmd)
}
