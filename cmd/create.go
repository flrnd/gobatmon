package cmd

import (
	"time"

	"github.com/flrnd/gobatmon/db"
	"github.com/flrnd/gobatmon/util"
	_ "github.com/mattn/go-sqlite3"
	"github.com/muesli/coral"
)

func init() {
	rootCmd.AddCommand(createCmd)
}

var createCmd = &coral.Command{
	Use:   "create",
	Short: "create a timestamp",
	Long:  ``,
	Run: func(cmd *coral.Command, args []string) {
		currentCharge := util.Stats().Capacity
		timestamp := time.Now()
		db.Insert(currentCharge, timestamp)
	},
}
