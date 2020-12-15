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

func run(inputs []int, stop int) int {
	numbers := make([]int, stop)
	lastNum := 0
	count := len(inputs)
	for i, num := range inputs {
		lastNum = num
		if i == count {
			break
		}
		numbers[num] = i + 1
	}

	for count < stop {
		lastDate := numbers[lastNum]
		if lastDate == 0 {
			numbers[lastNum] = count
			lastNum = 0
		} else {
			newNum := count - lastDate
			numbers[lastNum] = count
			lastNum = newNum
		}
		count++
	}

	return lastNum
}

func day15Part1(values []string) int {
	split := strings.Split(values[0], ",")
	inputs := make([]int, len(split))
	for i, spl := range split {
		inputs[i], _ = strconv.Atoi(spl)
	}
	return run(inputs, 2020)
}

func day15Part2(values []string) int {
	split := strings.Split(values[0], ",")
	inputs := make([]int, len(split))
	for i, spl := range split {
		inputs[i], _ = strconv.Atoi(spl)
	}
	return run(inputs, 30000000)
}
