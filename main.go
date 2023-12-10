package main

import (
	_ "embed"
	"fmt"
	"log"

	"github.com/raeperd/advent-of-code-2023/internal/day3"
)

//go:embed internal/day3/testdata/input.txt
var input string

//go:embed internal/day3/testdata/test_input.txt
var testInput string

func main() {
	answer, err := day3.Solve(input)
	if err != nil {
		log.Fatalf("Solve() = %v", err)
	}
	fmt.Printf("Answer: %v\n", answer)
}
