package main

import (
	"fmt"
)

// Day1 Day 1
func Day1() {
	fmt.Println(day1Part1())
	fmt.Println(day1Part2())
}

func day1Part1() int {
	values, err := ReadLinesToInt("inputs/day1.txt")
	if err != nil {
		panic(err)
	}

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

func day1Part2() int {
	values, err := ReadLinesToInt("inputs/day1.txt")
	if err != nil {
		panic(err)
	}

	valuesLen := len(values)
	for index1 := 0; index1 < valuesLen-2; index1++ {
		value1 := values[index1]
		for index2 := index1 + 1; index2 < valuesLen-1; index2++ {
			value2 := values[index2]
			for index3 := index2 + 1; index3 < valuesLen; index3++ {
				value3 := values[index3]
				if value1+value2+value3 == 2020 {
					return value1 * value2 * value3
				}
			}
		}
	}
	return 0
}
