package cmd

import "github.com/spf13/cobra"

type BatteryStamp struct {
	charge    int
	timestamp int
}

func init() {
	rootCmd.AddCommand(batteryCmd)
}

var batteryCmd = &cobra.Command{
	Use:   "battery",
	Short: "battery commands",
	Long:  ``,
	Run:   func(cmd *cobra.Command, args []string) {},
}
