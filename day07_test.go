package main

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestDay7(t *testing.T) {
	values, err := ReadLines("inputs/day7.txt")
	if err != nil {
		panic(err)
	}
	assert.Equal(t, 128, day7Part1(values))
	assert.Equal(t, 20189, day7Part2(values))
}

func BenchmarkDay7Part1(b *testing.B) {
	values, _ := ReadLines("inputs/day7.txt")
	for i := 0; i < b.N; i++ {
		day7Part1(values)
	}
}

func BenchmarkDay7Part2(b *testing.B) {
	values, _ := ReadLines("inputs/day7.txt")
	for i := 0; i < b.N; i++ {
		day7Part2(values)
	}
}
