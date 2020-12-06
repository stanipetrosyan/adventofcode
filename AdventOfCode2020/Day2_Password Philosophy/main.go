package main

import (
	"adventofcode/util"
	"log"
	"strconv"
	"strings"
)

type Password struct {
	policy []string
	char   byte
	value  string
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
	var solution = 0
	for _, line := range lines {
		password := passwordSerializator(line)
		count := strings.Count(password.value, string(password.char))
		min, _ := strconv.Atoi(password.policy[0])
		max, _ := strconv.Atoi(password.policy[1])
		if count >= min && count <= max {
			solution++
		}
	}
	println(solution)
}

func solvePartTwo(lines []string) {
	var solution = 0
	for _, line := range lines {
		password := passwordSerializator(line)
		start, _ := strconv.Atoi(password.policy[0])
		end, _ := strconv.Atoi(password.policy[1])
		if (password.value[start] == password.char) != (password.value[end] == password.char) {
			solution++
		}
	}
	println(solution)
}

func passwordSerializator(line string) Password {
	var password = Password{}
	line = strings.TrimSpace(line)
	var split = strings.Split(line, ":")
	var first = split[0]
	password.value = split[1]
	password.char = first[len(first)-1]
	password.policy = strings.Split(first[:len(first)-2], "-")
	return password
}
