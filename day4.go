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
	count := 0
	fieldNb := 0
	for _, line := range values {
		if line == "" {
			if fieldNb == 7 {
				count++
			}
			fieldNb = 0
			continue
		}
		fieldNb += strings.Count(line, ":") - strings.Count(line, "cid")
	}
	if fieldNb == 7 {
		count++
	}

	return count
}

func day4Part2(lines []string) int {
	colorRegex, _ := regexp.Compile("^#[0-9a-f]{6}$")
	idRegex, _ := regexp.Compile("^[0-9]{9}$")
	isValid := true
	count := 0
	fieldNb := 0
	linesLen := len(lines)
	for lineI := 0; lineI < linesLen; lineI++ {
		line := lines[lineI]
		if line == "" {
			if isValid && fieldNb == 7 {
				count++
			}
			isValid = true
			fieldNb = 0
			continue
		}
		if !isValid {
			isValid = true
			fieldNb = 0
			continue
		}
		pairs := strings.Split(line, " ")
		pairsLen := len(pairs)
		for pairI := 0; isValid && pairI < pairsLen; pairI++ {
			pair := pairs[pairI]
			if pair == "" {
				continue
			}
			key := pair[:3]
			value := pair[4:]
			switch key {
			case "byr":
				intValue, err := strconv.Atoi(value)
				if err != nil {
					isValid = false
				}
				if intValue < 1920 || intValue > 2002 {
					isValid = false
				}
				fieldNb++
				break
			case "iyr":
				intValue, err := strconv.Atoi(value)
				if err != nil || intValue < 2010 || intValue > 2020 {
					isValid = false
				}
				fieldNb++
				break
			case "eyr":
				intValue, err := strconv.Atoi(value)
				if err != nil || intValue < 2020 || intValue > 2030 {
					isValid = false
				}
				fieldNb++
				break
			case "hgt":
				unit := value[len(value)-2:]
				intValue, err := strconv.Atoi(value[:len(value)-2])
				if err != nil || (unit != "cm" && unit != "in") || (unit == "cm" && (intValue < 150 || intValue > 193)) || (unit == "in" && (intValue < 59 || intValue > 76)) {
					isValid = false
				}
				fieldNb++
			case "hcl":
				if !colorRegex.MatchString(value) {
					isValid = false
				}
				fieldNb++
				break
			case "ecl":
				if value != "amb" && value != "blu" && value != "brn" && value != "gry" && value != "grn" && value != "hzl" && value != "oth" {
					isValid = false
				}
				fieldNb++
				break
			case "pid":
				if !idRegex.MatchString(value) {
					isValid = false
				}
				fieldNb++
				break
			}
		}
	}
	if isValid && fieldNb == 7 {
		count++
	}

	return count
}
