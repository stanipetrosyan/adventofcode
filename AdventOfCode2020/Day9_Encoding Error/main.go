package main

import (
	"adventofcode/util"
	"log"
)

const INDEX = 25

func main() {
	lines, err := util.ConvertToIntLines("./input.txt")

	if err != nil {
		log.Fatal(err)
	}

	invalid := solvePartOne(lines)
	println(invalid)
	contiguosSet := solvePartTwo(lines, invalid)
	println(util.FindMax(contiguosSet) + util.FindMin(contiguosSet))
}

func solvePartOne(numbers []int) int {
	for index := INDEX; index < len(numbers[INDEX:]); {
		x := numbers[index]
		w := numbers[(index - INDEX):index]
		if isValid(w, x) == false {
			return x
		}
		index++
	}

	return 0
}

func solvePartTwo(numbers []int, invalid int) []int {
	var sum int = 0
	var start, end int = 0, 0

	for index := 0; index < len(numbers)-1; {
		if sum < invalid {
			sum += numbers[index]
			end++
			index++
		} else if sum > invalid {
			sum -= numbers[start]
			start++
		} else if sum == invalid {
			return numbers[start:end]
		}
	}

	return nil
}

func isValid(preamble []int, number int) bool {
	for _, item := range preamble {
		if find(preamble, number-item) {
			return true
		}
	}

	return false
}

func find(array []int, value int) bool {
	for _, item := range array {
		if item == value {
			return true
		}
	}

	return false
}
