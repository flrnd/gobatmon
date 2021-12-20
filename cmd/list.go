package cmd

import (
	"github.com/flrnd/gobatmon/db"
	"github.com/spf13/cobra"
)

func init() {
	rootCmd.AddCommand(listCmd)
	listCmd.AddCommand(savedCmd)
}

var listCmd = &cobra.Command{
	Use:   "list",
	Short: "list all timestamps. [help list] for more information.",
	Long:  "Output all stored timestamps. list saved shows stored discharging periods.",
	Run: func(cmd *cobra.Command, args []string) {
		db.List()
	},
}
