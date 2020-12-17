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
	solvePartTwo(instructions)
}

func solvePartOne(instructions []Instruction) {
	tmp := make([]Instruction, len(instructions))
	copy(tmp, instructions)
	var accumulator int = 0

	for index := 0; tmp[index].exec == false; {
		tmp[index].exec = true
		accumulator, index = tmp[index].execInstruction(accumulator, index)
	}
	println(accumulator)
}

func solvePartTwo(instructions []Instruction) {

	for index, item := range instructions {
		channel := make(chan int)
		tmp := make([]Instruction, len(instructions))
		copy(tmp, instructions)
		tmp[index].exec = false
		if item.op == "jmp" {
			tmp[index].op = "nop"
		} else if item.op == "nop" {
			tmp[index].op = "jmp"
		}
		go runHeldhandGame(tmp, channel)
		solution := <-channel
		if solution != 0 {
			println(solution)
		}
	}
}

func runHeldhandGame(instructions []Instruction, ch chan<- int) {
	var accumulator int = 0
	for index := 0; instructions[index].exec == false; {
		instructions[index].exec = true
		accumulator, index = instructions[index].execInstruction(accumulator, index)
		if index == len(instructions) {
			ch <- accumulator
			break
		}
	}
	ch <- 0
}

func (i *Instruction) execInstruction(accumulator, index int) (int, int) {
	switch i.op {
	case "acc":
		accumulator += i.arg
		index++
	case "jmp":
		index += i.arg
	case "nop":
		index++
	}
	return accumulator, index
}
