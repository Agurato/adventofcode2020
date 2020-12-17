package main

import "fmt"

// Day17 Day 17
func (aoc AOC) Day17() {
	values, err := ReadLines("inputs/day17.txt")
	if err != nil {
		panic(err)
	}
	fmt.Println(day17Part1(values))
	fmt.Println(day17Part2(values))
}

func day17Part1(values []string) int {
	pocket := make(map[Coord3]bool)

	min, max := Coord3{}, Coord3{}

	for y, vY := range values {
		for x, vXY := range vY {
			if vXY == '#' {
				pocket[Coord3{x, y, 0}] = true
			}
			if x < min.x {
				min.x = x
			} else if x > max.x {
				max.x = x
			}
			if y < min.y {
				min.y = y
			} else if y > max.y {
				max.y = y
			}
		}
	}
	var oldPocket map[Coord3]bool
	for i := 0; i < 6; i++ {
		oldPocket = copyMap(pocket)
		for x := min.x - 1; x <= max.x+1; x++ {
			for y := min.y - 1; y <= max.y+1; y++ {
				for z := min.z - 1; z <= max.z+1; z++ {
					coord := Coord3{x, y, z}
					if val, ok := oldPocket[coord]; ok && val {
						// If cube is active
						active := countNeighbors(oldPocket, coord)
						if !(active == 2 || active == 3) {
							pocket[coord] = false
						}
					} else {
						// If cube is inactive
						active := countNeighbors(oldPocket, coord)
						if active == 3 {
							pocket[coord] = true

							if coord.x < min.x {
								min.x = coord.x
							} else if coord.x > max.x {
								max.x = coord.x
							}
							if coord.y < min.y {
								min.y = coord.y
							} else if coord.y > max.y {
								max.y = coord.y
							}
							if coord.z < min.z {
								min.z = coord.z
							} else if coord.z > max.z {
								max.z = coord.z
							}
						}
					}
				}
			}
		}
	}

	count := 0
	for z := min.z; z <= max.z; z++ {
		for y := min.y; y <= max.y; y++ {
			for x := min.x; x <= max.x; x++ {
				if v, ok := pocket[Coord3{x, y, z}]; ok && v {
					count++
				}
			}
		}
	}
	return count
}

func day17Part2(values []string) int {
	pocket := make(map[Coord4]bool)

	min, max := Coord4{}, Coord4{}

	for y, vY := range values {
		for x, vXY := range vY {
			if vXY == '#' {
				pocket[Coord4{x, y, 0, 0}] = true
			}
			if x < min.x {
				min.x = x
			} else if x > max.x {
				max.x = x
			}
			if y < min.y {
				min.y = y
			} else if y > max.y {
				max.y = y
			}
		}
	}
	var oldPocket map[Coord4]bool
	for i := 0; i < 6; i++ {
		oldPocket = copyMap4D(pocket)
		// fmt.Printf("Cycle %d\n", i)
		// printMap4D(pocket, min, max)
		for x := min.x - 1; x <= max.x+1; x++ {
			for y := min.y - 1; y <= max.y+1; y++ {
				for z := min.z - 1; z <= max.z+1; z++ {
					for w := min.w - 1; w <= max.w+1; w++ {
						coord := Coord4{x, y, z, w}
						if val, ok := oldPocket[coord]; ok && val {
							// If cube is active
							active := countNeighbors4D(oldPocket, coord)
							if !(active == 2 || active == 3) {
								pocket[coord] = false
							}
						} else {
							// If cube is inactive
							active := countNeighbors4D(oldPocket, coord)
							if active == 3 {
								pocket[coord] = true

								if coord.x < min.x {
									min.x = coord.x
								} else if coord.x > max.x {
									max.x = coord.x
								}
								if coord.y < min.y {
									min.y = coord.y
								} else if coord.y > max.y {
									max.y = coord.y
								}
								if coord.z < min.z {
									min.z = coord.z
								} else if coord.z > max.z {
									max.z = coord.z
								}
								if coord.w < min.w {
									min.w = coord.w
								} else if coord.w > max.w {
									max.w = coord.w
								}
							}
						}
					}
				}
			}
		}
	}

	count := 0
	for z := min.z; z <= max.z; z++ {
		for w := min.w; w <= max.w; w++ {
			for y := min.y; y <= max.y; y++ {
				for x := min.x; x <= max.x; x++ {
					if v, ok := pocket[Coord4{x, y, z, w}]; ok && v {
						count++
					}
				}
			}
		}
	}
	return count
}

// Coord3 is a 3-dimensional coordinates representation (x, y, z)
type Coord3 struct {
	x, y, z int
}

// Coord4 is a 4-dimensional coordinates representation (x, y, z, w)
type Coord4 struct {
	x, y, z, w int
}

func copyMap(oldMap map[Coord3]bool) map[Coord3]bool {
	newMap := make(map[Coord3]bool)
	for k, v := range oldMap {
		newMap[k] = v
	}
	return newMap
}

func countNeighbors(m map[Coord3]bool, pos Coord3) int {
	count := 0
	for x := pos.x - 1; x <= pos.x+1; x++ {
		for y := pos.y - 1; y <= pos.y+1; y++ {
			for z := pos.z - 1; z <= pos.z+1; z++ {
				coord := Coord3{x, y, z}
				if coord == pos {
					continue
				}
				if v, ok := m[coord]; ok && v {
					count++
				}
			}
		}
	}
	return count
}

func printMap(m map[Coord3]bool, min, max Coord3) {
	for z := min.z; z <= max.z; z++ {
		fmt.Printf("z=%2d", z)

		for y := 0; y < max.y-min.y; y++ {
			fmt.Print(" ")
		}
	}
	fmt.Println()
	for y := min.y; y <= max.y; y++ {
		for z := min.z; z <= max.z; z++ {
			for x := min.x; x <= max.x; x++ {
				coord := Coord3{x, y, z}
				if v, ok := m[coord]; ok && v {
					fmt.Print("#")
				} else {
					fmt.Print(".")
				}
			}
			fmt.Print("   ")
		}
		fmt.Println()
	}
}

func copyMap4D(oldMap map[Coord4]bool) map[Coord4]bool {
	newMap := make(map[Coord4]bool)
	for k, v := range oldMap {
		newMap[k] = v
	}
	return newMap
}

func countNeighbors4D(m map[Coord4]bool, pos Coord4) int {
	count := 0
	for x := pos.x - 1; x <= pos.x+1; x++ {
		for y := pos.y - 1; y <= pos.y+1; y++ {
			for z := pos.z - 1; z <= pos.z+1; z++ {
				for w := pos.w - 1; w <= pos.w+1; w++ {
					coord := Coord4{x, y, z, w}
					if coord == pos {
						continue
					}
					if v, ok := m[coord]; ok && v {
						count++
					}
				}
			}
		}
	}
	return count
}

func printMap4D(m map[Coord4]bool, min, max Coord4) {
	for w := min.w; w <= max.w; w++ {
		fmt.Printf("w=%2d\n", w)
		for z := min.z; z <= max.z; z++ {
			fmt.Printf("z=%2d", z)

			for y := 0; y < max.y-min.y; y++ {
				fmt.Print(" ")
			}
		}
		fmt.Println()
		for y := min.y; y <= max.y; y++ {
			for z := min.z; z <= max.z; z++ {
				for x := min.x; x <= max.x; x++ {
					coord := Coord4{x, y, z, w}
					if v, ok := m[coord]; ok && v {
						fmt.Print("#")
					} else {
						fmt.Print(".")
					}
				}
				fmt.Print("   ")
			}
			fmt.Println()
		}
		fmt.Println()
	}
}
