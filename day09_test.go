package main

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestDay9(t *testing.T) {
	values, err := ReadLinesToInt("inputs/day9.txt")
	if err != nil {
		panic(err)
	}
	assert.Equal(t, 105950735, day9Part1(values))
	assert.Equal(t, 13826915, day9Part2(values))
}

func BenchmarkDay9Part1(b *testing.B) {
	values, _ := ReadLinesToInt("inputs/day9.txt")
	for i := 0; i < b.N; i++ {
		day9Part1(values)
	}
}

func BenchmarkDay9Part2(b *testing.B) {
	values, _ := ReadLinesToInt("inputs/day9.txt")
	for i := 0; i < b.N; i++ {
		day9Part2(values)
	}
}
