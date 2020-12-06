package main

import (
	"fmt"
	"strings"
)

// Day6 Day 6
func (aoc AOC) Day6() {
	values, err := ReadLinesSep("inputs/day6.txt", "\n\n")
	if err != nil {
		panic(err)
	}
	fmt.Println(day6Part1(values))
	fmt.Println(day6Part2(values))
}

func day6Part1(values []string) int {
	sum := 0
	for _, group := range values {
		group = strings.ReplaceAll(group, "\n", "")
		var answers [26]bool
		for _, question := range group {
			answers[int(question)-97] = true
		}
		for _, ans := range answers {
			if ans {
				sum++
			}
		}
	}

	return sum
}

func day6Part2(values []string) int {
	sum := 0
	for _, group := range values {
		everyone := strings.Split(group, "\n")
		everyoneLen := len(everyone)
		for _, answer := range everyone[0] {
			allYes := true
			for someoneI := 1; someoneI < everyoneLen; someoneI++ {
				someone := everyone[someoneI]
				if !strings.Contains(someone, string(rune(answer))) {
					allYes = false
					break
				}
			}
			if allYes {
				sum++
			}
		}
	}

	return sum
}
