package cmd

import (
	"fmt"
	"io/ioutil"
	"os"

	"github.com/spf13/cobra"
)

const energyFullDesignPath = "/sys/class/power_supply/BAT0/energy_full_design"
const energyFullPath = "/sys/class/power_supply/BAT0/energy_full"
const energyNowPath = "/sys/class/power_supply/BAT0/energy_now"
const powerNowPath = "/sys/class/power_supply/BAT0/power_now"

var statsCmd = &cobra.Command{
	Use:   "stats",
	Short: "print battery stats",
	Long: `Print battery stats like full charge
								full power, full fuck you...`,
	Run: func(cmd *cobra.Command, args []string) {
		// Read the current battery value
		current, err := ioutil.ReadFile(energyNowPath)
		Check(err)
		energyFull, err := ioutil.ReadFile(energyFullPath)
		check(err)
		percentage := calculateBatteryPercentage(parseBatteryValue(current), parseBatteryValue(energyFull))
		fmt.Printf("%d%%\n", percentage)
	},
}

func Execute() {
	if err := statsCmd.Execute(); err != nil {
		fmt.Fprintln(os.Stderr, err)
		os.Exit(1)
	}
}
