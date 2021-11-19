package util

import (
	"github.com/flrnd/gobatmon/db"
	_ "github.com/mattn/go-sqlite3"
)

func InitDatabase() {
	db.Init()
}
