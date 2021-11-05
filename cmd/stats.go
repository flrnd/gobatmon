package cmd

import (
	"fmt"
	"io/ioutil"

	"github.com/flrnd/gobatmon/util"
	"github.com/spf13/cobra"
)

func init() {
	rootCmd.AddCommand(statsCmd)
}

const cycleCountPath = "/sys/class/power_supply/BAT0/cycle_count"
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
		// Read design full charge
		fullDesign, err := ioutil.ReadFile(energyFullDesignPath)
		util.Check(err)
		// Read full charge
		energyFull, err := ioutil.ReadFile(energyFullPath)
		util.Check(err)
		// Read the current battery value
		current, err := ioutil.ReadFile(energyNowPath)
		util.Check(err)
		percentage := util.CalculateBatteryPercentage(util.ParseBatteryValue(current), util.ParseBatteryValue(energyFull))
		// read power
		power, err := ioutil.ReadFile(powerNowPath)
		util.Check(err)

		// read cycles
		cycles, err := ioutil.ReadFile(cycleCountPath)
		util.Check(err)

		// print the battery stats
		fmt.Printf("Full design capacity: %dmWh\n", util.ParseBatteryValue(fullDesign))
		fmt.Printf("Full charge capacity: %dmWh\n", util.ParseBatteryValue(energyFull))
		fmt.Printf("Current charge at: %d%%\n", percentage)
		fmt.Printf("Cycle count: %s", string(cycles))
		fmt.Printf("Current power discharge at: %.2fW\n", util.ParseMilliWats(util.ParseBatteryValue(power)))
	},
}
