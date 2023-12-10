package day3

import (
	"fmt"
	"strconv"
	"strings"
)

func Solve(input string) (int, error) {
	lines := strings.Split(input, "\n")
	items := make([][]rune, len(lines))
	for r, line := range lines {
		items[r] = []rune(line)
	}

	windowToNumber := make(map[Window]string)
	for r, row := range items {
		number := make([]rune, 0, len(row))
		for c, item := range row {
			if '0' <= item && item <= '9' {
				number = append(number, item)
				continue
			}
			if len(number) > 0 {
				windowToNumber[Window{Position{r - 1, c - len(number) - 1}, Position{r + 1, c}}] = string(number)
				number = make([]rune, 0, len(row))
			}
		}
	}

	answers := make([]int, 0, len(windowToNumber))
	for window, number := range windowToNumber {
		found := false
		for r := window.from.r; r <= window.to.r; r++ {
			if found {
				break
			}
			for c := window.from.c; c <= window.to.c; c++ {
				if found {
					break
				}
				if r < 0 || c < 0 || r >= len(items) || c >= len(items[r]) {
					continue
				}
				if '0' <= items[r][c] && items[r][c] <= '9' {
					continue
				}
				if items[r][c] == '.' {
					continue
				}
				fmt.Printf("found number: %s with symbol: %c in %d,%d\n", number, items[r][c], r, c)
				num, err := strconv.Atoi(number)
				if err != nil {
					return 0, err
				}
				answers = append(answers, num)
				found = true
			}
		}
	}

	answer := 0
	for _, a := range answers {
		answer += a
	}
	return answer, nil
}

type Window struct {
	from, to Position
}

type Position struct {
	r, c int
}
