package main

import (
	"fmt"
	"strconv"
	"strings"
)

// Day2 Day 2
func (aoc AOC) Day2() {
	values, err := ReadLines("inputs/day2.txt")
	if err != nil {
		panic(err)
	}
	fmt.Println(day2Part1(values))
	fmt.Println(day2Part2(values))
}

func split(r rune) bool {
	return r == ' ' || r == '-'
}

func day2Part1(values []string) int {
	count := 0
	for _, line := range values {
		res := strings.FieldsFunc(line, split)
		min, _ := strconv.Atoi(res[0])
		max, _ := strconv.Atoi(res[1])
		letter := res[2][:1]
		password := res[3]

		repetitions := strings.Count(password, letter)
		if repetitions >= min && repetitions <= max {
			count++
		}
	}
	return count
}

func day2Part2(values []string) int {
	count := 0
	for _, line := range values {
		res := strings.FieldsFunc(line, split)
		pos1, _ := strconv.Atoi(res[0])
		pos2, _ := strconv.Atoi(res[1])
		letter := res[2][0]
		password := res[3]

		if (password[pos1-1] == letter) != (password[pos2-1] == letter) {
			count++
		}
	}
	return count
}
