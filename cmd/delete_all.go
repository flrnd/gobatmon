package cmd

import (
	"fmt"

	"github.com/flrnd/gobatmon/db"
	_ "github.com/mattn/go-sqlite3"
	"github.com/spf13/cobra"
)

var deleteAllCmd = &cobra.Command{
	Use:   "all",
	Short: "Delete all timestamps. Can't be undone",
	Long:  "Delete all created timestamps, this action can't be undone.",
	Run: func(cmd *cobra.Command, args []string) {
		var input string

		fmt.Printf("You are about to delete ALL timestamps. This action can't be undone.\n\nContinue [y/n ...? type Y to continue (n): ")
		fmt.Scanf("%s", &input)

		if input == "Y" {
			db.DeleteAll()
		}
	},
}
