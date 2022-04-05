package cmd

import (
	"fmt"
	"os"
	"strconv"

	"github.com/flrnd/gobatmon/db"
	"github.com/flrnd/gobatmon/util"
	_ "github.com/mattn/go-sqlite3"
	"github.com/spf13/cobra"
)

func init() {
	rootCmd.AddCommand(deleteCmd)
	deleteCmd.AddCommand(deleteAllCmd)
}

var deleteCmd = &cobra.Command{
	Use:   "delete id",
	Short: "Deletes a timestamp",
	Long:  `Delete a timestamp passing an id or delete all timestamps with all command. This action can't be undone.`,
	Run: func(cmd *cobra.Command, args []string) {
		if len(args) < 1 {
			fmt.Printf("Missing id.\nExample Usage: gobatmon delete 3\n")
			os.Exit(1)
		}
		intVarId, err := strconv.Atoi(args[0])
		util.Check(err)
		db.Delete(intVarId, "battery_charge")
	},
}
