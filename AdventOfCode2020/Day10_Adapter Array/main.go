package main

import (
	"adventofcode/util"
	"log"
	"math"
	"sort"
)

func main() {
	numbers, err := util.ConvertToIntLines("./input.txt")

	if err != nil {
		log.Fatal(err)
	}

	sort.Ints(numbers)

	solvePartOne(numbers)
	solvePartTwo(numbers)
}

func solvePartOne(numbers []int) {
	var oneDifferences, threeDifferences int = 0, 1
	if numbers[0] == 1 {
		oneDifferences++
	} else {
		threeDifferences++
	}
	for index := 0; index < len(numbers)-1; {
		switch numbers[index+1] - numbers[index] {
		case 1:
			oneDifferences++
		case 3:
			threeDifferences++
		}
		index++
	}
	println(oneDifferences, threeDifferences)
}

func solvePartTwo(numbers []int) {
	var length int = 1
	var number int = 0
	var solution int = 1
	for index := 0; index < len(numbers)-1; {
		if numbers[index] <= number+3 {
			length++
			index++
		} else {
			solution *= countCombination(length, numbers[index]-numbers[index-1])
			number = numbers[index-1]
			length = 1
		}
	}
	println(solution)
}

func countCombination(len int, next int) int {
	if len == 4 && next == 1 {
		return 7
	}
	return int(math.Pow(2, float64(len-2)))
}
