package main

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestDay4(t *testing.T) {
	values, err := ReadLines("inputs/day4.txt")
	if err != nil {
		panic(err)
	}
	assert.Equal(t, 216, day4Part1(values))
	assert.Equal(t, 150, day4Part2(values))
}

func BenchmarkDay4Part1(b *testing.B) {
	values, _ := ReadLines("inputs/day4.txt")
	for i := 0; i < b.N; i++ {
		day4Part1(values)
	}
}

func BenchmarkDay4Part2(b *testing.B) {
	values, _ := ReadLines("inputs/day4.txt")
	for i := 0; i < b.N; i++ {
		day4Part2(values)
	}
}
