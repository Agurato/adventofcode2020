package main

import "testing"

func BenchmarkDayPart1(b *testing.B) {
	values, _ := ReadLines("inputs/day.txt")
	for i := 0; i < b.N; i++ {
		dayPart1(values)
	}
}

func BenchmarkDayPart2(b *testing.B) {
	values, _ := ReadLines("inputs/day.txt")
	for i := 0; i < b.N; i++ {
		dayPart2(values)
	}
}
