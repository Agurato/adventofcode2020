package main

import (
	"fmt"
	"math"
	"strconv"
	"strings"
)

// Day14 Day 14
func (aoc AOC) Day14() {
	values, err := ReadLines("inputs/day14.txt")
	if err != nil {
		panic(err)
	}
	fmt.Println(day14Part1(values))
	fmt.Println(day14Part2(values))
}

func day14Part1(values []string) int {
	mem := make(map[int]int)
	mask := ""
	for _, line := range values {
		info := strings.Split(line, " = ")
		if info[0] == "mask" {
			mask = info[1]
		} else {
			num, _ := strconv.Atoi(info[1])
			numString := fmt.Sprintf("%036b", num)
			res := 0
			for maskI, maskV := range mask {
				if (maskV == 'X' && numString[maskI] == '1') || maskV == '1' {
					res += int(math.Pow(2, float64(35-maskI)))
				}
			}
			address, _ := strconv.Atoi(info[0][4 : len(info[0])-1])
			mem[address] = res
		}
	}
	sum := 0
	for _, addressValue := range mem {
		sum += addressValue
	}

	return sum
}

func day14Part2(values []string) int {
	mem := make(map[int]int)
	mask := ""
	for _, line := range values {
		info := strings.Split(line, " = ")
		if info[0] == "mask" {
			mask = info[1]
		} else {
			num, _ := strconv.Atoi(info[1])
			addressOrig, _ := strconv.Atoi(info[0][4 : len(info[0])-1])
			addressString := fmt.Sprintf("%036b", addressOrig)
			var allAddresses []int
			allAddresses = append(allAddresses, 0)
			for maskI, maskV := range mask {
				switch maskV {
				case 'X':
					lenAllAddresses := len(allAddresses)
					allAddresses = append(allAddresses, allAddresses...)
					add := int(math.Pow(2, float64(35-maskI)))
					for i := lenAllAddresses; i < 2*lenAllAddresses; i++ {
						allAddresses[i] += add
					}
				case '1':
					add := int(math.Pow(2, float64(35-maskI)))
					for i := range allAddresses {
						allAddresses[i] += add
					}
				case '0':
					if addressString[maskI] == '1' {
						add := int(math.Pow(2, float64(35-maskI)))
						for i := range allAddresses {
							allAddresses[i] += add
						}
					}
				}
			}
			for _, address := range allAddresses {
				mem[address] = num
			}
		}
	}
	sum := 0
	for _, addressValue := range mem {
		sum += addressValue
	}

	return sum
}
