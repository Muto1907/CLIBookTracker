package cmd

import (
	"fmt"

	"github.com/spf13/cobra"
)

var reviewCmd = &cobra.Command{
	Use:   "review",
	Short: "Leaves a Review on the Book",
	Run: func(cmd *cobra.Command, args []string) {
		fmt.Println("What would you rate the Book?")
	},
}

func init() {
	rootCmd.AddCommand(reviewCmd)
}
