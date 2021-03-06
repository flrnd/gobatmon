package cmd

import (
	"time"

	"github.com/flrnd/gobatmon/db"
	"github.com/flrnd/gobatmon/util"
	"github.com/spf13/cobra"
)

func init() {
	rootCmd.AddCommand(lastCmd)
	lastCmd.AddCommand(cmdSave)
}

var lastCmd = &cobra.Command{
	Use:   "last",
	Short: "print discharge % since the last recorded timestamp. [help last] for more information",
	Long:  "Print the discharge percentage since last recorded timestamp. Accepts save flag to create an entry on the database.",
	Run: func(cmd *cobra.Command, args []string) {
		_, charge, timestamp := db.Last()

		period := util.NewPeriod(timestamp, time.Now(), charge)
		period.Print()
	},
}
