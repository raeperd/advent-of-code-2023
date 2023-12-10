package day1

import (
	_ "embed"
	"testing"
)

//go:embed testdata/input.txt
var input string

func TestDay1(t *testing.T) {
	answer, err := Solve(input)
	if err != nil {
		t.Errorf("Solve() error = %v", err)
	}
	want := 54877
	if answer != want {
		t.Errorf("Solve() = %v, want %v", answer, want)
	}
}
