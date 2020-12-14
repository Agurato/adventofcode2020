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
	mem := make(map[int]uint64)
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
			mem[address] = uint64(res)
		}
	}
	var sum uint64
	for _, addressValue := range mem {
		sum += addressValue
	}

	return int(sum)
}

func day14Part2(values []string) int {
	mem := make(map[uint64]int)
	mask := ""
	for _, line := range values {
		info := strings.Split(line, " = ")
		if info[0] == "mask" {
			mask = info[1]
		} else {
			num, _ := strconv.Atoi(info[1])
			addressOrig, _ := strconv.Atoi(info[0][4 : len(info[0])-1])
			addressString := fmt.Sprintf("%036b", addressOrig)
			var allAddresses []string
			allAddresses = append(allAddresses, "")
			for maskI, maskV := range mask {
				switch maskV {
				case 'X':
					lenAllAddresses := len(allAddresses)
					allAddresses = append(allAddresses, allAddresses...)
					for i := 0; i < lenAllAddresses; i++ {
						allAddresses[i] += "0"
					}
					for i := lenAllAddresses; i < 2*lenAllAddresses; i++ {
						allAddresses[i] += "1"
					}
				case '1':
					for i := range allAddresses {
						allAddresses[i] += "1"
					}
				case '0':
					for i := range allAddresses {
						allAddresses[i] += string(addressString[maskI])
					}
				}
			}
			for _, possAddress := range allAddresses {
				address, _ := strconv.ParseUint(possAddress, 2, 64)
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
