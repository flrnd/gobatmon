package util

import (
	"strconv"
	"strings"
)

func CalculateBatteryPercentage(current, full int) int {
	return current * 100 / full
}

func ParseBatteryValue(d []byte) int {
	v, err := strconv.Atoi(strings.Trim(string(d), "\n"))
	Check(err)

	return v / 1000
}
