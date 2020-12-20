package main

import (
	"fmt"
	"strconv"
	"strings"
)

// Day19 Day 19
func (aoc AOC) Day19() {
	values, err := ReadLinesSep("inputs/day19.txt", "\n\n")
	if err != nil {
		panic(err)
	}
	fmt.Println(day19Part1(values))
	// fmt.Println(day19Part2(values))
}

func day19Part1(values []string) int {
	rules := getRules(strings.Split(values[0], "\n"))

	count := 0
	for _, textLine := range strings.Split(values[1], "\n") {
		if match(rules, 0, &textLine) && len(textLine) == 0 {
			count++
		}
	}

	return count
}

func day19Part2(values []string) int {
	rules := getRules(strings.Split(values[0], "\n"))
	rules[8] = Rule{rules: [][]int{[]int{42}, []int{42, 8}}}
	rules[11] = Rule{rules: [][]int{[]int{42, 31}, []int{42, 11, 31}}}

	count := 0
	for i, textLine := range strings.Split(values[1], "\n") {
		if match(rules, 0, &textLine) && len(textLine) == 0 {
			fmt.Println(strings.Split(values[1], "\n")[i])
			count++
		}
	}

	return count
}

// Rule rule
type Rule struct {
	char  byte
	rules [][]int
}

func getRules(lines []string) map[int]Rule {
	rules := make(map[int]Rule)
	for _, ruleLine := range lines {
		ruleLineSplit := strings.Split(ruleLine, ": ")
		ruleID, _ := strconv.Atoi(ruleLineSplit[0])
		// Text rule
		if strings.Contains(ruleLineSplit[1], "\"") {
			rules[ruleID] = Rule{char: ruleLineSplit[1][1]}
			continue
		}
		var subrules []int
		var subrulesGroup [][]int
		for _, subRule := range strings.Split(ruleLineSplit[1], " ") {
			if subRule == "|" {
				subrulesGroup = append(subrulesGroup, subrules)
				subrules = nil
			} else {
				num, _ := strconv.Atoi(subRule)
				subrules = append(subrules, num)
			}
		}
		subrulesGroup = append(subrulesGroup, subrules)
		rules[ruleID] = Rule{rules: subrulesGroup}
	}
	return rules
}

func match(rules map[int]Rule, ruleIndex int, text *string) bool {
	if len(*text) == 0 {
		return true
	}
	r := rules[ruleIndex]
	if r.char != 0 {
		// fmt.Printf("Testing rule %d: %c\n", ruleIndex, r.char)
		if (*text)[0] == r.char {
			// fmt.Printf("\tMatched %d\n", ruleIndex)
			(*text) = (*text)[1:]
			return true
		}
		return false
	}
	// fmt.Printf("Testing rule %d: %#v\n", ruleIndex, r.rules)
	matchesOr := false
	for _, subrules := range r.rules {
		matches := true
		textCopy := *text
		rulesRan := 0
		for _, subrule := range subrules {
			if !match(rules, subrule, &textCopy) {
				matches = false
				break
			}
			rulesRan++
		}
		if matches && rulesRan == len(subrules) {
			matchesOr = true
			*text = textCopy
			break
		}
	}
	// fmt.Printf("\tMatched %d\n", ruleIndex)
	return matchesOr
}
