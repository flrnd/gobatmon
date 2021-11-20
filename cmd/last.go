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
		currentCharge := util.Stats().Capacity
		currentTime := time.Now()
		_, charge, timestamp := db.Last()

		discharge := util.CalculateDischarge(currentCharge, charge)
		dischargeTime := currentTime.Sub(timestamp)
		fmt.Printf("Discharge      : %d%%\n", discharge)
		fmt.Printf("Time elapsed   : %v\n", dischargeTime)
		fmt.Printf("Discharge ratio: %0.3fWh\n", util.CalculateDischargeRatePerHour(discharge, dischargeTime))
	},
}
