package days

import (
	"testing"

	"github.com/klmz/adventofcode2017/test"
)

type AnswerCase struct {
	Input  string
	Answer int
}

func TestCalculateAnswer(t *testing.T) {
	cases := []AnswerCase{
		AnswerCase{"1111", 4},
		AnswerCase{"1112", 2},
		AnswerCase{"1122", 3},
		AnswerCase{"1234", 0},
		AnswerCase{"91212129", 9},
	}

	for _, c := range cases {
		Day1.Input = c.Input
		test.Equals(t, Day1.CalculateAnswer(), c.Answer)
	}

}
