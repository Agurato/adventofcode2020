package main

import (
	"fmt"
)

// Day5 Day 5
func (aoc AOC) Day5() {
	values, err := ReadLines("inputs/day5.txt")
	if err != nil {
		panic(err)
	}
	fmt.Println(day5Part1(values))
	fmt.Println(day5Part2(values))
}

func day5Part1(values []string) int {
	var maxSeatID int
	for _, line := range values {
		seatID := getSeatID(line)
		if seatID > maxSeatID {
			maxSeatID = seatID
		}
	}

	return maxSeatID
}

func day5Part2(values []string) int {
	// minSeat := 1023
	var usedSeats [1023]bool
	for _, line := range values {
		usedSeats[getSeatID(line)] = true
	}
	start := false
	for i := 0; i < 1023; i++ {
		if !start {
			if usedSeats[i] {
				start = true
			}
		} else if !usedSeats[i] {
			return i
		}
	}

	return 0
}

func getSeatID(partitioning string) int {
	min, max := 0, 127
	for _, half := range partitioning[:7] {
		if half == 'F' {
			max = (max + min) / 2
		} else {
			min = (max+min)/2 + 1
		}
	}
	row := min
	min, max = 0, 7
	for _, half := range partitioning[7:] {
		if half == 'L' {
			max = (max + min) / 2
		} else {
			min = (max+min)/2 + 1
		}
	}
	return row*8 + min
}
