package main

import (
	"adventofcode/util"
	"log"
)

type Position struct {
	x int
	y int
}

func main() {
	lines, err := util.ReadLines("./input.txt")

	if err != nil {
		log.Fatal(err)
	}

	solvePartOne(lines)
	solvePartTwo(lines)

}

func solvePartOne(lines []string) {
	println(calcTreesEncountered(lines, 3, 1))
}

func solvePartTwo(lines []string) {
	println(calcTreesEncountered(lines, 1, 1) *
		calcTreesEncountered(lines, 3, 1) *
		calcTreesEncountered(lines, 5, 1) *
		calcTreesEncountered(lines, 7, 1) *
		calcTreesEncountered(lines, 1, 2))
}

func calcTreesEncountered(lines []string, right int, down int) int {
	var solution = 0
	var position int = -right

	for index := 0; index < len(lines); {
		line := lines[index]
		length := len(line) - 1
		if (position + right) > length {
			position = right - ((length + 1) - position)
		} else {
			position += right
		}
		if line[position] == '#' {
			solution++
		}

		index += down
	}
	return solution
}
