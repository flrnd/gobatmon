package cmd

import (
	"fmt"
	"io/ioutil"

	"github.com/flrnd/gobatmon/util"
	"github.com/spf13/cobra"
)

func getParameterPath(p string) string {
	var batteryPath = "/sys/class/power_supply/BAT0/"
	return fmt.Sprintf("%s%s", batteryPath, p)
}

func init() {
	rootCmd.AddCommand(statsCmd)
}

var statsCmd = &cobra.Command{
	Use:   "stats",
	Short: "print battery stats",
	Long: `Print battery stats like full charge
								full power, full fuck you...`,
	Run: func(cmd *cobra.Command, args []string) {
		// Read design full charge
		fullCapacityDesign, err := ioutil.ReadFile(getParameterPath("energy_full_design"))
		util.Check(err)

		// Read full charge
		fullCapacity, err := ioutil.ReadFile(getParameterPath("energy_full"))
		util.Check(err)

		// Read the currentCharge battery value
		currentCharge, err := ioutil.ReadFile(getParameterPath("energy_now"))
		util.Check(err)
		currentChargePercentage := util.CalculateBatteryPercentage(util.ParseBatteryValue(currentCharge), util.ParseBatteryValue(fullCapacity))

		// read dischargePower
		dischargePower, err := ioutil.ReadFile(getParameterPath("power_now"))
		util.Check(err)

		// read cycles
		cycles, err := ioutil.ReadFile(getParameterPath("cycle_count"))
		util.Check(err)

		// print the battery stats
		fmt.Printf("Full design capacity: %dmWh\n", util.ParseBatteryValue(fullCapacityDesign))
		fmt.Printf("Full charge capacity: %dmWh\n", util.ParseBatteryValue(fullCapacity))
		fmt.Printf("Current charge at: %d%%\n", currentChargePercentage)
		fmt.Printf("Cycle count: %s", string(cycles))
		fmt.Printf("Current power discharge at: %.2fW\n", util.ParseMilliWats(util.ParseBatteryValue(dischargePower)))
	},
}
