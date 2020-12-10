package main

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestDay2(t *testing.T) {
	values, err := ReadLines("inputs/day2.txt")
	if err != nil {
		panic(err)
	}
	assert.Equal(t, 517, day2Part1(values))
	assert.Equal(t, 284, day2Part2(values))
}

func BenchmarkDay2Part1(b *testing.B) {
	values, _ := ReadLines("inputs/day2.txt")
	for i := 0; i < b.N; i++ {
		day2Part1(values)
	}
}

func BenchmarkDay2Part2(b *testing.B) {
	values, _ := ReadLines("inputs/day2.txt")
	for i := 0; i < b.N; i++ {
		day2Part2(values)
	}
}
