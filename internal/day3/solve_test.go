package day3

import (
	_ "embed"
	"testing"
)

func TestDay3(t *testing.T) {
	tc := map[string]struct {
		input    string
		expected int
	}{"example": {example, 4361}, "actual": {input, 546312}}

	for name, tt := range tc {
		t.Run(name, func(t *testing.T) {
			answer, err := Solve(tt.input)
			if err != nil {
				t.Error(err)
			}
			if answer != tt.expected {
				t.Errorf("expected %d, got %d", tt.expected, answer)
			}
		})
	}
}

//go:embed testdata/test_input.txt
var example string

//go:embed testdata/input.txt
var input string
