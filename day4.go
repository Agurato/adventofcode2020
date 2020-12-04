package main

import (
	"fmt"
	"regexp"
	"strconv"
	"strings"
)

// Day4 Day 4
func (aoc AOC) Day4() {
	values, err := ReadLines("inputs/day4.txt")
	if err != nil {
		panic(err)
	}
	fmt.Println(day4Part1(values))
	fmt.Println(day4Part2(values))
}

func day4Part1(values []string) int {
	var newLines []string
	passport := ""
	for _, line := range values {
		if line == "" {
			newLines = append(newLines, passport)
			passport = ""
			continue
		}
		passport += line + " "
	}
	newLines = append(newLines, passport)

	count := 0
	for _, line := range newLines {
		colonNumber := strings.Count(line, ":")
		if colonNumber == 8 || (colonNumber == 7 && !strings.Contains(line, "cid")) {
			count++
		}
	}

	return count
}

func day4Part2(values []string) int {
	var newLines []string
	passport := ""
	for _, line := range values {
		if line == "" {
			newLines = append(newLines, passport)
			passport = ""
			continue
		}
		passport += line + " "
	}
	newLines = append(newLines, passport)

	count := 0
	for _, line := range newLines {
		colonNumber := strings.Count(line, ":")
		if !(colonNumber == 8 || (colonNumber == 7 && !strings.Contains(line, "cid"))) {
			continue
		}
		pairs := strings.Split(line, " ")
		pairsLen := len(pairs)
		isValid := true
		for pairI := 0; isValid && pairI < pairsLen; pairI++ {
			pair := pairs[pairI]
			if pair == "" {
				continue
			}
			separated := strings.Split(pair, ":")
			value := separated[1]
			switch separated[0] {
			case "byr":
				intValue, err := strconv.Atoi(value)
				if err != nil {
					isValid = false
				}
				if intValue < 1920 || intValue > 2002 {
					isValid = false
				}
				break
			case "iyr":
				intValue, err := strconv.Atoi(value)
				if err != nil || intValue < 2010 || intValue > 2020 {
					isValid = false
				}
				break
			case "eyr":
				intValue, err := strconv.Atoi(value)
				if err != nil || intValue < 2020 || intValue > 2030 {
					isValid = false
				}
				break
			case "hgt":
				unit := value[len(value)-2:]
				intValue, err := strconv.Atoi(value[:len(value)-2])
				if err != nil || (unit != "cm" && unit != "in") || (unit == "cm" && (intValue < 150 || intValue > 193)) || (unit == "in" && (intValue < 59 || intValue > 76)) {
					isValid = false
				}
			case "hcl":
				colorRegex, _ := regexp.Compile("^#[0-9a-f]{6}$")
				if !colorRegex.MatchString(value) {
					isValid = false
				} else {
					isValid = true
				}
				break
			case "ecl":
				if value != "amb" && value != "blu" && value != "brn" && value != "gry" && value != "grn" && value != "hzl" && value != "oth" {
					isValid = false
				}
				break
			case "pid":
				idRegex, _ := regexp.Compile("^[0-9]{9}$")
				if !idRegex.MatchString(value) {
					isValid = false
				}
				break
			}
		}
		if isValid {
			count++
		}
	}

	return count
}

// Passport holds passport info
type Passport struct {
	byr string
	iyr string
	eyr string
	hgt string
	hcl string
	ecl string
	pid string
	cid string
}

// HasRequiredFields checks if a passport has all required fields
func (p Passport) HasRequiredFields() bool {
	return p.byr != "" && p.iyr != "" && p.eyr != "" && p.hgt != "" && p.hcl != "" && p.ecl != "" && p.pid != ""
}
