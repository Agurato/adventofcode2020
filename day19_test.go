package main

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestDay19(t *testing.T) {
	values, err := ReadLines("inputs/day19.txt")
	if err != nil {
		panic(err)
	}
	assert.Equal(t, 0, day19Part1(values))
	// assert.Equal(t, 0, day19Part2(values))
}

func BenchmarkDay19Part1(b *testing.B) {
	values, _ := ReadLines("inputs/day19.txt")
	for i := 0; i < b.N; i++ {
		day19Part1(values)
	}
}

func BenchmarkDay19Part2(b *testing.B) {
	values, _ := ReadLines("inputs/day19.txt")
	for i := 0; i < b.N; i++ {
		day19Part2(values)
	}
}
