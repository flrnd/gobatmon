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

func Insert(charge int, timestamp time.Time) {
	db, err := sql.Open("sqlite3", Path())
	util.Check(err)
	defer db.Close()

	stmt, err := db.Prepare("INSERT INTO battery_charge(charge, timestamp) values(?,?)")
	util.Check(err)

	res, err := stmt.Exec(charge, timestamp)
	util.Check(err)

	id, err := res.LastInsertId()

	fmt.Printf("created timestamp (%d) %d%% at %s\n", id, charge, util.ParseTime(timestamp))
}

func Last() (id, charge int, timestamp time.Time) {
	db, err := sql.Open("sqlite3", Path())
	util.Check(err)
	defer db.Close()
	sqlStatement := "SELECT * FROM battery_charge WHERE id = (SELECT MAX(id) FROM battery_charge)"
	rows, err := db.Query(sqlStatement)
	util.Check(err)

	if rows.Next() {
		err = rows.Scan(&id, &charge, &timestamp)
		util.Check(err)
	}
	rows.Close()
	return id, charge, timestamp
}

func List() {
	db, err := sql.Open("sqlite3", Path())
	util.Check(err)
	defer db.Close()

	rows, err := db.Query("SELECT * FROM battery_charge")
	util.Check(err)

	for rows.Next() {
		var id int
		var charge string
		var timestamp time.Time
		err = rows.Scan(&id, &charge, &timestamp)
		util.Check(err)
		fmt.Printf("id: %d charge: %s%% created: %s\n", id, charge, util.ParseTime(timestamp))
	}
	rows.Close()
}

func createTable(batteryDB *sql.DB) {
	createBatteryTable := `CREATE TABLE IF NOT EXISTS "battery_charge" (
		"id"	INTEGER NOT NULL UNIQUE,
		"charge"	INTEGER NOT NULL,
		"timestamp"	datetime NOT NULL,
		PRIMARY KEY("id" AUTOINCREMENT));`

	if statement, err := batteryDB.Prepare(createBatteryTable); err != nil {
		util.Check(err)
	} else {
		statement.Exec()
	}
}
