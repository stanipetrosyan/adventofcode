package main

import (
	"bufio"
	"log"
	"os"
	"sort"
	"strconv"
)

const PROBLEM = 2020

func main() {
	lines, err := convertToIntLines("./input.txt")

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

func convertToIntLines(path string) ([]int, error) {
	file, err := os.Open(path)

	if err != nil {
		return nil, err
	}

	defer file.Close()

	scanner := bufio.NewScanner(file)
	scanner.Split(bufio.ScanLines)

	var lines []int

	for scanner.Scan() {
		n, _ := strconv.Atoi(scanner.Text())
		lines = append(lines, n)
	}

	return lines, nil
}
