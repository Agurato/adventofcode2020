package main

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestDay11(t *testing.T) {
	values, err := ReadLines("inputs/day11.txt")
	if err != nil {
		panic(err)
	}
	assert.Equal(t, 2166, day11Part1(values))
	assert.Equal(t, 1955, day11Part2(values))
}

func BenchmarkDay11Part1(b *testing.B) {
	values, _ := ReadLines("inputs/day11.txt")
	for i := 0; i < b.N; i++ {
		day11Part1(values)
	}
}

func BenchmarkDay11Part2(b *testing.B) {
	values, _ := ReadLines("inputs/day11.txt")
	for i := 0; i < b.N; i++ {
		day11Part2(values)
	}
}
