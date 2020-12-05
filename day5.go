package main

import (
	"fmt"
	"sort"
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
	var seats []int
	for _, line := range values {
		seats = append(seats, getSeatID(line))
	}
	sort.Ints(seats)
	seatsLen := len(seats)
	prev := seats[0]
	for seatI := 1; seatI < seatsLen; seatI++ {
		seat := seats[seatI]
		if seat != prev+1 {
			return seat - 1
		}
		prev = seat
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
