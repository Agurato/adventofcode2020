package main

import (
	"fmt"
	"regexp"
	"strconv"
	"strings"

	"github.com/thoas/go-funk"
)

// Day16 Day 16
func (aoc AOC) Day16() {
	values, err := ReadLinesSep("inputs/day16.txt", "\n\n")
	if err != nil {
		panic(err)
	}
	fmt.Println(day16Part1(values))
	fmt.Println(day16Part2(values))
}

func day16Part1(values []string) int {
	var ranges []Range
	rangeRegex, _ := regexp.Compile(".*: ([0-9]+)-([0-9]+) or ([0-9]+)-([0-9]+)")
	for _, category := range strings.Split(values[0], "\n") {
		matches := rangeRegex.FindStringSubmatch(category)
		min1, _ := strconv.Atoi(matches[1])
		max1, _ := strconv.Atoi(matches[2])
		min2, _ := strconv.Atoi(matches[3])
		max2, _ := strconv.Atoi(matches[4])
		ranges = append(ranges, Range{min1, max1, min2, max2})
	}
	sum := 0
	for _, nearby := range strings.Split(values[2], "\n")[1:] {
		for _, ticketField := range strings.Split(nearby, ",") {
			ticketFieldNb, _ := strconv.Atoi(ticketField)
			isValid := false
			for _, r := range ranges {
				if r.isValid(ticketFieldNb) {
					isValid = true
				}
			}
			if !isValid {
				sum += ticketFieldNb
			}
		}
	}
	return sum
}

func day16Part2(values []string) int {
	var categories []Category
	rangeRegex, _ := regexp.Compile("(.*): ([0-9]+)-([0-9]+) or ([0-9]+)-([0-9]+)")
	for _, category := range strings.Split(values[0], "\n") {
		matches := rangeRegex.FindStringSubmatch(category)
		min1, _ := strconv.Atoi(matches[2])
		max1, _ := strconv.Atoi(matches[3])
		min2, _ := strconv.Atoi(matches[4])
		max2, _ := strconv.Atoi(matches[5])
		categories = append(categories, Category{matches[1], Range{min1, max1, min2, max2}})
	}
	var validTickets [][]int
	for _, nearby := range strings.Split(values[2], "\n")[1:] {
		isTicketValid := true
		ticket := make([]int, len(categories))
		for fieldI, ticketField := range strings.Split(nearby, ",") {
			ticketFieldNb, _ := strconv.Atoi(ticketField)
			isFieldValid := false
			for _, c := range categories {
				if c.r.isValid(ticketFieldNb) {
					isFieldValid = true
				}
			}
			if !isFieldValid {
				isTicketValid = false
				break
			}
			ticket[fieldI] = ticketFieldNb
		}
		if isTicketValid {
			validTickets = append(validTickets, ticket)
		}
	}

	categoryPosition := make([]int, len(categories))
	categoryPossible := make([][]int, len(categories))
	for catI, cat := range categories {
		for catJ := range categories {
			catCanBeHere := true
			for _, ticket := range validTickets {
				if !cat.r.isValid(ticket[catJ]) {
					catCanBeHere = false
					break
				}
			}
			if catCanBeHere {
				categoryPossible[catI] = append(categoryPossible[catI], catJ)
			}
		}
	}

	allFound := false
	for !allFound {
		for catI, catPoss := range categoryPossible {
			if len(catPoss) == 1 {
				categoryPosition[catI] = catPoss[0]
				for catJ, categoryPoss := range categoryPossible {
					if index := funk.IndexOf(categoryPoss, catPoss[0]); index != -1 {
						categoryPossible[catJ] = remove(categoryPossible[catJ], index)
					}
				}
			}
		}

		allFound = true
		for _, catPoss := range categoryPossible {
			if len(catPoss) > 0 {
				allFound = false
				break
			}
		}
	}

	myTicket := strings.Split(strings.Split(values[1], "\n")[1], ",")

	total := 1
	for i, c := range categories {
		if strings.Contains(c.name, "departure") {
			field, _ := strconv.Atoi(myTicket[categoryPosition[i]])
			total *= field
		}
	}

	return total
}

// Range double range
type Range struct {
	min1, max1, min2, max2 int
}

// Category holds range and name
type Category struct {
	name string
	r    Range
}

func (r Range) isValid(number int) bool {
	if (number >= r.min1 && number <= r.max1) || (number >= r.min2 && number <= r.max2) {
		return true
	}
	return false
}

func remove(s []int, i int) []int {
	s[len(s)-1], s[i] = s[i], s[len(s)-1]
	return s[:len(s)-1]
}
