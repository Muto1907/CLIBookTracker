package cmd

import (
	"fmt"

	"github.com/spf13/cobra"
)

var rootCmd = &cobra.Command{
	Use:   "boocktracker",
	Short: "BookTracker is a CLI tool to manage your readings",
	Long:  "BookTracker helps you store, track, and manage your reading progress and make notes on them.",
	Run: func(cmd *cobra.Command, args []string) {
		fmt.Printf("Welcome to BookTrackerCLI")
	},
}

func Execute() error {
	return rootCmd.Execute()
}
