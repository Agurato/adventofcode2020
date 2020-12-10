package main

import (
	"fmt"
	"sort"
)

// Day10 Day 10
func (aoc AOC) Day10() {
	values, err := ReadLinesToInt("inputs/day10.txt")
	if err != nil {
		panic(err)
	}
	fmt.Println(day10Part1(values))
	fmt.Println(day10Part2(values))
}

func day10Part1(values []int) int {
	sort.Ints(values)

	diff1 := 0
	diff3 := 1
	lenValues := len(values)
	if values[0] == 1 {
		diff1++
	}
	for valueI := 1; valueI < lenValues; valueI++ {
		diff := values[valueI] - values[valueI-1]
		if diff == 1 {
			diff1++
		} else if diff == 3 {
			diff3++
		}
	}

	return diff1 * diff3
}

func day10Part2(values []int) int {
	sort.Ints(values)
	lenValues := len(values)
	ways := make([]int, values[lenValues-1]+1)
	ways[0] = 1
	for _, value := range values {
		sum := 0
		if value-3 >= 0 {
			sum += ways[value-3]
		}
		if value-2 >= 0 {
			sum += ways[value-2]
		}
		if value-1 >= 0 {
			sum += ways[value-1]
		}
		ways[value] = sum
	}

	return ways[len(ways)-1]
}
