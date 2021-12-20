package cmd

import (
	"github.com/flrnd/gobatmon/db"
	"github.com/spf13/cobra"

	_ "github.com/mattn/go-sqlite3"
)

var cmdSave = &cobra.Command{
	Use:   "save",
	Short: "Save last period",
	Long:  "",
	Run: func(cmd *cobra.Command, args []string) {
		db.SaveLastPeriod()
	},
}
