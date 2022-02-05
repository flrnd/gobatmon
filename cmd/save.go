package cmd

import (
	"github.com/flrnd/gobatmon/db"
	_ "github.com/mattn/go-sqlite3"
	"github.com/muesli/coral"
)

var cmdSave = &coral.Command{
	Use:   "save",
	Short: "Save last period",
	Long:  "",
	Run: func(cmd *coral.Command, args []string) {
		db.SaveLastPeriod()
	},
}
