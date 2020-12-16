package main

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestDay16(t *testing.T) {
	values, err := ReadLinesSep("inputs/day16.txt", "\n\n")
	if err != nil {
		panic(err)
	}
	assert.Equal(t, 20013, day16Part1(values))
	assert.Equal(t, 5977293343129, day16Part2(values))
}

func BenchmarkDay16Part1(b *testing.B) {
	values, _ := ReadLinesSep("inputs/day16.txt", "\n\n")
	for i := 0; i < b.N; i++ {
		day16Part1(values)
	}
}

func BenchmarkDay16Part2(b *testing.B) {
	values, _ := ReadLinesSep("inputs/day16.txt", "\n\n")
	for i := 0; i < b.N; i++ {
		day16Part2(values)
	}
}
