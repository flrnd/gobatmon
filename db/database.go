package db

import (
	"database/sql"
	"log"
	"os"

	"github.com/flrnd/gobatmon/util"
)

const defaultDbPath = "/.config/batmon/battery.db"

func Path() string {
	homeDir, _ := os.UserHomeDir()

	return homeDir + defaultDbPath
}

func Init() error {
	homeDir, _ := os.UserHomeDir()
	databasePath := Path()

	if _, err := os.Stat(databasePath); os.IsNotExist(err) {
		os.MkdirAll(homeDir+"/.config/batmon", 0755)
		log.Println("Creating the database")
		if file, err := os.Create(databasePath); err != nil {
			return err
		} else {
			file.Close()
			log.Println("Database created")
		}
	}

	// Open the database
	if batteryDB, err := sql.Open("sqlite3", databasePath); err != nil {
		return err
	} else {
		batteryDB.Close()
		defer batteryDB.Close()
		createTable(batteryDB)
	}
	return nil
}

func createTable(batteryDB *sql.DB) {
	createBatteryTable := `CREATE TABLE IF NOT EXISTS "battery_charge" (
		"id"	INTEGER NOT NULL UNIQUE,
		"charge"	INTEGER NOT NULL,
		"timestamp"	INTEGER NOT NULL,
		PRIMARY KEY("id" AUTOINCREMENT));`

	if statement, err := batteryDB.Prepare(createBatteryTable); err != nil {
		util.Check(err)
	} else {
		statement.Exec()
	}
}
