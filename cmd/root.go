package cmd

import (
	"fmt"

	"github.com/spf13/cobra"
)

var RootCmd = &cobra.Command{
	Use:   "aoc",
	Short: "Compute the answers for Advent of Code 2017",
	Long: `Use this tool to compute the answers for Advent Of Code 2017. 
			Usage: 'aoc day 1' or 'aoc day 2'`,
	Run: func(cmd *cobra.Command, args []string) {
		fmt.Println(`Usage: 'aoc day 1' or 'aoc day 2'`)
	},
}
