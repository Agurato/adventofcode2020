package main

import (
	"fmt"
	"strconv"
	"strings"
)

// Day7 Day 7
func (aoc AOC) Day7() {
	values, err := ReadLines("inputs/test.txt")
	if err != nil {
		panic(err)
	}
	fmt.Println(day7Part1(values))
	fmt.Println(day7Part2(values))
}

func day7Part1(values []string) int {
	// Thanks akkes for the idea of reversing the map!
	// It was useful to beat your perf tests

	bagsRev := make(map[string][]string) // Bag is inside which bags

	for _, line := range values {
		bagRules := strings.Split(line, " bag")
		bagRules[1] = bagRules[1][9:] // remove " contain"
		bagRulesLen := len(bagRules) - 1
		for bagRuleI := 1; bagRuleI < bagRulesLen; bagRuleI++ {
			bagRule := bagRules[bagRuleI]
			if bagRule[1:3] == "no" {
				continue
			}
			info := strings.SplitN(bagRule, " ", 3)
			if val, ok := bagsRev[info[2]]; ok {
				bagsRev[info[2]] = append(val, bagRules[0])
			} else {
				bagsRev[info[2]] = []string{bagRules[0]}
			}
		}
	}

	parents := make(map[string]bool)
	parents["shiny gold"] = true
	oldLength := -1
	for len(parents) != oldLength {
		oldLength = len(parents)
		for color := range parents {
			for _, bag := range bagsRev[color] {
				parents[bag] = true
			}
		}
	}
	return len(parents) - 1
}

func day7Part2(values []string) int {
	bags := make(map[string]Bag)

	for _, line := range values {
		bagRules := strings.Split(line, " bag")
		var contains []string
		var containsNb []int

		bagRules[1] = bagRules[1][9:] // remove " contain"
		bagRulesLen := len(bagRules) - 1
		for bagRuleI := 1; bagRuleI < bagRulesLen; bagRuleI++ {
			bagRule := bagRules[bagRuleI]
			if bagRule[1:3] == "no" {
				continue
			}
			info := strings.SplitN(bagRule, " ", 3)
			nb, _ := strconv.Atoi(info[1])
			containsNb = append(containsNb, nb)
			contains = append(contains, info[2])
		}

		bags[bagRules[0]] = Bag{
			contains:   contains,
			containsNb: containsNb,
		}
	}

	return bagsInside(bags, bags["shiny gold"])
}

// Bag contains other bags
type Bag struct {
	contains   []string
	containsNb []int
}

func bagsInside(bags map[string]Bag, bag Bag) int {
	total := 0
	for i, bagInside := range bag.contains {
		total += bag.containsNb[i] * (bagsInside(bags, bags[bagInside]) + 1)
	}
	return total
}
