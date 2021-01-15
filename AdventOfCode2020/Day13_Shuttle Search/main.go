package main

import (
	"adventofcode/util"
	"log"
	"math"
	"strconv"
	"strings"
)

func main() {
	lines, err := util.ReadLines("./input.txt")

	if err != nil {
		log.Fatal(err)
	}

	timestamp, _ := strconv.Atoi(lines[0])
	var busIds []string = strings.Split(lines[1], ",")
	var min = math.MaxInt64
	var busID int = 0

	for _, item := range busIds {
		if item != "x" {
			id, _ := strconv.Atoi(item)
			var entire int = timestamp / id
			stop := (entire * id) + id
			if stop > timestamp {
				if stop-timestamp < min {
					min = stop - timestamp
					busID = id
				}
			}
		}
	}
	println(min * busID)
}
