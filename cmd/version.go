package cmd

import (
	"fmt"

	"github.com/muesli/coral"
)

func init() {
	rootCmd.AddCommand(versionCmd)
}

var versionCmd = &coral.Command{
	Use:   "version",
	Short: "Print the version number of gobatmon",
	Long:  `All software has versions.`,
	Run: func(cmd *coral.Command, args []string) {
		fmt.Println("gobatmon v0.1")
	},
}
