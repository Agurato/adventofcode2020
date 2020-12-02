package main

import (
	"fmt"
	"regexp"
	"strconv"
	"strings"
)

// Day2 Day 2
func Day2() {
	values, err := ReadLines("inputs/day2.txt")
	if err != nil {
		panic(err)
	}
	fmt.Println(day2Part1(values))
	fmt.Println(day2Part2(values))
}

func day2Part1(values []string) int {
	count := 0
	regex, _ := regexp.Compile(`(\d+)\-(\d+)\s([a-z]):\s(\w+)`)
	for _, line := range values {
		res := regex.FindStringSubmatch(line)
		min, _ := strconv.Atoi(res[1])
		max, _ := strconv.Atoi(res[2])
		letter := res[3]
		password := res[4]

		repetitions := strings.Count(password, letter)
		if repetitions >= min && repetitions <= max {
			count++
		}
	}
	return count
}

func day2Part2(values []string) int {
	count := 0
	regex, _ := regexp.Compile(`(\d+)\-(\d+)\s([a-z]):\s(\w+)`)
	for _, line := range values {
		res := regex.FindStringSubmatch(line)
		pos1, _ := strconv.Atoi(res[1])
		pos2, _ := strconv.Atoi(res[2])
		letter := res[3][0]
		password := res[4]

		if (password[pos1-1] == letter) != (password[pos2-1] == letter) {
			count++
		}
	}
	return count
}
