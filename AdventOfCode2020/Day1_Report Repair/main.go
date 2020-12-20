package main

import (
	"adventofcode/util"
	"log"
	"sort"
)

const PROBLEM = 2020

func main() {
	lines, err := util.ConvertToIntLines("./input.txt")

	if err != nil {
		log.Fatal(err)
	}

	sort.Ints(lines)
	println(solvePartOne(lines, 0, len(lines)-1))
	println(solvePartTwo(lines, 0, 1, len(lines)-1))
}

func solvePartOne(numbers []int, start int, end int) int {
	if (numbers[start] + numbers[end]) == PROBLEM {
		return numbers[end] * numbers[start]
	}
	if (numbers[start] + numbers[end]) > PROBLEM {
		return solvePartOne(numbers, start, end-1)
	}
	return solvePartOne(numbers, start+1, end)
}

func solvePartTwo(numbers []int, start int, middle int, end int) int {
	if (numbers[start] + numbers[middle] + numbers[end]) == PROBLEM {
		return numbers[start] * numbers[middle] * numbers[end]
	}
	if (numbers[start] + numbers[middle] + numbers[end]) > PROBLEM {
		return solvePartTwo(numbers, start, middle, end-1)
	}
	return solvePartTwo(numbers, start+1, start+2, end)
}
