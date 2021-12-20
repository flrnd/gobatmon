package main

import (
	"github.com/flrnd/gobatmon/cmd"
	database "github.com/flrnd/gobatmon/db"
)

func main() {
	database.Init()
	cmd.Execute()
}
