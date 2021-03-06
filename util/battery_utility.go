package util

import (
	"fmt"
	"io/ioutil"
	"log"
	"os"
	"strconv"
	"strings"
	"time"
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

type BatteryPeriod struct {
	Timestamp        time.Time
	CurrentTimestamp time.Time
	Discharge        int
	DischargeTime    time.Duration
	DischargeRatio   float32
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

func DischargeStats(timestamp time.Time, currentTimestamp time.Time, charge int) (d int, dt time.Duration, dr float32) {
	currentCharge := Stats().Capacity
	d = CalculateDischarge(currentCharge, charge)
	dt = currentTimestamp.Sub(timestamp)
	dr = CalculateDischargeRatePerHour(d, dt)

	return d, dt, dr
}

func NewPeriod(timestamp time.Time, currentTimestamp time.Time, charge int) BatteryPeriod {
	discharge, dischargeTime, dischargeRatio := DischargeStats(timestamp, currentTimestamp, charge)

	return BatteryPeriod{
		Timestamp:        timestamp,
		CurrentTimestamp: currentTimestamp,
		Discharge:        discharge,
		DischargeTime:    dischargeTime,
		DischargeRatio:   dischargeRatio,
	}
}

func (p BatteryPeriod) Print() {
	fmt.Printf("Discharge      : %d%%\n", p.Discharge)
	fmt.Printf("Time elapsed   : %v\n", p.DischargeTime)
	fmt.Printf("Ratio: %0.3fWh\n", p.DischargeRatio)

}

func CalculateDischarge(current, old int) int {
	if (old - current) < 0 {
		fmt.Printf("It seems you charged your battery since last recorded time.\n")
		fmt.Printf("To start a new period run: gobatmon create.\n")
		os.Exit(1)
	}

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

/*
* Since time Duration is in nanoseconds, we need to work in nanoseconds
* to calculate the discharge ratio.
* d is discharge percentage, h time duration in nanoseconds
 */
func CalculateDischargeRatePerHour(d int, h time.Duration) float32 {
	return float32(d) * 3600000000000 / float32(h)
}
