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
	adjacent := []Coord{
		Coord{-1, -1},
		Coord{-1, 0},
		Coord{-1, 1},
		Coord{0, -1},
		Coord{0, 1},
		Coord{1, -1},
		Coord{1, 0},
		Coord{1, 1},
	}
	for !areGridsEqual(grid, newGrid) || firstTurn {
		grid = copyGrid(newGrid)
		for rowI, row := range grid {
			for columnI, val := range row {
				if val == '.' {
					continue
				}
				occupiedAdj := 0
				for _, adjDiff := range adjacent {
					adj := Coord{rowI + adjDiff.row, columnI + adjDiff.col}
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
	grid := makeGrid(values)
	newGrid := copyGrid(grid)
	firstTurn := true
	directions := []Coord{
		Coord{-1, -1},
		Coord{-1, 0},
		Coord{-1, 1},
		Coord{0, -1},
		Coord{0, 1},
		Coord{1, -1},
		Coord{1, 0},
		Coord{1, 1},
	}
	for !areGridsEqual(grid, newGrid) || firstTurn {
		grid = copyGrid(newGrid)
		for rowI, row := range grid {
			for columnI, val := range row {
				if val == '.' {
					continue
				}
				occupiedDir := 0
				for _, dir := range directions {
					if !freeSeatInDirection(grid, Coord{rowI, columnI}, dir) {
						occupiedDir++
					}
				}
				if val == 'L' && occupiedDir == 0 {
					newGrid[rowI][columnI] = '#'
				} else if val == '#' && occupiedDir >= 5 {
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

func freeSeatInDirection(grid [][]rune, curr Coord, dir Coord) bool {
	nextCheck := Coord{curr.row + dir.row, curr.col + dir.col}
	rowSize := len(grid[0])
	colSize := len(grid)
	for nextCheck.row >= 0 && nextCheck.row < colSize && nextCheck.col >= 0 && nextCheck.col < rowSize {
		newSeat := grid[nextCheck.row][nextCheck.col]
		if newSeat == '#' {
			return false
		} else if newSeat == 'L' {
			return true
		}
		nextCheck.row += dir.row
		nextCheck.col += dir.col
	}

	return true
}

// Coord struct holds row and col coordinates
type Coord struct {
	row, col int
}
