package main

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestDay3(t *testing.T) {
	values, err := ReadLines("inputs/day3.txt")
	if err != nil {
		panic(err)
	}
	assert.Equal(t, day3Part1(values), 198)
	assert.Equal(t, day3Part2(values), 5140884672)
}

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
