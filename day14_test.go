package main

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestDay14(t *testing.T) {
	values, err := ReadLines("inputs/day14.txt")
	if err != nil {
		panic(err)
	}
	assert.Equal(t, 12408060320841, day14Part1(values))
	assert.Equal(t, 4466434626828, day14Part2(values))
}

func BenchmarkDay14Part1(b *testing.B) {
	values, _ := ReadLines("inputs/day14.txt")
	for i := 0; i < b.N; i++ {
		day14Part1(values)
	}
}

func BenchmarkDay14Part2(b *testing.B) {
	values, _ := ReadLines("inputs/day14.txt")
	for i := 0; i < b.N; i++ {
		day14Part2(values)
	}
}
