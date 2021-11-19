package util

import (
	"fmt"
	"io/ioutil"
	"log"
	"os"
	"strconv"
	"strings"
)

type BatteryStats struct {
	Manufacturer       string
	FullCapacityDesign int
	FullCapacity       int
	CurrentCharge      string
	Cycles             string
	DischargeRate      float32
	Status             string
}

func ParameterPath(p string) string {
	var batteryPath = "/sys/class/power_supply/BAT0/"
	return fmt.Sprintf("%s%s", batteryPath, p)
}

func Stats() BatteryStats {
	// check if the battery is present
	if _, err := os.Stat(ParameterPath("present")); os.IsNotExist(err) {
		log.Fatal("No battery on this system")
		os.Exit(1)
	}

	// Read design full charge
	fullCapacityDesign, err := ioutil.ReadFile(ParameterPath("energy_full_design"))
	Check(err)

	// Read full charge
	fullCapacity, err := ioutil.ReadFile(ParameterPath("energy_full"))
	Check(err)

	/*
		 * Read the currentCharge battery value
		currentCharge, err := ioutil.ReadFile(getParameterPath("energy_now"))
		util.Check(err)
		currentChargePercentage := util.CalculateBatteryPercentage(util.ParseBatteryValue(currentCharge), util.ParseBatteryValue(fullCapacity))
	*/

	// read capacity
	currentCapacity, err := ioutil.ReadFile(ParameterPath("capacity"))
	Check(err)

	// read dischargePower
	dischargePower, err := ioutil.ReadFile(ParameterPath("power_now"))
	Check(err)
	dischargeRate := ParseMilliWats(ParseBatteryValue(dischargePower))

	// read cycles
	cycles, err := ioutil.ReadFile(ParameterPath("cycle_count"))
	Check(err)

	// read manufacurer
	manufacturer, err := ioutil.ReadFile(ParameterPath("manufacturer"))
	Check(err)

	// read status
	status, err := ioutil.ReadFile(ParameterPath("status"))
	Check(err)

	return BatteryStats{
		Manufacturer:       TrimValue(manufacturer),
		Status:             TrimValue(status),
		FullCapacityDesign: ParseBatteryValue(fullCapacityDesign),
		FullCapacity:       ParseBatteryValue(fullCapacity),
		CurrentCharge:      string(TrimValue(currentCapacity)),
		Cycles:             string(cycles),
		DischargeRate:      dischargeRate,
	}
}

func CalculateBatteryPercentage(current, full int) int {
	return current * 100 / full
}

func ParseMilliWats(m int) float32 {
	return float32(m) / 1000
}

func TrimValue(d []byte) string {
	return strings.Trim(string(d), "\n")
}

func ParseBatteryValue(d []byte) int {
	v, err := strconv.Atoi(TrimValue(d))
	Check(err)

	return v / 1000
}
