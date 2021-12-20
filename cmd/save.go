package cmd

import (
	"fmt"

	"github.com/spf13/cobra"
)

var cmdSave = &cobra.Command{
	Use:   "save",
	Short: "Save last period",
	Long:  "",
	Run: func(cmd *cobra.Command, args []string) {
		fmt.Println("Saved")
	},
}
