package main

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestDay10(t *testing.T) {
	values, err := ReadLinesToInt("inputs/day10.txt")
	if err != nil {
		panic(err)
	}
	assert.Equal(t, 1984, day10Part1(values))
	assert.Equal(t, 3543369523456, day10Part2(values))
}

func BenchmarkDay10Part1(b *testing.B) {
	values, _ := ReadLinesToInt("inputs/day10.txt")
	for i := 0; i < b.N; i++ {
		day10Part1(values)
	}
}

func BenchmarkDay10Part2(b *testing.B) {
	values, _ := ReadLinesToInt("inputs/day10.txt")
	for i := 0; i < b.N; i++ {
		day10Part2(values)
	}
}
