package main

import (
	"adventofcode/util"
	"log"
	"strconv"
	"strings"
)

type Passport struct {
	fields map[string]string
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
			passport.passportSerialization(passportString)
			if passport.isValid() {
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
			if passport.isValid() && passport.isStrictValid() {
				solution++
			}
			passportString = ""
		}
	}
	println(solution)
}

func (p *Passport) passportSerialization(lines string) {
	passport := strings.Split(lines, " ")
	fields := make(map[string]string)
	for _, field := range passport[1:] {
		fieldStructure := strings.Split(field, ":")
		fieldType := fieldStructure[0]
		fields[fieldType] = fieldStructure[1]
	}
	p.fields = fields
}

func (p *Passport) isValid() bool {
	fields := len(p.fields)
	if fields == 8 {
		return true
	}

	if fields == 7 {
		_, exist := p.fields["cid"]
		return !exist
	}
	return false
}

func (p *Passport) isStrictValid() bool {
	byr, _ := strconv.Atoi(p.fields["byr"])
	if (byr < 1920) || (byr > 2002) {
		return false
	}
	iyr, _ := strconv.Atoi(p.fields["iyr"])
	if (iyr < 2010) || (iyr > 2020) {
		return false
	}

	eyr, _ := strconv.Atoi(p.fields["eyr"])
	if (eyr < 2020) || (eyr > 2030) {
		return false
	}

	if strings.Contains(p.fields["hgt"], "cm") {
		hgt, _ := strconv.Atoi(p.fields["hgt"][:len(p.fields["hgt"])-2])
		if (hgt < 150) || (hgt > 193) {
			return false
		}
	} else if strings.Contains(p.fields["hgt"], "in") {
		hgt, _ := strconv.Atoi(p.fields["hgt"][:len(p.fields["hgt"])-2])
		if (hgt < 59) || (hgt > 76) {
			return false
		}
	} else {
		return false
	}
	if p.fields["hcl"][0] != '#' || len(p.fields["hcl"]) != 7 {
		return false
	}
	if p.fields["ecl"] != "amb" && p.fields["ecl"] != "blu" && p.fields["ecl"] != "brn" && p.fields["ecl"] != "gry" && p.fields["ecl"] != "grn" && p.fields["ecl"] != "hzl" && p.fields["ecl"] != "oth" {
		return false
	}
	return len(p.fields["pid"]) == 9
}
