package main

import (
	"fmt"
	"strconv"
)

// Day12 Day 12
func (aoc AOC) Day12() {
	values, err := ReadLines("inputs/day12.txt")
	if err != nil {
		panic(err)
	}
	fmt.Println(day12Part1(values))
	fmt.Println(day12Part2(values))
}

func day12Part1(values []string) int {
	// Pos{X, Y} : X -> east, Y -> north
	pos := Pos{0, 0}
	dir := EAST
	for _, val := range values {
		num, _ := strconv.Atoi(val[1:])
		switch val[0] {
		case 'N':
			pos.y += num
		case 'S':
			pos.y -= num
		case 'E':
			pos.x += num
		case 'W':
			pos.x -= num
		case 'L':
			rotation := num / 90
			dir -= rotation
			dir %= 4
			for dir < 0 {
				dir += 4
			}
		case 'R':
			rotation := num / 90
			dir += rotation
			dir %= 4
			for dir < 0 {
				dir += 4
			}
		case 'F':
			switch dir {
			case EAST:
				pos.x += num
			case SOUTH:
				pos.y -= num
			case WEST:
				pos.x -= num
			case NORTH:
				pos.y += num
			}
		}
	}
	return abs(pos.x) + abs(pos.y)
}

func day12Part2(values []string) int {
	// Pos{X, Y} : X -> east, Y -> north
	pos := Pos{0, 0}
	wp := Pos{10, 1}
	for _, val := range values {
		num, _ := strconv.Atoi(val[1:])
		switch val[0] {
		case 'N':
			wp.y += num
		case 'S':
			wp.y -= num
		case 'E':
			wp.x += num
		case 'W':
			wp.x -= num
		case 'L':
			for num > 0 {
				num -= 90
				wp.x, wp.y = -wp.y, wp.x
			}
		case 'R':
			for num > 0 {
				num -= 90
				wp.x, wp.y = wp.y, -wp.x
			}
		case 'F':
			pos.x += num * wp.x
			pos.y += num * wp.y
		}
	}
	return abs(pos.x) + abs(pos.y)
}

// Directions
const (
	EAST  = iota
	SOUTH = iota
	WEST  = iota
	NORTH = iota
)

// Pos holds x, y int
type Pos struct {
	x, y int
}

func abs(x int) int {
	if x < 0 {
		return -x
	}
	return x
}
