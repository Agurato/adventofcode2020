package main

import "testing"

func BenchmarkDay2Part1(b *testing.B) {
	values, _ := ReadLines("inputs/day2.txt")
	for i := 0; i < b.N; i++ {
		day2Part1(values)
	}
}

func BenchmarkDay2Part2(b *testing.B) {
	values, _ := ReadLines("inputs/day2.txt")
	for i := 0; i < b.N; i++ {
		day2Part2(values)
	}
}
