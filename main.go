package main

import (
	"github.com/flrnd/gobatmon/cmd"
	"github.com/flrnd/gobatmon/db"
)

func main() {
	db.Init()
	cmd.Execute()
}
