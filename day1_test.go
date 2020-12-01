package main

import "testing"

func BenchmarkDay1Part1(b *testing.B) {
	values, _ := ReadLinesToInt("inputs/day1.txt")
	for i := 0; i < b.N; i++ {
		day1Part1(values)
	}
}

func BenchmarkDay1Part2(b *testing.B) {
	values, _ := ReadLinesToInt("inputs/day1.txt")
	for i := 0; i < b.N; i++ {
		day1Part2(values)
	}
}
