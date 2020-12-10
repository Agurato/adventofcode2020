package main

import (
	"fmt"
	"strconv"
	"strings"
)

// Day8 Day 8
func (aoc AOC) Day8() {
	values, err := ReadLines("inputs/day8.txt")
	if err != nil {
		panic(err)
	}
	fmt.Println(day8Part1(values))
	fmt.Println(day8Part2(values))
}

func day8Part1(values []string) int {
	executed := make(map[int]bool)
	accumulator := 0
	lineI := 0
	for {
		// If we already went through this line, quit
		if _, ok := executed[lineI]; ok {
			break
		}
		executed[lineI] = true
		line := values[lineI]
		instruction := strings.Split(line, " ")
		switch instruction[0] {
		case "acc":
			value, _ := strconv.Atoi(instruction[1])
			accumulator += value
			fallthrough
		case "nop":
			lineI++
		case "jmp":
			value, _ := strconv.Atoi(instruction[1])
			lineI += value
		}
	}

	return accumulator
}

func day8Part2(values []string) int {
	oldInst := ""
	accumulator := 0
	lenValues := len(values)
	for tryI, try := range values {
		switch try[:3] {
		case "nop":
			oldInst = "nop"
			values[tryI] = "jmp" + try[3:]
		case "jmp":
			oldInst = "jmp"
			values[tryI] = "nop" + try[3:]
		default:
			continue
		}

		executed := make(map[int]bool)
		accumulator = 0
		lineI := 0
		finished := false
		for {
			// If we already went through this line, quit with failure
			if _, ok := executed[lineI]; ok {
				break
			}
			// If we reached the end of the program, quit with success
			if lineI >= lenValues {
				finished = true
				break
			}
			executed[lineI] = true
			line := values[lineI]
			instruction := strings.Split(line, " ")
			switch instruction[0] {
			case "acc":
				value, _ := strconv.Atoi(instruction[1])
				accumulator += value
				fallthrough
			case "nop":
				lineI++
			case "jmp":
				value, _ := strconv.Atoi(instruction[1])
				lineI += value
			}
		}
		if finished {
			break
		}

		values[tryI] = oldInst + try[3:]
	}

	return accumulator
}
