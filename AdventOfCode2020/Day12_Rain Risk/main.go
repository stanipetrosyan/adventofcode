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

	solvePartOne(lines)
	solvePartTwo(lines)
}

func solvePartOne(lines []string) {
	var ship Position = Position{0, 0, "E"}
	for _, line := range lines {
		value, _ := strconv.Atoi(line[1:])
		action := Action{string(line[0]), value}
		if action.name == "R" || action.name == "L" {
			ship.orientation = moveOrientation(action, ship.orientation)
		} else {
			if action.name == "F" {
				ship.movePosition(ship.orientation, value)
			} else {
				ship.movePosition(action.name, value)
			}
		}
	}

	println(manhattanDistance(ship))
}

func solvePartTwo(lines []string) {
	var waypoint Position = Position{10, 1, "E"}
	var ship Position = Position{0, 0, "E"}
	for _, line := range lines {
		value, _ := strconv.Atoi(line[1:])
		action := Action{string(line[0]), value}
		if action.name == "R" || action.name == "L" {
			waypoint.rotate(action)
		} else {
			if action.name == "F" {
				orientation := waypointOrientation(waypoint)
				ship.movePosition(orientation[0], int(math.Abs(float64(waypoint.x*value))))
				ship.movePosition(orientation[1], int(math.Abs(float64(waypoint.y*value))))
			} else {
				waypoint.movePosition(action.name, value)
			}
		}
	}

	println(manhattanDistance(ship))
}

func (p *Position) movePosition(name string, value int) {
	switch name {
	case "N":
		p.y += value
	case "S":
		p.y -= value
	case "W":
		p.x -= value
	case "E":
		p.x += value
	}
}

func moveOrientation(action Action, precOrientation string) string {
	var move int = action.value / 90
	if action.name == "L" {
		move = (move * -1) + len(coordinates)
	}
	return coordinates[(findCoordinate(precOrientation)+move)%len(coordinates)]
}

func findCoordinate(name string) int {
	for index := range coordinates {
		if coordinates[index] == name {
			return index
		}
	}

	return -1
}

func (p *Position) rotate(action Action) {
	tmpPosition := Position{0, 0, "E"}

	for _, item := range waypointOrientation(*p) {
		dimensioneRotation := moveOrientation(action, item)
		tmpPosition.setOrientation(*p, item, dimensioneRotation)
	}

	p.x = tmpPosition.x
	p.y = tmpPosition.y
}

func (tmpPosition *Position) setOrientation(p Position, orientation string, dimensioneRotation string) {
	if orientation == "E" {
		switch dimensioneRotation {
		case "N":
			tmpPosition.y = p.x
		case "W":
			tmpPosition.x = -p.x
		case "S":
			tmpPosition.y = -p.x
		}
	}

	if orientation == "N" {
		switch dimensioneRotation {
		case "E":
			tmpPosition.x = p.y
		case "W":
			tmpPosition.x = -p.y
		case "S":
			tmpPosition.y = -p.y
		}
	}

	if orientation == "S" {
		switch dimensioneRotation {
		case "E":
			tmpPosition.x = -p.y
		case "W":
			tmpPosition.x = p.y
		case "N":
			tmpPosition.y = -p.y
		}
	}

	if orientation == "W" {
		switch dimensioneRotation {
		case "E":
			tmpPosition.x = -p.x
		case "N":
			tmpPosition.y = -p.x
		case "S":
			tmpPosition.y = p.x
		}
	}
}

func waypointOrientation(p Position) []string {
	var orientation []string = []string{"", ""}
	if p.x > 0 {
		orientation[0] = "E"
	} else {
		orientation[0] = "W"
	}

	if p.y > 0 {
		orientation[1] = "N"
	} else {
		orientation[1] = "S"
	}

	return orientation
}

func manhattanDistance(ship Position) int {
	return int(math.Abs(float64(ship.x)) + math.Abs(float64(ship.y)))
}
