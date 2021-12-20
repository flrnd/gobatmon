package cmd

import (
	"time"

	"github.com/flrnd/gobatmon/db"
	"github.com/flrnd/gobatmon/util"
	"github.com/spf13/cobra"
)

var Period = util.BatteryPeriod{}

func init() {
	rootCmd.AddCommand(lastCmd)
	lastCmd.AddCommand(cmdSave)
}

var lastCmd = &cobra.Command{
	Use:   "last",
	Short: "print discharge % since the last recorded timestamp",
	Long:  "",
	Run: func(cmd *cobra.Command, args []string) {
		currentCharge := util.Stats().Capacity
		currentTime := time.Now()
		_, charge, timestamp := db.Last()

		Period.Timestamp = timestamp
		Period.Discharge = util.CalculateDischarge(currentCharge, charge)
		Period.DischargeTime = currentTime.Sub(timestamp)
		Period.DischargeRatio = util.CalculateDischargeRatePerHour(Period.Discharge, Period.DischargeTime)

		Period.Print()
	},
}
