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
		var answers [26]bool
		everyone := strings.Split(group, "\n")
		everyoneLen := len(everyone)
		for answerI := range answers {
			count := 0
			for _, someone := range everyone {
				if strings.Contains(someone, string(rune(answerI+97))) {
					count++
				}
			}
			if count == everyoneLen {
				sum++
			}
		}
	}

	return sum
}
