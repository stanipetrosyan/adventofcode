package main

import (
	"adventofcode/util"
	"fmt"
	"log"
	"strconv"
	"strings"
)

const MYBAG = "shinygold"

type Bag struct {
	name     string
	contains []Rule
}

type Rule struct {
	number int
	bag    string
}

func main() {

	lines, err := util.ReadLines("./input.txt")

	if err != nil {
		log.Fatal(err)
	}
	var rules = make(map[string]Bag)

	for _, line := range lines {
		line = strings.Replace(line, ".", "", -1)
		bag := setRule(line)
		rules[bag.name] = bag
	}

	var solution int = 0
	for _, bag := range rules {
		if bag.name != MYBAG {
			solution += solvePartOne(rules, bag)
		}
	}

	println(solution)

	fmt.Println(solvePartTwo(rules, rules[MYBAG], 1) - 1)
}

func solvePartOne(rules map[string]Bag, bag Bag) int {
	if bag.name == MYBAG {
		return 1
	}

	for _, contain := range bag.contains {
		if solvePartOne(rules, rules[contain.bag]) == 1 {
			return 1
		}
	}
	return 0
}

func solvePartTwo(rules map[string]Bag, bag Bag, level int) int {
	if len(bag.contains) == 0 {
		return level
	}

	for _, rule := range bag.contains {
		solvePartTwo(rules, rules[rule.bag], level*rule.number)
	}
	return level
}

func setRule(line string) Bag {
	bag := Bag{}
	rule := strings.Split(line, " ")
	if !(rule[3]+rule[4] == "contain"+"no") {
		for i := 0; i <= strings.Count(line, ","); i++ {
			index := 4 + (i * 4)
			number, _ := strconv.Atoi(rule[index])
			rule := Rule{number, rule[index+1] + rule[index+2]}
			bag.contains = append(bag.contains, rule)
		}
	}
	bag.name = rule[0] + rule[1]
	return bag
}
