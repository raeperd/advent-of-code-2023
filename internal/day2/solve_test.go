package day2

import (
	_ "embed"
	"testing"
)

//go:embed testdata/input.txt
var input string

func TestDay1(t *testing.T) {
	limit := Cube{
		red:   12,
		green: 13,
		blue:  14,
	}

	want := 2006
	answer := SolvePart1(input, limit)
	if answer != want {
		t.Errorf("SolvePart1() = %v, want %v", answer, want)
	}

	want = 84911
	answer = SolvePart2(input)
	if answer != want {
		t.Errorf("SolvePart2() = %v, want %v", answer, want)
	}
}
