package main

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestDay17(t *testing.T) {
	values, err := ReadLines("inputs/day17.txt")
	if err != nil {
		panic(err)
	}
	assert.Equal(t, 322, day17Part1(values))
	assert.Equal(t, 2000, day17Part2(values))
}

func BenchmarkDay17Part1(b *testing.B) {
	values, _ := ReadLines("inputs/day17.txt")
	for i := 0; i < b.N; i++ {
		day17Part1(values)
	}
}

func BenchmarkDay17Part2(b *testing.B) {
	values, _ := ReadLines("inputs/day17.txt")
	for i := 0; i < b.N; i++ {
		day17Part2(values)
	}
}
