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
	solve(seats, 4, false)
	solve(seats, 5, true)

}

func solve(seats [][]string, near int, partTwo bool) {
	matrix, variation := anotherRound(seats, near, partTwo)

	for variation == true {
		matrix, variation = anotherRound(matrix, near, partTwo)
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

func anotherRound(seats [][]string, near int, partTwo bool) ([][]string, bool) {
	round := duplicateMatrix(seats)
	var changed bool = false
	var adjacent int = 0
	for row := range seats {
		for col := range seats[row] {
			if partTwo == false {
				adjacent = getAllAdjacent(seats, row, col, len(seats[row]))
			} else {
				adjacent = getAllAdjacentInRow(seats, row, col, len(seats[row]))
			}
			if seats[row][col] == "L" && adjacent == 0 {
				round[row][col] = "#"
				changed = true
			}
			if seats[row][col] == "#" && adjacent >= near {
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

func getAdjacent(x, y int, dx, dy int, a [][]string, lenRow int) string {
	if !isValid(x, y, len(a), lenRow) {
		return "."
	}

	if a[x][y] == "." {
		return getAdjacent(x+dx, y+dy, dx, dy, a, lenRow)
	}

	return a[x][y]

}

func getAllAdjacentInRow(a [][]string, x, y int, lenRow int) int {
	var number string = ""
	var directions []int = []int{1, 0, -1}

	for _, dx := range directions {
		for _, dy := range directions {
			if !(dx == 0 && dy == 0) {
				number += getAdjacent(x+dx, y+dy, dx, dy, a, lenRow)
			}
		}
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
