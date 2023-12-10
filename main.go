package main

import (
	_ "embed"
	"fmt"

	"github.com/raeperd/advent-of-code-2023/internal/day2"
)

//go:embed internal/day2/testdata/input.txt
var input string

func main() {
	answer := day2.SolvePart2(input)
	fmt.Printf("Answer: %v\n", answer)
}
