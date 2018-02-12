package test

import "testing"

func Equals(t *testing.T, actual, expected int) {
	if expected != actual {
		t.Errorf("Sum was incorrect, got: %d, want: %d.", actual, expected)
	}
}
