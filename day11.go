package main

import "fmt"

// Day11 Day 11
func (aoc AOC) Day11() {
	values, err := ReadLines("inputs/day11.txt")
	if err != nil {
		panic(err)
	}
	fmt.Println(day11Part1(values))
	fmt.Println(day11Part2(values))
}

func day11Part1(values []string) int {
	grid := makeGrid(values)
	newGrid := copyGrid(grid)
	rowSize := len(grid[0])
	colSize := len(grid)
	firstTurn := true
	for !areGridsEqual(grid, newGrid) || firstTurn {
		grid = copyGrid(newGrid)
		for rowI, row := range grid {
			for columnI, val := range row {
				if val == '.' {
					continue
				}
				adjacent := []Coord{
					Coord{rowI - 1, columnI - 1},
					Coord{rowI - 1, columnI},
					Coord{rowI - 1, columnI + 1},
					Coord{rowI, columnI - 1},
					Coord{rowI, columnI + 1},
					Coord{rowI + 1, columnI - 1},
					Coord{rowI + 1, columnI},
					Coord{rowI + 1, columnI + 1},
				}
				occupiedAdj := 0
				for _, adj := range adjacent {
					if adj.row >= 0 && adj.row < colSize && adj.col >= 0 && adj.col < rowSize {
						if grid[adj.row][adj.col] == '#' {
							occupiedAdj++
						}
					}
				}
				if val == 'L' && occupiedAdj == 0 {
					newGrid[rowI][columnI] = '#'
				} else if val == '#' && occupiedAdj >= 4 {
					newGrid[rowI][columnI] = 'L'
				}
			}
		}
		firstTurn = false
	}

	occupied := 0
	for _, row := range grid {
		for _, val := range row {
			if val == '#' {
				occupied++
			}
		}
	}

	return occupied
}

func day11Part2(values []string) int {
	return 0
}

// Returns 2D slice: [row][column]rune
func makeGrid(values []string) [][]rune {
	grid := make([][]rune, len(values))
	rowSize := len(values[0])
	for i, val := range values {
		grid[i] = make([]rune, rowSize)
		for j, r := range val {
			grid[i][j] = r
		}
	}
	return grid
}

func copyGrid(grid [][]rune) [][]rune {
	newGrid := make([][]rune, len(grid))
	rowSize := len(grid[0])
	for i, row := range grid {
		newGrid[i] = make([]rune, rowSize)
		copy(newGrid[i], row)
	}
	return newGrid
}

func dispGrid(grid [][]rune) {
	for _, row := range grid {
		for _, val := range row {
			fmt.Printf("%c", val)
		}
		fmt.Println()
	}
	fmt.Println("------------------------")
}

func areGridsEqual(grid1, grid2 [][]rune) bool {
	for rowI, row := range grid1 {
		for colI, val := range row {
			if val != grid2[rowI][colI] {
				return false
			}
		}
	}
	return true
}

type Coord struct {
	row, col int
}
