package util

import (
	"strconv"
	"strings"
)

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
