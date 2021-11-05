package cmd

import (
	"fmt"
	"io/ioutil"
	"log"
	"os"

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
		// check if the battery is present
		if _, err := os.Stat(getParameterPath("present")); os.IsNotExist(err) {
			log.Fatal("No battery on this system")
			os.Exit(1)
		}
		// Read design full charge
		fullCapacityDesign, err := ioutil.ReadFile(getParameterPath("energy_full_design"))
		util.Check(err)

		// Read full charge
		fullCapacity, err := ioutil.ReadFile(getParameterPath("energy_full"))
		util.Check(err)

		/*
			 * Read the currentCharge battery value
			currentCharge, err := ioutil.ReadFile(getParameterPath("energy_now"))
			util.Check(err)
			currentChargePercentage := util.CalculateBatteryPercentage(util.ParseBatteryValue(currentCharge), util.ParseBatteryValue(fullCapacity))
		*/

		// read capacity
		currentCapacity, err := ioutil.ReadFile(getParameterPath("capacity"))
		util.Check(err)

		// read dischargePower
		dischargePower, err := ioutil.ReadFile(getParameterPath("power_now"))
		util.Check(err)
		dischargeRate := util.ParseMilliWats(util.ParseBatteryValue(dischargePower))

		// read cycles
		cycles, err := ioutil.ReadFile(getParameterPath("cycle_count"))
		util.Check(err)

		// read manufacurer
		manufacturer, err := ioutil.ReadFile(getParameterPath("manufacturer"))
		util.Check(err)

		// read status
		status, err := ioutil.ReadFile(getParameterPath("status"))
		util.Check(err)

		// print the battery stats
		fmt.Println()
		fmt.Printf("Manufacturer: %s", manufacturer)
		fmt.Printf("Status: %s", status)
		fmt.Printf("Full design capacity: %d mWh\n", util.ParseBatteryValue(fullCapacityDesign))
		fmt.Printf("Full charge capacity: %d mWh\n", util.ParseBatteryValue(fullCapacity))
		fmt.Printf("Current charge at: %s%% | Discharge rate of %.2f W\n", string(util.TrimValue(currentCapacity)), dischargeRate)
		fmt.Printf("Cycle count: %s", string(cycles))
		fmt.Println()
	},
}
