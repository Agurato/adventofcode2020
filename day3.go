package main

import "fmt"

// Day3 Day 3
func (aoc AOC) Day3() {
	values, err := ReadLines("inputs/day3.txt")
	if err != nil {
		panic(err)
	}
	fmt.Println(day3Part1(values))
	fmt.Println(day3Part2(values))
}

func day3Part1(values []string) int {
	treeCount := 0
	rowSize := len(values[0])

	for rowIndex, row := range values {
		if row[(rowIndex*3)%rowSize] == '#' {
			treeCount++
		}
	}

	return treeCount
}

func day3Part2(values []string) int {
	slopes := []Slope{
		Slope{
			right: 1,
			down:  1,
		},
		Slope{
			right: 3,
			down:  1,
		},
		Slope{
			right: 5,
			down:  1,
		},
		Slope{
			right: 7,
			down:  1,
		},
		Slope{
			right: 1,
			down:  2,
		},
	}
	rowSize := len(values[0])

	for rowIndex, row := range values {
		for slopeI, slope := range slopes {
			if rowIndex%slope.down == 0 && row[(rowIndex*slope.right/slope.down)%rowSize] == '#' {
				slopes[slopeI].trees++
			}
		}
	}

	res := 1
	for _, slope := range slopes {
		res *= slope.trees
	}

	return res
}

// Slope contains the slope parameters and number of trees
type Slope struct {
	right int
	down  int
	trees int
}
