package main

import "testing"

func BenchmarkDay3Part1(b *testing.B) {
	values, _ := ReadLines("inputs/day3.txt")
	for i := 0; i < b.N; i++ {
		day3Part1(values)
	}
}

func BenchmarkDay3Part2(b *testing.B) {
	values, _ := ReadLines("inputs/day3.txt")
	for i := 0; i < b.N; i++ {
		day3Part2(values)
	}
}
