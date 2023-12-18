package day3

import (
	"sort"
	"strconv"
	"strings"
	"unicode"
)

func Solve(input string) (int, error) {
	lines := strings.Fields(input)
	items := make([][]rune, len(lines))
	for r, line := range lines {
		items[r] = []rune(line)
	}

	symbols := make(map[Position]bool)
	for r, row := range items {
		for c, d := range row {
			if d != '.' && !unicode.IsDigit(d) {
				symbols[Position{r, c}] = true
			}
		}
	}

	partNumbers := make([]int, 0)
	for r, row := range items {
		isPartNumber := false
		digits := make([]rune, 0)
		for c, d := range row {
			if unicode.IsDigit(d) {
				digits = append(digits, d)
				if isPartNumber || hasAdjacentSymbol(symbols, Position{r, c}) {
					isPartNumber = true
				}
				continue
			}
			if len(digits) > 0 {
				if isPartNumber {
					partNumber, err := strconv.Atoi(string(digits))
					if err != nil {
						return 0, err
					}
					partNumbers = append(partNumbers, partNumber)
				}
				isPartNumber = false
				digits = make([]rune, 0)
			}
		}
	}

	// not necessary, but for better debugging
	sort.Slice(partNumbers, func(i, j int) bool {
		return partNumbers[i] < partNumbers[j]
	})

	answer := 0
	for _, partNumber := range partNumbers {
		answer += partNumber
	}
	return answer, nil
}

func hasAdjacentSymbol(symbols map[Position]bool, pos Position) bool {
	for _, diff := range []Position{{-1, -1}, {-1, 0}, {-1, 1}, {0, -1}, {0, 1}, {1, -1}, {1, 0}, {1, 1}} {
		p := pos.AddPosition(diff)
		_, exist := symbols[p]
		if exist {
			return true
		}
	}
	return false
}

type Position struct {
	r, c int
}

func (p Position) AddPosition(diff Position) Position {
	return Position{p.r + diff.r, p.c + diff.c}
}
