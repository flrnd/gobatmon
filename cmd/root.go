package cmd

import (
	"fmt"
	"os"

	"github.com/muesli/coral"
)

var rootCmd = &coral.Command{Use: "gobatmon"}

func Execute() {
	if err := rootCmd.Execute(); err != nil {
		fmt.Fprintln(os.Stderr, err)
		os.Exit(1)
	}
}
