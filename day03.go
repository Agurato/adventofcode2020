package main

import (
	"fmt"
	"sync"
)

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

	var wg sync.WaitGroup
	trees := make(chan int, len(slopes))
	valuesLen := len(values)
	for _, slope := range slopes {
		wg.Add(1)
		go func(slope Slope) {
			defer wg.Done()
			treeCount := 0
			for rowIndex := 0; rowIndex < valuesLen; rowIndex += slope.down {
				if values[rowIndex][(rowIndex*slope.right/slope.down)%rowSize] == '#' {
					treeCount++
				}
			}
			trees <- treeCount
		}(slope)
	}
	wg.Wait()
	close(trees)

	res := 1
	for hits := range trees {
		res *= hits
	}

	return res
}

// Slope contains the slope parameters
type Slope struct {
	right int
	down  int
}
