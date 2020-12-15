package main

import (
	"fmt"
	"strconv"
	"strings"
)

// Day15 Day 15
func (aoc AOC) Day15() {
	values, err := ReadLines("inputs/day15.txt")
	if err != nil {
		panic(err)
	}
	fmt.Println(day15Part1(values))
	fmt.Println(day15Part2(values))
}

func day15Part1(values []string) int {
	numbers := make(map[int]int)
	lastNum := 0
	inputs := strings.Split(values[0], ",")
	count := len(inputs) - 1
	for i, input := range inputs {
		num, _ := strconv.Atoi(input)
		lastNum = num
		if i == count {
			break
		}
		numbers[num] = i
	}

	for count < 2019 {
		if lastDate, ok := numbers[lastNum]; ok {
			newNum := count - lastDate
			numbers[lastNum] = count
			lastNum = newNum
		} else {
			numbers[lastNum] = count
			lastNum = 0
		}
		count++
	}

	return lastNum
}

func day15Part2(values []string) int {
	numbers := make(map[int]int)
	lastNum := 0
	inputs := strings.Split(values[0], ",")
	count := len(inputs) - 1
	for i, input := range inputs {
		num, _ := strconv.Atoi(input)
		lastNum = num
		if i == count {
			break
		}
		numbers[num] = i
	}

	for count < 29999999 {
		if lastDate, ok := numbers[lastNum]; ok {
			newNum := count - lastDate
			numbers[lastNum] = count
			lastNum = newNum
		} else {
			numbers[lastNum] = count
			lastNum = 0
		}
		count++
	}

	return lastNum
}
