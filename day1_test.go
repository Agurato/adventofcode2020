package main

import "testing"

func BenchmarkDay1Part1(b *testing.B) {
	for i := 0; i < b.N; i++ {
		day1Part1()
	}
}

func BenchmarkDay1Part2(b *testing.B) {
	for i := 0; i < b.N; i++ {
		day1Part2()
	}
}
