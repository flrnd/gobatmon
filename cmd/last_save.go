package cmd

import (
	"github.com/flrnd/gobatmon/db"
	_ "github.com/mattn/go-sqlite3"
	"github.com/spf13/cobra"
)

var cmdSave = &cobra.Command{
	Use:   "save",
	Short: "Save last period",
	Long:  "",
	Run: func(cmd *cobra.Command, args []string) {
		db.SaveLastPeriod()
	},
}
