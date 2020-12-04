package main

import (
	"bufio"
	"fmt"
	"os"
	"reflect"
	"strconv"
)

// AOC contains method for all challenges
type AOC struct{}

func main() {
	if len(os.Args) < 2 {
		fmt.Fprintf(os.Stderr, "You need to pass the challenge number as an argument\n")
		os.Exit(1)
	}
	dayNum, err := strconv.Atoi(os.Args[1])
	if err != nil {
		fmt.Fprintf(os.Stderr, "Argument should be a number\n")
		os.Exit(1)
	}
	if dayNum < 1 || dayNum > 25 {
		fmt.Fprintf(os.Stderr, "Challenge number should be between 1 and 25\n")
		os.Exit(1)
	}

	aoc := AOC{}
	challenge := reflect.ValueOf(aoc).MethodByName("Day" + os.Args[1])
	var emptyMethod reflect.Value
	if challenge == emptyMethod {
		fmt.Fprintf(os.Stderr, "Challenge is not implemented yet\n")
		os.Exit(1)
	}
	challenge.Call(nil)
}

// ReadLines returns slice of the file's lines
func ReadLines(path string) ([]string, error) {
	file, err := os.Open(path)
	if err != nil {
		return nil, err
	}
	defer file.Close()

	var lines []string
	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		lines = append(lines, scanner.Text())
	}
	return lines, scanner.Err()
}

// ReadLinesToInt returns slice of the file's lines converted to integers
func ReadLinesToInt(path string) ([]int, error) {
	file, err := os.Open(path)
	if err != nil {
		return nil, err
	}
	defer file.Close()

	var lines []int
	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		val, err := strconv.Atoi(scanner.Text())
		if err != nil {
			panic(err)
		}
		lines = append(lines, val)
	}
	return lines, scanner.Err()
}
