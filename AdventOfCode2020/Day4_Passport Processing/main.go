package main

import (
	"adventofcode/util"
	"log"
	"strconv"
	"strings"
)

type Passport struct {
	complete []string
	byr      string
	iyr      string
	eyr      string
	hgt      string
	hcl      string
	ecl      string
	pid      string
	cid      string
	fields   map[string]string
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
	var solution int = 0
	var passportString string = ""

	for _, line := range lines {
		if line != "" {
			passportString = passportString + " " + line
		} else {
			passport := Passport{}
			//passport.passportSerialization(passportString)
			if passport.isValid(passportString) {
				solution++
			}
			passportString = ""
		}
	}
	println(solution)
}

func solvePartTwo(lines []string) {
	var solution int = 0
	var passportString string = ""

	for _, line := range lines {
		if line != "" {
			passportString = passportString + " " + line
		} else {
			passport := Passport{}
			passport.passportSerialization(passportString)
			if passport.isValid(passportString) {
				if passport.isStrictValid() {
					solution++
				}
			}
			passportString = ""
		}
	}
	println(solution)
}

func (p *Passport) passportSerialization(lines string) {
	p.complete = strings.Split(lines, " ")

	for _, field := range p.complete {
		fieldStructure := strings.Split(field, ":")
		fieldType := fieldStructure[0]
		//p.fields[fieldType] = fieldStructure[1] Usare la mappa risolverrebe un sacco di cose la devo guardare per bene
		switch fieldType {
		case "byr":
			p.byr = fieldStructure[1]
		case "iyr":
			p.iyr = fieldStructure[1]
		case "eyr":
			p.eyr = fieldStructure[1]
		case "hgt":
			p.hgt = fieldStructure[1]
		case "hcl":
			p.hcl = fieldStructure[1]
		case "ecl":
			p.ecl = fieldStructure[1]
		case "pid":
			p.pid = fieldStructure[1]
		case "cid":
			p.cid = fieldStructure[1]
		}
	}
}

func (p *Passport) isValid(lines string) bool {
	p.complete = strings.Split(lines, " ")
	completeLength := len(p.complete)
	if completeLength == 9 {
		return true
	}

	if completeLength == 8 {
		return !strings.Contains(lines, "cid")
	}
	return false
}

func (p *Passport) isStrictValid() bool {
	byr, _ := strconv.Atoi(p.byr)
	if (byr < 1920) || (byr > 2002) {
		return false
	}
	iyr, _ := strconv.Atoi(p.iyr)
	if (iyr < 2010) || (iyr > 2020) {
		return false
	}

	eyr, _ := strconv.Atoi(p.eyr)
	if (eyr < 2020) || (eyr > 2030) {
		return false
	}

	if strings.Contains(p.hgt, "cm") {
		hgt, _ := strconv.Atoi(p.hgt[:len(p.hgt)-2])
		if (hgt < 150) || (hgt > 193) {
			return false
		}
	} else if strings.Contains(p.hgt, "in") {
		hgt, _ := strconv.Atoi(p.hgt[:len(p.hgt)-2])
		if (hgt < 59) || (hgt > 76) {
			return false
		}
	} else {
		return false
	}
	if p.hcl[0] != '#' || len(p.hcl) != 7 {
		return false
	}
	if p.ecl != "amb" && p.ecl != "blu" && p.ecl != "brn" && p.ecl != "gry" && p.ecl != "grn" && p.ecl != "hzl" && p.ecl != "oth" {
		return false
	}
	return len(p.pid) == 9
}
