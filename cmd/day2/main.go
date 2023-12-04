package main

import (
	_ "embed"
	"fmt"
	"log"
	"strconv"
	"strings"
)

//go:embed input.txt
var input string

type Game struct {
	id    int
	cubes []Cube
}

type Cube struct {
	red   int
	green int
	blue  int
}

func NewGame(line string) Game {
	items := strings.Split(line, ": ")
	if len(items) == 0 {
		log.Fatal("error")
	}
	id, err := strconv.Atoi(strings.TrimPrefix(items[0], "Game "))
	if err != nil {
		log.Fatal(err)
	}
	game := Game{id: id, cubes: make([]Cube, 0, 3)}

	rgbs := strings.Split(items[1], "; ")
	for _, rgb := range rgbs {
		cubes := strings.Split(rgb, ", ")
		c := Cube{}
		for _, cube := range cubes {
			terms := strings.Split(cube, " ")
			if len(terms) != 2 {
				log.Fatal("term has not length 2", terms)
			}
			i, err := strconv.Atoi(terms[0])
			if err != nil {
				log.Fatal(err)
			}
			switch terms[1] {
			case "red":
				c.red = i
			case "green":
				c.green = i
			case "blue":
				c.blue = i
			}
		}
		game.cubes = append(game.cubes, c)
	}
	return game
}

func main() {
	answer := 0
	limit := Cube{
		red:   12,
		green: 13,
		blue:  14,
	}
	lines := strings.Split(input, "\n")
	for _, line := range lines {
		game := NewGame(line)
		possible := true
		for _, cube := range game.cubes {
			if limit.red < cube.red {
				possible = false
				break
			}
			if limit.green < cube.green {
				possible = false
				break
			}
			if limit.blue < cube.blue {
				possible = false
				break
			}
		}
		if possible {
			answer += game.id
		}
	}
	fmt.Printf("answer: %d\n", answer)
}
