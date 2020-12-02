package main

import (
	"fmt"
)

// Day1 Day 1
func (aoc AOC) Day1() {
	values, err := ReadLinesToInt("inputs/day1.txt")
	if err != nil {
		panic(err)
	}
	fmt.Println(day1Part1(values))
	fmt.Println(day1Part2(values))
}

func day1Part1(values []int) int {
	valuesLen := len(values)
	for index1 := 0; index1 < valuesLen-1; index1++ {
		value1 := values[index1]
		for index2 := index1 + 1; index2 < valuesLen; index2++ {
			value2 := values[index2]
			if value1+value2 == 2020 {
				return value1 * value2
			}
		}
	}
	return 0
}

func day1Part2(values []int) int {
	valuesLen := len(values)
	for index1 := 0; index1 < valuesLen-2; index1++ {
		value1 := values[index1]
		for index2 := index1 + 1; index2 < valuesLen-1; index2++ {
			value2 := values[index2]
			add := value1 + value2
			if add >= 2020 {
				continue
			}
			for index3 := index2 + 1; index3 < valuesLen; index3++ {
				value3 := values[index3]
				if add+value3 == 2020 {
					return value1 * value2 * value3
				}
			}
		}
	}
	return 0
}
