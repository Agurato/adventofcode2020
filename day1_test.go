package main

import "testing"

func BenchmarkDay1Part1(b *testing.B) {
	values := []int{1721, 979, 366, 299, 675, 1456}
	for i := 0; i < b.N; i++ {
		day1Part1(values)
	}
}

func BenchmarkDay1Part2(b *testing.B) {
	values := []int{1721, 979, 366, 299, 675, 1456}
	for i := 0; i < b.N; i++ {
		day1Part2(values)
	}
}
