package util

import (
	"fmt"
	"io/ioutil"
	"log"
	"os"
	"strconv"
	"strings"
)

/*
 * From https://www.kernel.org/doc/html/latest/power/power_supply_class.html
 *
 * Quoting include/linux/power_supply.h:
 * All voltages, currents, charges, energies, time and temperatures in µV, µA, µAh, µWh, seconds and tenths of degree Celsius
 * unless otherwise stated. It’s driver’s job to convert its raw values to units in which this class operates.
 *
 * CHARGE_* : attributes represents capacity in µAh only.
 * ENERGY_* : attributes represents capacity in µWh only.
 * CAPACITY : attribute represents capacity in percents, from 0 to 100.
 */

type BatteryStats struct {
	Manufacturer     string
	EnergyFullDesign float32
	EnergyFull       float32
	EnergyNow        float32
	Capacity         int
	Cycles           int
	PowerNow         float32
	Status           string
}

const BATTERY_PATH = "/sys/class/power_supply/BAT0/"

func ParameterPath(p string) string {
	return fmt.Sprintf("%s%s", BATTERY_PATH, p)
}

func EnergyFullDesign() float32 {
	value, err := ioutil.ReadFile(ParameterPath("energy_full_design"))
	Check(err)

	// return value in Wh
	return MilliWattsToWatts(ByteStringToInt(value) / 1000)
}

func EnergyFull() float32 {
	value, err := ioutil.ReadFile(ParameterPath("energy_full"))
	Check(err)

	// return value in Wh
	return MilliWattsToWatts(ByteStringToInt(value) / 1000)
}

func EnergyNow() float32 {
	value, err := ioutil.ReadFile(ParameterPath("energy_now"))
	Check(err)

	return MilliWattsToWatts(ByteStringToInt(value) / 1000)
}

func Manufacturer() string {
	// read manufacurer
	value, err := ioutil.ReadFile(ParameterPath("manufacturer"))
	Check(err)

	return ByteStringToString(value)
}

func Capacity() int {
	return InputByteArrayToInt("capacity")
}

func PowerNow() float32 {
	pn, err := ioutil.ReadFile(ParameterPath("power_now"))
	Check(err)

	//return value in W
	return MilliWattsToWatts(ByteStringToInt(pn) / 1000)
}

func CycleCount() int {
	return InputByteArrayToInt("cycle_count")
}

func Status() string {
	status, err := ioutil.ReadFile(ParameterPath("status"))
	Check(err)

	return ByteStringToString(status)
}

func CheckBattery() {
	// check if the battery is present
	if _, err := os.Stat(ParameterPath("present")); os.IsNotExist(err) {
		log.Fatal("No battery on this system")
		os.Exit(1)
	}
}

func Stats() BatteryStats {
	CheckBattery()

	return BatteryStats{
		Manufacturer:     Manufacturer(),
		Status:           Status(),
		EnergyFullDesign: EnergyFullDesign(),
		EnergyFull:       EnergyFull(),
		EnergyNow:        EnergyNow(),
		Capacity:         Capacity(),
		Cycles:           CycleCount(),
		PowerNow:         PowerNow(),
	}
}

func CalculateDischarge(current, old int) int {
	return old - current
}

func MilliWattsToWatts(m int) float32 {
	return float32(m) / 1000
}

func ByteStringToString(d []byte) string {
	return strings.Trim(string(d), "\n")
}

func ByteStringToInt(d []byte) int {
	v, err := strconv.Atoi(ByteStringToString(d))
	Check(err)

	return v
}

func InputByteArrayToInt(s string) int {
	is, err := ioutil.ReadFile(ParameterPath(s))
	Check(err)
	v, err := strconv.Atoi(ByteStringToString(is))
	Check(err)

	return v
}
