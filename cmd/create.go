package cmd

import (
	"time"

	"github.com/flrnd/gobatmon/db"
	"github.com/flrnd/gobatmon/util"
	"github.com/spf13/cobra"

	_ "github.com/mattn/go-sqlite3"
)

type BatteryStamp struct {
	charge    int
	timestamp int
}

func init() {
	rootCmd.AddCommand(createCmd)
}

var createCmd = &cobra.Command{
	Use:   "create",
	Short: "create a timestamp",
	Long:  ``,
	Run: func(cmd *cobra.Command, args []string) {
		currentCharge := util.Stats().CurrentCharge
		timestamp := time.Now()
		db.Insert(currentCharge, timestamp)
	},
}