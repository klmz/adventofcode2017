package days

import (
	"fmt"

	"github.com/spf13/cobra"
)

type dayOne struct {
	Input string
}

var Day1 = dayOne{}

func init() {
	AddDay(Day1)
}

func (day dayOne) AddFlags(cmd *cobra.Command) {

}

func (day dayOne) Execute() {
	if Input == "" {
		fmt.Println("Input not set. Use: 'aoc day1 -i [input string]")
		return
	}
	day.Input = Input
	sum := Day1.CalculateAnswer()
	fmt.Println()
	fmt.Println("Answer:")
	fmt.Println(sum)
}

func (day dayOne) CalculateAnswer() int {
	sum := 0
	for i, c := range Day1.Input {
		next := (i + 1) % len(Day1.Input)
		if Day1.Input[i] == Day1.Input[next] {
			digit := int(c) - '0'
			sum += digit
		}
	}

	return sum
}
