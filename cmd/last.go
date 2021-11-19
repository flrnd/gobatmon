package cmd

import (
	"fmt"
	"time"

	"github.com/flrnd/gobatmon/db"
	"github.com/flrnd/gobatmon/util"
	"github.com/spf13/cobra"
)

func init() {
	rootCmd.AddCommand(lastCmd)
}

var lastCmd = &cobra.Command{
	Use:   "last",
	Short: "print discharge % since the last recorded timestamp",
	Long:  "",
	Run: func(cmd *cobra.Command, args []string) {
		currentCharge := util.Stats().CurrentCharge
		currentTime := time.Now()
		_, charge, timestamp := db.Last()

		fmt.Printf("%d%% discharge in %v\n", util.CalculateDischarge(currentCharge, charge), currentTime.Sub(timestamp))
	},
}
