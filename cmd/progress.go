package cmd

import (
	"fmt"

	"github.com/spf13/cobra"
)

var progCmd = &cobra.Command{
	Use:   "progress",
	Short: "Updates your reading progress",
	Run: func(cmd *cobra.Command, args []string) {
		fmt.Println("Updating reading progress")
	},
}

func init() {
	rootCmd.AddCommand(progCmd)
}
