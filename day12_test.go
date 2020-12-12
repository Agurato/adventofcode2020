package main

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestDay12(t *testing.T) {
	values, err := ReadLines("inputs/day12.txt")
	if err != nil {
		panic(err)
	}
	assert.Equal(t, 1496, day12Part1(values))
	assert.Equal(t, 63843, day12Part2(values))
}

func BenchmarkDay12Part1(b *testing.B) {
	values, _ := ReadLines("inputs/day12.txt")
	for i := 0; i < b.N; i++ {
		day12Part1(values)
	}
}

func BenchmarkDay12Part2(b *testing.B) {
	values, _ := ReadLines("inputs/day12.txt")
	for i := 0; i < b.N; i++ {
		day12Part2(values)
	}
}
