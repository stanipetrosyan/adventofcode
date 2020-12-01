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
	solvePartOne(lines)
	solvePartTwo(lines)
}

func solvePartOne(numbers []int) {
	start := 0
	end := len(numbers) - 1

	for (numbers[start] + numbers[end]) != PROBLEM {
		if (numbers[start] + numbers[end]) > PROBLEM {
			end--
		} else {
			start++
		}
	}

	println(numbers[end] * numbers[start])
}

func solvePartTwo(numbers []int) {
	start := 0
	middle := 1
	end := len(numbers) - 1

	for (numbers[start] + numbers[middle] + numbers[end]) != PROBLEM {
		if (numbers[start] + numbers[middle] + numbers[end]) > PROBLEM {
			end--
		} else {
			for (numbers[start] + numbers[middle] + numbers[end]) < PROBLEM {
				middle++
			}
			start++
			middle = start + 1
		}
	}

	println(numbers[end] * numbers[middle] * numbers[start])
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
