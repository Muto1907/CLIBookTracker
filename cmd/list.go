package cmd

import (
	"fmt"

	"github.com/spf13/cobra"
)

var listCmd = &cobra.Command{
	Use:   "list",
	Short: "List all books in your reading list",
	Run: func(cmd *cobra.Command, args []string) {
		fmt.Println("Here are all your books:")
	},
}

func init() {
	rootCmd.AddCommand(listCmd)
}
