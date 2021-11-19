package cmd

import (
	"github.com/flrnd/gobatmon/db"
	"github.com/spf13/cobra"
)

func init() {
	rootCmd.AddCommand(listCmd)
}

var listCmd = &cobra.Command{
	Use:   "list",
	Short: "list all timestamps",
	Long:  ``,
	Run: func(cmd *cobra.Command, args []string) {
		db.List()
	},
}
