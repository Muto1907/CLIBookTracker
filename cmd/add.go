package cmd

import (
	"fmt"

	"github.com/spf13/cobra"
)

var addCmd = &cobra.Command{
	Use:   "add",
	Short: "Adds a Book to your List",
	Run: func(cmd *cobra.Command, args []string) {
		fmt.Println("Adding the following Book to your list:")
	},
}

func init() {
	rootCmd.AddCommand(addCmd)
}
