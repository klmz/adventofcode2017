package days

import (
	"fmt"
	"strconv"
	"strings"

	"github.com/klmz/adventofcode2017/days"
	"github.com/spf13/cobra"
)

type day2 struct {
	Input string
	Part2 bool
}

var Day2 = day2{}

var (
	maxInt int = 1<<32 - 1
	minInt int = -1 << 31
)

func init() {
	AddDay(Day2)
}

func (day day2) AddFlags(cmd *cobra.Command) {
	cmd.Flags().StringVarP(&days.Input, "input", "i", "", "Set the input for day 1 ")
	cmd.Flags().BoolVarP(&Day2.Part2, "part2", "p2", false, "Calculate part 2")
}

func (day day2) Execute() {
	if Input == "" {
		fmt.Println("Input not set. Use: 'aoc day1 -i [input string]")
		return
	}
	day.Input = Input
	sum := Day2.CalculateAnswer()
	fmt.Println()
	fmt.Println("Answer:")
	fmt.Println(sum)
}

func (day day2) CalculateAnswer() int {
	lines := strings.Split(day.Input, "\n")
	checkSum := 0
	for _, line := range lines {
		tokens := strings.Split(line, "\t")
		max, min := minInt, maxInt

		for _, token := range tokens {
			digit, err := strconv.Atoi(token)
			if err != nil {
				panic(err)
			}

			if digit > max {
				max = digit
			}
			if digit < min {
				min = digit
			}
		}

		checkSum += max - min
	}
	return checkSum
}
