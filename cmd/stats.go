package cmd

import (
	"fmt"
	"log"
	"os"

	"github.com/flrnd/gobatmon/util"
	"github.com/spf13/cobra"
)

func init() {
	rootCmd.AddCommand(statsCmd)
}

var statsCmd = &cobra.Command{
	Use:   "stats",
	Short: "print battery stats",
	Long: `Print battery stats like full charge
								design capacity, current charge, current discharge rate...`,
	Run: func(cmd *cobra.Command, args []string) {
		// check if the battery is present
		if _, err := os.Stat(util.ParameterPath("present")); os.IsNotExist(err) {
			log.Fatal("No battery on this system")
			os.Exit(1)
		}

		stats := util.Stats()

		// print the battery stats
		fmt.Println()
		fmt.Printf("Manufacturer: %s\n", stats.Manufacturer)
		fmt.Printf("Status: %s\n", stats.Status)
		fmt.Printf("Energy full design : %0.2f Wh\n", stats.EnergyFullDesign)
		fmt.Printf("Energy full        : %0.2f Wh\n", stats.EnergyFull)
		fmt.Printf("Energy now         : %0.2f Wh\n", stats.EnergyNow)
		fmt.Printf("Energy rate        : %.2f W\n", stats.PowerNow)
		fmt.Printf("Charge             : %d%%\n", stats.Capacity)
		fmt.Printf("Cycle count: %d\n", stats.Cycles)
		fmt.Println()
	},
}
