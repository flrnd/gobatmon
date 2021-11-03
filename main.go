package main

import (
	"fmt"
	"io/ioutil"
	"strconv"
	"strings"
)

const energyFullDesign = "/sys/class/power_supply/BAT0/energy_full_design"
const energyFull = "/sys/class/power_supply/BAT0/energy_full"
const energyNow = "/sys/class/power_supply/BAT0/energy_now"
const powerNow = "/sys/class/power_supply/BAT0/power_now"
const batteryStatus = "/sys/class/power_supply/BAT0/status"

func check(e error) {
	if e != nil {
		panic(e)
	}
}

func parseBatteryValue(d []byte) int {
	v, err := strconv.Atoi(strings.Trim(string(d), "\n"))
	check(err)

	return v / 1000
}

func calculateBatteryPercentage(current, full int) int {
	return current * 100 / full
}

func main() {
	current, err := ioutil.ReadFile(energyNow)
	check(err)
	energyFull, err := ioutil.ReadFile(energyFull)
	check(err)
	percentage := calculateBatteryPercentage(parseBatteryValue(current), parseBatteryValue(energyFull))
	fmt.Printf("%d%%\n", percentage)
}

