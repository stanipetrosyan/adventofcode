package main

import (
	"adventofcode/util"
	"log"
	"sort"
)

func main() {
	lines, err := util.ReadLines("./input.txt")

	if err != nil {
		log.Fatal(err)
	}
	var solution int = 0
	var ids []int
	for _, line := range lines {
		seatID := (recRow(line[:7], 0, 0, 127) * 8) + recColumn(line[7:], 0, 0, 7)
		ids = append(ids, seatID)
		if seatID > solution {
			solution = seatID
		}
	}

	println(solution)

	sort.Ints(ids)
	for i := 1; i < len(ids); i++ {
		if ids[i]-ids[i-1] == 2 {
			println(ids[i] - 1)
		}
	}
}

func recRow(line string, index int, min int, max int) int {
	if index == len(line)-1 {
		if line[index] == 'F' {
			return min
		}
		return max
	}
	if line[index] == 'F' {
		return recRow(line, index+1, min, (min+max)/2)
	}
	return recRow(line, index+1, ((min+max)/2)+1, max)
}

func recColumn(line string, index int, min int, max int) int {
	if index == len(line)-1 {
		if line[index] == 'L' {
			return min
		}
		return max
	}
	if line[index] == 'L' {
		return recColumn(line, index+1, min, (min+max)/2)
	}
	return recColumn(line, index+1, ((min+max)/2)+1, max)
}
