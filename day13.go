package main

import (
	"fmt"
	"strconv"
	"strings"
)

// Day13 Day 13
func (aoc AOC) Day13() {
	values, err := ReadLines("inputs/day13.txt")
	if err != nil {
		panic(err)
	}
	fmt.Println(day13Part1(values))
	fmt.Println(day13Part2(values))
}

func day13Part1(values []string) int {
	earliest, _ := strconv.Atoi(values[0])
	minTime := -1
	minTimeID := 0
	for _, busID := range strings.Split(values[1], ",") {
		if busID == "x" {
			continue
		}
		loop, _ := strconv.Atoi(busID)
		nextDeparture := loop - earliest%loop
		if nextDeparture < minTime || minTime == -1 {
			minTime = nextDeparture
			minTimeID = loop
		}
	}

	return minTime * minTimeID
}

func day13Part2(values []string) int {

	buses := make(map[int]int)
	for i, busID := range strings.Split(values[1], ",") {
		if busID == "x" {
			continue
		}
		loop, _ := strconv.Atoi(busID)
		buses[loop] = i
	}

	timestamp := 0
	incr := 1
	for busID, busTime := range buses {
		for (timestamp+busTime)%busID != 0 {
			timestamp += incr
		}
		incr *= busID
	}

	return timestamp
}
