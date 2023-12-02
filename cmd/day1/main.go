package main

import (
	_ "embed"
	"fmt"
	"strings"
)

//go:embed input.txt
var input string

func main() {
	answer := 0
	lines := strings.Split(input, "\n")
	for _, line := range lines {
		numbers := make([]int, 0, len(line))
		for _, char := range line {
			if '0' < char && char <= '9' {
				numbers = append(numbers, int(char-'0'))
			}
		}
		number := numbers[0]*10 + numbers[len(numbers)-1]
		fmt.Printf("number: %d\n", number)
		answer += number
	}
	fmt.Printf("answer: %d\n", answer)
}
