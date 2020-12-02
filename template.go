package main

import "fmt"

// Day Day
func Day() {
	values, err := ReadLines("inputs/day.txt")
	if err != nil {
		panic(err)
	}
	fmt.Println(dayPart1(values))
	fmt.Println(dayPart2(values))
}

func dayPart1(values []string) int {
	return 0
}

func dayPart2(values []string) int {
	return 0
}
