package day2

import (
	_ "embed"
	"log"
	"strconv"
	"strings"
)

func SolvePart1(input string, limit Cube) int {
	answer := 0
	lines := strings.Split(input, "\n")
	for _, line := range lines {
		game := NewGame(line)
		if game.IsPossible(limit) {
			answer += game.id
		}
	}
	return answer
}

func SolvePart2(input string) int {
	answer := 0
	for _, line := range strings.Split(input, "\n") {
		game := NewGame(line)
		minCube := Cube{}
		for _, cube := range game.cubes {
			if minCube.red < cube.red {
				minCube.red = cube.red
			}
			if minCube.green < cube.green {
				minCube.green = cube.green
			}
			if minCube.blue < cube.blue {
				minCube.blue = cube.blue
			}
		}
		answer += minCube.red * minCube.green * minCube.blue
	}
	return answer
}

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

func (g Game) IsPossible(limit Cube) bool {
	for _, cube := range g.cubes {
		if limit.red < cube.red {
			return false
		}
		if limit.green < cube.green {
			return false
		}
		if limit.blue < cube.blue {
			return false
		}
	}
	return true
}
