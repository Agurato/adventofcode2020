package main

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestDay5(t *testing.T) {
	values, err := ReadLines("inputs/day5.txt")
	if err != nil {
		panic(err)
	}
	assert.Equal(t, 904, day5Part1(values))
	assert.Equal(t, 669, day5Part2(values))
}

func BenchmarkDay5Part1(b *testing.B) {
	values, _ := ReadLines("inputs/day5.txt")
	for i := 0; i < b.N; i++ {
		day5Part1(values)
	}
}

func BenchmarkDay5Part2(b *testing.B) {
	values, _ := ReadLines("inputs/day5.txt")
	for i := 0; i < b.N; i++ {
		day5Part2(values)
	}
}
