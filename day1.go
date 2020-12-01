package main

import (
	"fmt"
	"strconv"
)

// Day1 Day 1
func Day1() {
	fmt.Println(day1Part1())
	fmt.Println(day1Part2())
}

func day1Part1() int {
	lines, err := ReadLines("inputs/day1.txt")
	if err != nil {
		panic(err)
	}

	for _, line1 := range lines {
		value1, _ := strconv.Atoi(line1)
		for _, line2 := range lines {
			value2, _ := strconv.Atoi(line2)
			if value1+value2 == 2020 {
				return value1 * value2
			}
		}
	}
	return 0
}

func day1Part2() int {
	lines, err := ReadLines("inputs/day1.txt")
	if err != nil {
		panic(err)
	}

	for _, line1 := range lines {
		value1, _ := strconv.Atoi(line1)
		for _, line2 := range lines {
			value2, _ := strconv.Atoi(line2)
			for _, line3 := range lines {
				value3, _ := strconv.Atoi(line3)
				if value1+value2+value3 == 2020 {
					return value1 * value2 * value3
				}
			}
		}
	}
	return 0
}
