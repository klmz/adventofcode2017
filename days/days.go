package days

import "github.com/spf13/cobra"

var Days []Day
var Input string

type Day interface {
	Execute()
	AddFlags(cmd *cobra.Command)
}

func AddDay(day Day) {
	Days = append(Days, day)
}
