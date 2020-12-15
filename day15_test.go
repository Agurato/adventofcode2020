package main

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestDay15(t *testing.T) {
	values, err := ReadLines("inputs/day15.txt")
	if err != nil {
		panic(err)
	}
	assert.Equal(t, 496, day15Part1(values))
	assert.Equal(t, 883, day15Part2(values))
}

func BenchmarkDay15Part1(b *testing.B) {
	values, _ := ReadLines("inputs/day15.txt")
	for i := 0; i < b.N; i++ {
		day15Part1(values)
	}
}

func BenchmarkDay15Part2(b *testing.B) {
	values, _ := ReadLines("inputs/day15.txt")
	for i := 0; i < b.N; i++ {
		day15Part2(values)
	}
}
