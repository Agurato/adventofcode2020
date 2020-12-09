package main

import (
	"fmt"

	"github.com/thoas/go-funk"
)

// Day9 Day 9
func (aoc AOC) Day9() {
	values, err := ReadLinesToInt("inputs/day9.txt")
	if err != nil {
		panic(err)
	}
	fmt.Println(day9Part1(values))
	fmt.Println(day9Part2(values))
}

func day9Part1(values []int) int {
	preamble := 25
	val := 0
	lenValues := len(values)
	for valI := preamble; valI < lenValues; valI++ {
		val = values[valI]
		lastN := values[valI-preamble : valI]
		found := false
		for _, part1 := range lastN {
			if inSlice(val-part1, lastN) {
				found = true
				break
			}
		}
		if !found {
			break
		}
	}
	return val
}

func day9Part2(values []int) int {
	preamble := 25
	invalid, invalidI := 0, 0
	lenValues := len(values)
	for invalidI = preamble; invalidI < lenValues; invalidI++ {
		invalid = values[invalidI]
		lastN := values[invalidI-preamble : invalidI]
		found := false
		for _, part1 := range lastN {
			if inSlice(invalid-part1, lastN) {
				found = true
				break
			}
		}
		if !found {
			break
		}
	}

	var contiguous []int
	for tryI := 0; tryI < invalidI-2; tryI++ {
		sum := values[tryI]
		tryJ := tryI + 1
		contiguous = append(contiguous, sum)
		for sum < invalid {
			try := values[tryJ]
			sum += try
			contiguous = append(contiguous, try)
			tryJ++
		}
		if sum == invalid {
			break
		}
		contiguous = nil
	}

	return funk.MinInt(contiguous).(int) + funk.MaxInt(contiguous).(int)
}

func inSlice(value int, slice []int) bool {
	for _, val := range slice {
		if val == value {
			return true
		}
	}
	return false
}
