package cmd

import (
	"strconv"

	"github.com/klmz/adventofcode2017/days"
	"github.com/spf13/cobra"
)

func init() {
	RootCmd.AddCommand(dayCmd)
	dayCmd.Flags().StringVarP(&days.Input, "input", "i", "", "Set the input for day 1 ")

	for _, d := range days.Days {
		d.AddFlags(dayCmd)
	}
}

var dayCmd = &cobra.Command{
	Use:   "day",
	Short: "Return the answer for a specific day",
	Long:  `Return the answer for a specific day`,
	Run: func(cmd *cobra.Command, args []string) {

		day, err := strconv.ParseInt(args[0], 10, 64)
		if err != nil {
			println(err)
		}

		if day > int64(len(days.Days)) || day < 1 {
			println("That day doesn't exist.")
			return
		}
		days.Days[day-1].Execute()
	},
}
