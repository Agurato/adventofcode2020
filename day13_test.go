package main

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestDay13(t *testing.T) {
	values, err := ReadLines("inputs/day13.txt")
	if err != nil {
		panic(err)
	}
	assert.Equal(t, 156, day13Part1(values))
	assert.Equal(t, 404517869995362, day13Part2(values))
}

func BenchmarkDay13Part1(b *testing.B) {
	values, _ := ReadLines("inputs/day13.txt")
	for i := 0; i < b.N; i++ {
		day13Part1(values)
	}
}

func BenchmarkDay13Part2(b *testing.B) {
	values, _ := ReadLines("inputs/day13.txt")
	for i := 0; i < b.N; i++ {
		day13Part2(values)
	}
}
