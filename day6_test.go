package main

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestDay6(t *testing.T) {
	values, err := ReadLinesSep("inputs/day6.txt", "\n\n")
	if err != nil {
		panic(err)
	}
	assert.Equal(t, 6590, day6Part1(values))
	assert.Equal(t, 3288, day6Part2(values))
}

func BenchmarkDay6Part1(b *testing.B) {
	values, _ := ReadLinesSep("inputs/day6.txt", "\n\n")
	for i := 0; i < b.N; i++ {
		day6Part1(values)
	}
}

func BenchmarkDay6Part2(b *testing.B) {
	values, _ := ReadLinesSep("inputs/day6.txt", "\n\n")
	for i := 0; i < b.N; i++ {
		day6Part2(values)
	}
}
