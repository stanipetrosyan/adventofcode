package main

import (
	"adventofcode/util"
	"log"
	"strings"
)

func main() {
	lines, err := util.ReadLines("./input.txt")

	if err != nil {
		log.Fatal(err)
	}

	var seats [][]string = [][]string{}

	for _, line := range lines {
		seats = append(seats, strings.Split(line, ""))
	}

	matrix, variation := anotherRound(seats)

	for variation == true {
		matrix, variation = anotherRound(matrix)
	}
	var solution int = 0
	for _, item := range matrix {
		for _, seat := range item {
			if seat == "#" {
				solution++
			}
		}
	}
	println(solution)
}

func anotherRound(seats [][]string) ([][]string, bool) {
	round := duplicateMatrix(seats)
	var changed bool = false
	for row := range seats {
		for col := range seats[row] {
			adjacent := getAllAdjacent(seats, row, col, len(seats[row]))
			if seats[row][col] == "L" && adjacent == 0 {
				round[row][col] = "#"
				changed = true
			}
			if seats[row][col] == "#" && adjacent >= 4 {
				round[row][col] = "L"
				changed = true
			}
		}

	}
	return round, changed
}

func getAllAdjacent(a [][]string, x, y int, lenRow int) int {
	var number string = ""
	if isValid(x+1, y-1, len(a), lenRow) {
		number += a[x+1][y-1]
	}
	if isValid(x-1, y+1, len(a), lenRow) {
		number += a[x-1][y+1]
	}
	if isValid(x-1, y-1, len(a), lenRow) {
		number += a[x-1][y-1]
	}
	if isValid(x+1, y+1, len(a), lenRow) {
		number += a[x+1][y+1]
	}
	if isValid(x+1, y, len(a), lenRow) {
		number += a[x+1][y]
	}
	if isValid(x, y+1, len(a), lenRow) {
		number += a[x][y+1]
	}
	if isValid(x, y-1, len(a), lenRow) {
		number += a[x][y-1]
	}
	if isValid(x-1, y, len(a), lenRow) {
		number += a[x-1][y]
	}
	return strings.Count(number, "#")
}

func isValid(x int, y int, n int, w int) bool {
	if x < 0 || y < 0 || x > n-1 || y > w-1 {
		return false
	}
	return true
}

func duplicateMatrix(matrix [][]string) [][]string {
	duplicate := make([][]string, len(matrix))
	for i := range matrix {
		duplicate[i] = make([]string, len(matrix[i]))
		copy(duplicate[i], matrix[i])
	}
	return duplicate
}
