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
	assert.Equal(t, day2Part1(values), 517)
	assert.Equal(t, day2Part2(values), 284)
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
