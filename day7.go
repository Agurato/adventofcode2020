package main

import (
	"fmt"
	"strconv"
	"strings"
)

// Day7 Day 7
func (aoc AOC) Day7() {
	values, err := ReadLines("inputs/day7.txt")
	if err != nil {
		panic(err)
	}
	fmt.Println(day7Part1(values))
	fmt.Println(day7Part2(values))
}

func day7Part1(values []string) int {
	bags := make(map[string]Bag)

	for _, line := range values {
		bagRules := strings.Split(line, " bag")
		var contains []string

		bagRules[1] = bagRules[1][9:] // remove " contain"
		bagRulesLen := len(bagRules) - 1
		for bagRuleI := 1; bagRuleI < bagRulesLen; bagRuleI++ {
			bagRule := bagRules[bagRuleI]
			if bagRule[1:3] == "no" {
				continue
			}
			info := strings.SplitN(bagRule, " ", 3)
			contains = append(contains, info[2])
		}

		bags[bagRules[0]] = Bag{
			color:    bagRules[0],
			contains: contains,
		}
	}

	parents := make(map[string]bool)
	bagsContaining(bags, "shiny gold", parents)
	return len(parents)
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
			color:      bagRules[0],
			contains:   contains,
			containsNb: containsNb,
		}
	}

	return bagsInside(bags, bags["shiny gold"])
}

// Bag contains other bags
type Bag struct {
	color      string
	contains   []string
	containsNb []int
}

func stringInSlice(a string, slice []string) bool {
	for _, el := range slice {
		if a == el {
			return true
		}
	}
	return false
}

func bagsContaining(bags map[string]Bag, colorSearch string, parents map[string]bool) {
	for color, bag := range bags {
		if stringInSlice(colorSearch, bag.contains) {
			parents[color] = true
			bagsContaining(bags, color, parents)
		}
	}
}

func bagsInside(bags map[string]Bag, bag Bag) int {
	total := 0
	for i, bagInside := range bag.contains {
		total += bag.containsNb[i] * (bagsInside(bags, bags[bagInside]) + 1)
	}
	return total
}
