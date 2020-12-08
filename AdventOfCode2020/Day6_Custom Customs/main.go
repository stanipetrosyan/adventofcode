package main

import (
	"adventofcode/util"
	"log"
)

func main() {
	lines, err := util.ReadLines("./input.txt")

	if err != nil {
		log.Fatal(err)
	}

	solvePartOne(lines)
	solvePartTwo(lines)

}

func solvePartOne(lines []string) {
	var custom = make(map[byte]int)
	var solution int = 0

	for _, line := range lines {
		if line != "" {
			for _, char := range line {
				_, exist := custom[byte(char)]
				if !exist {
					custom[byte(char)] = 1
					solution++
				}
			}
		} else {
			custom = make(map[byte]int)
		}
	}

	println(solution)
}

func solvePartTwo(lines []string) {
	var custom = make(map[byte]int)
	var solution int = 0
	var count int = 0
	var group string = ""

	for _, line := range lines {
		if line != "" {
			count++
			group += line
		} else {
			for _, char := range group {
				_, exist := custom[byte(char)]
				if !exist {
					custom[byte(char)] = 1
				} else {
					custom[byte(char)] = custom[byte(char)] + 1
				}
				if (custom[byte(char)] - count) == 0 {
					solution++
				}
			}
			custom = make(map[byte]int)
			count = 0
			group = ""
		}
	}

	println(solution)
}
