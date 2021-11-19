package cmd

import "github.com/spf13/cobra"

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

	},
}
