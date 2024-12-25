package cmd

import (
	"fmt"

	"github.com/spf13/cobra"
)

var noteCmd = &cobra.Command{
	Use:   "note",
	Short: "Adds a Node to the current Chapter",
	Run: func(cmd *cobra.Command, args []string) {
		fmt.Println("What are you notes for today?")
	},
}

func init() {
	rootCmd.AddCommand(noteCmd)
}
