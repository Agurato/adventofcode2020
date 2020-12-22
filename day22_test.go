package main

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestDay22(t *testing.T) {
	values, err := ReadLines("inputs/day22.txt")
	if err != nil {
		panic(err)
	}
	assert.Equal(t, 35397, day22Part1(values))
	assert.Equal(t, 0, day22Part2(values))
}

func BenchmarkDay22Part1(b *testing.B) {
	values, _ := ReadLines("inputs/day22.txt")
	for i := 0; i < b.N; i++ {
		day22Part1(values)
	}
}

func BenchmarkDay22Part2(b *testing.B) {
	values, _ := ReadLines("inputs/day22.txt")
	for i := 0; i < b.N; i++ {
		day22Part2(values)
	}
}
