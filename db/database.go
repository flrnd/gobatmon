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
	createTables(batteryDB)

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

func Delete(columnId int, table string) {
	db, err := sql.Open("sqlite3", Path())
	util.Check(err)
	defer db.Close()

	stmt := fmt.Sprintf("DELETE FROM %s WHERE id = %d", table, columnId)
	util.Check(err)

	res, err := db.Exec(stmt)
	util.Check(err)

	id, _ := res.RowsAffected()

	fmt.Printf("Deleted %d timestamp (id=%d)\n", id, columnId)
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

func SaveLastPeriod() {
	db, err := sql.Open("sqlite3", Path())
	util.Check(err)
	defer db.Close()

	stmt, err := db.Prepare("INSERT INTO battery_last_period(timestamp_start, timestamp_end, discharge, discharge_ratio, last_charge) values(?,?,?,?,?)")
	util.Check(err)

	lastId, charge, timestamp := Last()
	currentTimestamp := time.Now()
	period := util.NewPeriod(timestamp, currentTimestamp, charge)

	res, err := stmt.Exec(period.Timestamp, period.CurrentTimestamp, period.Discharge, period.DischargeRatio, lastId)
	util.Check(err)

	id, err := res.LastInsertId()

	fmt.Printf("Saved (%d) %d%% elapsed: %v dr: %0.3fWh\n", id, period.Discharge, period.DischargeTime, period.DischargeRatio)
}

func ListSavedPeriods() {
	db, err := sql.Open("sqlite3", Path())
	util.Check(err)
	defer db.Close()

	rows, err := db.Query("SELECT * FROM battery_last_period")
	util.Check(err)

	for rows.Next() {
		var id int
		var timestamp time.Time
		var currentTimestamp time.Time
		var discharge int
		var dischargeRatio float32
		var lastId int
		err = rows.Scan(&id, &timestamp, &currentTimestamp, &discharge, &dischargeRatio, &lastId)
		util.Check(err)

		fmt.Printf("id: %d | lastId: %d\n", id, lastId)
		fmt.Printf("Discharge    : %d%%\n", discharge)
		fmt.Printf("From         : %s\n", util.ParseTime(timestamp))
		fmt.Printf("To           : %s\n", util.ParseTime(currentTimestamp))
		fmt.Printf("Time elapsed : %v\n", currentTimestamp.Sub(timestamp).String())
		fmt.Printf("Ratio        : %0.3fWh\n\n", dischargeRatio)
	}
	rows.Close()
}

func createTables(batteryDB *sql.DB) {
	createBatteryTable := `CREATE TABLE IF NOT EXISTS "battery_charge" (
		"id"	INTEGER NOT NULL UNIQUE,
		"charge"	INTEGER NOT NULL,
		"timestamp"	datetime NOT NULL,
		PRIMARY KEY("id" AUTOINCREMENT));`

	createBatteryPeriodTable := `CREATE TABLE IF NOT EXISTS "battery_last_period" (
			"id" INTEGER NOT NULL UNIQUE,
			"timestamp_start" datetime NOT NULL,
			"timestamp_end" datetime NOT NULL,
			"discharge" INTEGER NOT NULL,
			"discharge_ratio" FLOAT NOT NULL,
			"last_charge" INTEGER NOT NULL,
			FOREIGN KEY(last_charge) REFERENCES battery_charge(id),
			PRIMARY KEY("id" AUTOINCREMENT));`

	if createBatteryTablestatement, err := batteryDB.Prepare(createBatteryTable); err != nil {
		util.Check(err)
	} else {
		createBatteryTablestatement.Exec()
	}

	if createBatteryPeriodTableStatement, err := batteryDB.Prepare(createBatteryPeriodTable); err != nil {
		util.Check(err)
	} else {
		createBatteryPeriodTableStatement.Exec()
	}
}
