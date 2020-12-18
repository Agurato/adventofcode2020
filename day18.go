package main

import (
	"fmt"
	"strconv"
	"strings"
)

// Day18 Day 18
func (aoc AOC) Day18() {
	values, err := ReadLines("inputs/day18.txt")
	if err != nil {
		panic(err)
	}
	fmt.Println(day18Part1(values))
	fmt.Println(day18Part2(values))
}

func day18Part1(values []string) int {
	sum := 0

	for _, line := range values {
		sum += doOperation(strings.Split(line, " "))
	}

	return sum
}

func day18Part2(values []string) int {
	// Add parentheses to both operands of +, then do the same thing as part 1
	sum := 0

	for _, line := range values {
		lineSplit := strings.Split(line, " ")
		for i, currOp := range lineSplit {
			if currOp != "+" {
				continue
			}
			numOpen := 0
			for j := i - 1; j >= 0; j-- {
				numOpen += strings.Count(lineSplit[j], ")")
				if numOpen > 0 {
					numOpen -= strings.Count(lineSplit[j], "(")
				}
				if numOpen <= 0 {
					lineSplit[j] = "(" + lineSplit[j]
					break
				}
			}
			numOpen = 0
			for j := i + 1; j < len(lineSplit); j++ {
				numOpen += strings.Count(lineSplit[j], "(")
				if numOpen > 0 {
					numOpen -= strings.Count(lineSplit[j], ")")
				}
				if numOpen <= 0 {
					lineSplit[j] = lineSplit[j] + ")"
					break
				}
			}
		}
		sum += doOperation(lineSplit)
	}

	return sum
}

func doOperation(operation []string) int {
	total := 0
	lastOperation := "+"
	lenOperations := len(operation)
	// Remove 1st and last parenthesis
	for i := 0; i < lenOperations; i++ {
		currOp := operation[i]
		num := 0
		if currOp == "+" || currOp == "*" {
			// current operand is a + or *
			lastOperation = currOp
			continue
		} else if !strings.Contains(currOp, "(") {
			// current operand is a number
			num, _ = strconv.Atoi(currOp)
		} else {
			// find corresponding parenthesis
			corresponding := i
			numOpen := 0
			for j := i; j < lenOperations; j++ {
				numOpen += strings.Count(operation[j], "(")
				numOpen -= strings.Count(operation[j], ")")
				if numOpen == 0 {
					corresponding = j
					break
				}
			}
			// Remove parentheses
			operation[i] = operation[i][1:]
			operation[corresponding] = operation[corresponding][:len(operation[corresponding])-1]
			// Recursive operation
			num = doOperation(operation[i : corresponding+1])
			i = corresponding
		}
		switch lastOperation {
		case "+":
			total += num
		case "*":
			total *= num
		}
	}
	return total
}
