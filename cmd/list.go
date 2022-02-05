package cmd

import (
	"github.com/flrnd/gobatmon/db"
	"github.com/muesli/coral"
)

func init() {
	rootCmd.AddCommand(listCmd)
	listCmd.AddCommand(savedCmd)
}

var listCmd = &coral.Command{
	Use:   "list",
	Short: "list all timestamps. [help list] for more information.",
	Long:  "Output all stored timestamps. list saved shows stored discharging periods.",
	Run: func(cmd *coral.Command, args []string) {
		db.List()
	},
}
