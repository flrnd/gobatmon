package cmd

import (
	"fmt"
	"time"

	"github.com/flrnd/gobatmon/db"
	"github.com/flrnd/gobatmon/util"
	"github.com/spf13/cobra"
)

type BatteryPeriod struct {
	timestamp      time.Time
	discharge      int
	dischargeTime  time.Duration
	dischargeRatio float32
}

var Period = BatteryPeriod{}

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

		Period.timestamp = timestamp
		Period.discharge = util.CalculateDischarge(currentCharge, charge)
		Period.dischargeTime = currentTime.Sub(timestamp)
		Period.dischargeRatio = util.CalculateDischargeRatePerHour(Period.discharge, Period.dischargeTime)

		fmt.Printf("Discharge      : %d%%\n", Period.discharge)
		fmt.Printf("Time elapsed   : %v\n", Period.dischargeTime)
		fmt.Printf("Discharge ratio: %0.3fWh\n", Period.dischargeRatio)
	},
}
