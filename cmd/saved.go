package cmd

import (
	"github.com/flrnd/gobatmon/db"
	"github.com/muesli/coral"
)

var savedCmd = &coral.Command{
	Use:   "saved",
	Short: "list all saved periods",
	Long:  ``,
	Run: func(cmd *coral.Command, args []string) {
		db.ListSavedPeriods()
	},
}
