// Package cmd
/*
Copyright © 2022 Riku Järvinen
*/
package cmd

import (
	"github.com/jarvainen/advent-of-code-22/calendar"
	"github.com/spf13/cobra"
)

// allCmd represents the all command
var allCmd = &cobra.Command{
	Use:   "all",
	Short: "Run all challenges",
	Run: func(cmd *cobra.Command, args []string) {
		for _, solution := range calendar.Solutions {
			solution()
		}
	},
}

func init() {
	rootCmd.AddCommand(allCmd)
}
