package main

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestDay8(t *testing.T) {
	values, err := ReadLines("inputs/day8.txt")
	if err != nil {
		panic(err)
	}
	assert.Equal(t, 1137, day8Part1(values))
	assert.Equal(t, 0, day8Part2(values))
}

func BenchmarkDay8Part1(b *testing.B) {
	values, _ := ReadLines("inputs/day8.txt")
	for i := 0; i < b.N; i++ {
		day8Part1(values)
	}
}

func BenchmarkDay8Part2(b *testing.B) {
	values, _ := ReadLines("inputs/day8.txt")
	for i := 0; i < b.N; i++ {
		day8Part2(values)
	}
}
