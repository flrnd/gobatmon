package cmd

import (
	"fmt"
	"os"

	"github.com/muesli/coral"
)

var deleteAllCmd = &coral.Command{
	Use:   "all",
	Short: "Delete all timestamps. Can't be undone",
	Long:  "Delete all created timestamps, this action can't be undone.",
	Run: func(cmd *coral.Command, args []string) {
		fmt.Printf("All timestamp deleted!!")
		os.Exit(1)
	},
}
