package main

import (
	"github.com/flrnd/gobatmon/cmd"
	"github.com/flrnd/gobatmon/util"
)

func main() {
	util.InitDatabase()
	cmd.Execute()
}
