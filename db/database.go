package db

import (
	"database/sql"
	"fmt"
	"os"
	"time"

	"github.com/flrnd/gobatmon/util"
)

const defaultDbPath = "/.config/batmon/battery.db"

func Path() string {
	homeDir, _ := os.UserHomeDir()

	return homeDir + defaultDbPath
}

func Init() {
	homeDir, _ := os.UserHomeDir()
	databasePath := Path()

	if _, err := os.Stat(databasePath); os.IsNotExist(err) {
		os.MkdirAll(homeDir+"/.config/batmon", 0755)
		// log.Println("Creating the database")
		file, err := os.Create(databasePath)
		util.Check(err)

		file.Close()
		// log.Println("Database created")
	}

	// Open the database
	batteryDB, err := sql.Open("sqlite3", databasePath)
	util.Check(err)
	defer batteryDB.Close()
	createTable(batteryDB)

}

func Insert(charge string, timestamp time.Time) {
	db, err := sql.Open("sqlite3", Path())
	util.Check(err)
	defer db.Close()

	stmt, err := db.Prepare("INSERT INTO battery_charge(charge, timestamp) values(?,?)")
	util.Check(err)

	res, err := stmt.Exec(charge, timestamp)
	util.Check(err)

	id, err := res.LastInsertId()

	fmt.Printf("created timestamp: %d\n %s - %s\n", id, timestamp.UTC(), charge)
}

func createTable(batteryDB *sql.DB) {
	createBatteryTable := `CREATE TABLE IF NOT EXISTS "battery_charge" (
		"id"	INTEGER NOT NULL UNIQUE,
		"charge"	string NOT NULL,
		"timestamp"	INTEGER NOT NULL,
		PRIMARY KEY("id" AUTOINCREMENT));`

	if statement, err := batteryDB.Prepare(createBatteryTable); err != nil {
		util.Check(err)
	} else {
		statement.Exec()
	}
}
