package day1

import (
	_ "embed"
	"strings"
)

func Solve(input string) (int, error) {
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
		answer += number
	}
	return answer, nil
}
