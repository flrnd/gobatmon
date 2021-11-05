package util

import (
	"database/sql"
	"log"
	"os"

	_ "github.com/mattn/go-sqlite3"
)

func createTable(batteryDB *sql.DB) {
	createBatteryTable := `CREATE TABLE IF NOT EXISTS battery (
		"id" INTEGER PRIMARY KEY AUTOINCREMENT,
		"percentage" INTEGER NOT NULL,
		"timestamp" INTEGER NOT NULL,
		"FOREIGN KEY(percentage) REFERENCES percentage(id),
		"FOREIGN KEY(timestamp) REFERENCES timestamp(id),
	);`

	statement, err := batteryDB.Prepare(createBatteryTable)
	Check(err)
	statement.Exec()
}

func InitDatabase() {
	homeDir, err := os.UserHomeDir()
	Check(err)

	databasePath := homeDir + "/.config/batmon/battery.db"

	// Check if database exists
	if _, err := os.Stat(databasePath); os.IsNotExist(err) {
		os.MkdirAll(homeDir+"/.config/batmon", 0755)
		log.Println("Creating the database")
		file, err := os.Create(databasePath)
		Check(err)
		file.Close()
		log.Println("Database created")
	}

	// Open the database
	batteryDB, err := sql.Open("sqlite3", databasePath)
	Check(err)
	defer batteryDB.Close()
	createTable(batteryDB)
}
