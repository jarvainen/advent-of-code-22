// Package cmd
/*
Copyright © 2022 Riku Järvinen
*/
package cmd

import (
	"fmt"
	"github.com/jarvainen/advent-of-code-22/calendar"
	"github.com/spf13/cobra"
	"os"
	"time"
)

// rootCmd represents the base command when called without any subcommands
var (
	day     int
	rootCmd = &cobra.Command{
		Use:   "advent-of-code-22",
		Short: "Advent of Code 2022 challenge runner",
		Long:  `TODO`,
		Run: func(cmd *cobra.Command, args []string) {
			solutionIndex := day - 1
			calendar.Solutions[solutionIndex]()
		},
	}
)

// Execute adds all child commands to the root command and sets flags appropriately.
func Execute() {
	err := rootCmd.Execute()
	if err != nil {
		fmt.Fprintln(os.Stderr, err)
		os.Exit(1)
	}
}

func getCurrentDateIfInCalendarRange() int {
	currentTime := time.Now()
	adventOfCodeTimezone := time.FixedZone("UTC-5", -5*60*60)
	startOfCalendar := time.Date(2022, time.December, 1, 0, 0, 0, 0, adventOfCodeTimezone)
	endOfCalendar := time.Date(2022, time.December, 25, 0, 0, 0, 0, adventOfCodeTimezone)
	if currentTime.After(startOfCalendar) && currentTime.Before(endOfCalendar) {
		return currentTime.Day()
	}

	return 1
}

func init() {
	rootCmd.Flags().IntVarP(&day, "day", "d", getCurrentDateIfInCalendarRange(), "Run challenge for given day")
}
