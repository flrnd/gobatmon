package cmd

import (
	"github.com/flrnd/gobatmon/db"
	"github.com/spf13/cobra"
)

var savedCmd = &cobra.Command{
	Use:   "saved",
	Short: "list all saved periods",
	Long:  ``,
	Run: func(cmd *cobra.Command, args []string) {
		db.ListSavedPeriods()
	},
}
