package cmd

import (
	"time"

	"github.com/flrnd/gobatmon/db"
	"github.com/flrnd/gobatmon/util"
	_ "github.com/mattn/go-sqlite3"
	"github.com/spf13/cobra"
)

func init() {
	rootCmd.AddCommand(createCmd)
}

var createCmd = &cobra.Command{
	Use:   "create",
	Short: "create a timestamp",
	Long:  ``,
	Run: func(cmd *cobra.Command, args []string) {
		currentCharge := util.Stats().Capacity
		timestamp := time.Now()
		db.Insert(currentCharge, timestamp)
	},
}
