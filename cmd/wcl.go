package cmd

import (
	"bufio"
	"fmt"
	"log"
	"os"

	"github.com/spf13/cobra"
)

// wclCmd represents the wcl command
var wclCmd = &cobra.Command{
	Use:   "wcl",
	Short: "wcl – line count",
	Long:  `The wcl utility displays the number of lines, contained in each input file.`,
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
		lineCount := 0
		for scanner.Scan() {
			lineCount++
		}
		fmt.Println(lineCount, args[0])
	},
}

func init() {
	rootCmd.AddCommand(wclCmd)
}
