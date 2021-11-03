package main

import (
	"database/sql"
	"fmt"
	"io/ioutil"
	"log"
	"os"
	"strconv"
	"strings"

	_ "github.com/mattn/go-sqlite3"
)

// const energyFullDesign = "/sys/class/power_supply/BAT0/energy_full_design"
const energyFull = "/sys/class/power_supply/BAT0/energy_full"
const energyNow = "/sys/class/power_supply/BAT0/energy_now"

// const powerNow = "/sys/class/power_supply/BAT0/power_now"
// const batteryStatus = "/sys/class/power_supply/BAT0/status"

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
	homeDir, err := os.UserHomeDir()
	check(err)

	databasePath := homeDir + "/.config/batmon/battery.db"

	// Check if database exists
	if _, err := os.Stat(databasePath); os.IsNotExist(err) {
		os.MkdirAll(homeDir+"/.config/batmon", 0755)
		log.Println("Creating the database")
		file, err := os.Create(databasePath)
		check(err)
		file.Close()
		log.Println("Database created")
	}

	// Open the database
	batteryDB, err := sql.Open("sqlite3", databasePath)
	check(err)
	defer batteryDB.Close()
	createTable(batteryDB)

	// Read the current battery value
	current, err := ioutil.ReadFile(energyNow)
	check(err)
	energyFull, err := ioutil.ReadFile(energyFull)
	check(err)
	percentage := calculateBatteryPercentage(parseBatteryValue(current), parseBatteryValue(energyFull))
	fmt.Printf("%d%%\n", percentage)

}

func createTable(batteryDB *sql.DB) {
	createBatteryTable := `CREATE TABLE IF NOT EXISTS battery (
		"id" INTEGER PRIMARY KEY AUTOINCREMENT,
		"percentage" INTEGER NOT NULL,
		"timestamp" INTEGER NOT NULL,
		"FOREIGN KEY(percentage) REFERENCES percentage(id),
		"FOREIGN KEY(timestamp) REFERENCES timestamp(id),
		"UNIQUE(percentage, timestamp)"
	);`

	statement, err := batteryDB.Prepare(createBatteryTable)
	check(err)
	statement.Exec()
}
