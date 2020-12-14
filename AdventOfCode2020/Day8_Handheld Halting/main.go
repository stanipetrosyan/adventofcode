package main

import (
	"adventofcode/util"
	"log"
	"strconv"
	"strings"
)

type Instruction struct {
	op   string
	arg  int
	exec bool
}

func main() {
	lines, err := util.ReadLines("./input.txt")

	if err != nil {
		log.Fatal(err)
	}
	var instructions = []Instruction{}

	for _, line := range lines {
		split := strings.Split(line, " ")
		arg, _ := strconv.Atoi(split[1])
		instructions = append(instructions, Instruction{split[0], arg, false})
	}

	solvePartOne(instructions)

}

func solvePartOne(instructions []Instruction) {
	var accumulator int = 0

	for index := 0; instructions[index].exec == false; {
		instructions[index].exec = true
		arg := instructions[index].arg

		switch instructions[index].op {
		case "acc":
			accumulator += arg
			index++
		case "jmp":
			index += arg
		case "nop":
			index++
		}
	}
	println(accumulator)
}
