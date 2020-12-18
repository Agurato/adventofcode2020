package main

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestDay18(t *testing.T) {
	values, err := ReadLines("inputs/day18.txt")
	if err != nil {
		panic(err)
	}
	assert.Equal(t, 86311597203806, day18Part1(values))
	assert.Equal(t, 276894767062189, day18Part2(values))
}

func BenchmarkDay18Part1(b *testing.B) {
	values, _ := ReadLines("inputs/day18.txt")
	for i := 0; i < b.N; i++ {
		day18Part1(values)
	}
}

func BenchmarkDay18Part2(b *testing.B) {
	values, _ := ReadLines("inputs/day18.txt")
	for i := 0; i < b.N; i++ {
		day18Part2(values)
	}
}
