package day3

import (
	"fmt"
	"strconv"
	"strings"
)

func Solve(input string) (int64, error) {
	lines := strings.Split(input, "\n")
	items := make([][]rune, len(lines))
	for r, line := range lines {
		items[r] = []rune(line)
	}

	numbers := make([]int, 0)
	for r, row := range items {
		number := make([]rune, 0, len(row))
		keepNumber := false
		for c, item := range row {
			if '0' <= item && item <= '9' {
				number = append(number, item)
				if keepNumber || hasAdjacentSymbol(items, r, c) {
					keepNumber = true
				}
				continue
			}
			if len(number) > 0 {
				if keepNumber {
					num, err := strconv.Atoi(string(number))
					if err != nil {
						return 0, err
					}
					numbers = append(numbers, num)
					fmt.Printf("Found number: %d\n", num)
				}
				number = make([]rune, 0, len(row))
				keepNumber = false
			}
		}
	}

	var answer int64
	for _, num := range numbers {
		answer += int64(num)
	}
	return answer, nil
}

func hasAdjacentSymbol(schema [][]rune, r, c int) bool {
	if r < 0 || c < 0 || r >= len(schema) || c >= len(schema[r]) {
		return false
	}
	for _, dr := range []int{-1, 0, 1} {
		nr := r + dr
		if nr < 0 || nr >= len(schema) {
			continue
		}
		for _, dc := range []int{-1, 0, 1} {
			nc := c + dc
			if nc < 0 || nc >= len(schema[nr]) {
				continue
			}
			if !isNumber(schema[nr][nc]) && schema[nr][nc] != '.' {
				return true
			}
		}
	}
	return false
}

func isNumber(r rune) bool {
	return '0' <= r && r <= '9'
}
