package main

import (
	"adventofcode/util"
	"log"
	"math"
	"strconv"
)

type Action struct {
	name  string
	value int
}

type Position struct {
	x           int
	y           int
	orientation string
}

var coordinates []string = []string{"E", "S", "W", "N"}

func main() {
	lines, err := util.ReadLines("./input.txt")

	if err != nil {
		log.Fatal(err)
	}
	var position Position = Position{0, 0, "E"}

	for _, line := range lines {
		value, _ := strconv.Atoi(line[1:])
		action := Action{string(line[0]), value}
		if action.name == "R" || action.name == "L" {
			position.setOrientation(action)
		} else {
			if action.name == "F" {
				position.movePosition(position.orientation, value)
			} else {
				position.movePosition(action.name, value)
			}
		}
	}

	println(manhattanDistance(position))
}

func (p *Position) movePosition(name string, value int) {
	switch name {
	case "N":
		p.x += value
	case "S":
		p.x -= value
	case "W":
		p.y -= value
	case "E":
		p.y += value
	}
}

func (p *Position) setOrientation(action Action) {
	var move int = action.value / 90
	if action.name == "L" {
		move = (move * -1) + len(coordinates)
	}
	var position int = (findCoordinate(p.orientation) + move) % len(coordinates)

	p.orientation = coordinates[position]
}

func findCoordinate(name string) int {
	for index := range coordinates {
		if coordinates[index] == name {
			return index
		}
	}

	return -1
}

func manhattanDistance(position Position) int {
	return int(math.Abs(float64(position.x)) + math.Abs(float64(position.y)))
}
