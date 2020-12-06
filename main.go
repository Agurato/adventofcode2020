package main

import (
	"fmt"
	"io/ioutil"
	"os"
	"reflect"
	"strconv"
	"strings"
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
	content, err := ioutil.ReadFile(path)
	if err != nil {
		return nil, err
	}
	return strings.Split(strings.ReplaceAll(string(content), "\r", ""), "\n"), nil
}

// ReadLinesSep returns slice of the file's lines, using a custom separator
func ReadLinesSep(path string, separator string) ([]string, error) {
	content, err := ioutil.ReadFile(path)
	if err != nil {
		return nil, err
	}
	return strings.Split(strings.ReplaceAll(string(content), "\r", ""), separator), nil
}

// ReadLinesToInt returns slice of the file's lines converted to integers
func ReadLinesToInt(path string) ([]int, error) {
	lines, err := ReadLines(path)
	if err != nil {
		return nil, err
	}
	var intLines []int
	for _, line := range lines {
		intVal, err := strconv.Atoi(line)
		if err != nil {
			return nil, err
		}
		intLines = append(intLines, intVal)
	}
	return intLines, err
}
