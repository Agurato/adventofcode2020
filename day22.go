package main

import (
	"fmt"
	"strconv"
	"strings"
)

// Day22 Day 22
func (aoc AOC) Day22() {
	values, err := ReadLinesSep("inputs/day22.txt", "\n\n")
	if err != nil {
		panic(err)
	}
	fmt.Println(day22Part1(values))
	fmt.Println(day22Part2(values))
}

func day22Part1(values []string) int {
	player1Split := strings.Split(values[0], "\n")
	lenPlayer1Split := len(player1Split)
	player1Deck := make([]int, lenPlayer1Split-1)
	player2Split := strings.Split(values[1], "\n")
	lenPlayer2Split := len(player2Split)
	player2Deck := make([]int, lenPlayer2Split-1)

	// Build decks
	for i := 1; i < lenPlayer1Split; i++ {
		num, _ := strconv.Atoi(player1Split[i])
		player1Deck[i-1] = num
	}
	for i := 1; i < lenPlayer2Split; i++ {
		num, _ := strconv.Atoi(player2Split[i])
		player2Deck[i-1] = num
	}

	// Game
	player1Wins := false
	for len(player1Deck) != 0 {
		card1 := player1Deck[0]
		card2 := player2Deck[0]
		if card1 > card2 {
			player1Deck = append(player1Deck, card1, card2)
		} else {
			player2Deck = append(player2Deck, card2, card1)
		}
		player1Deck = player1Deck[1:]
		player2Deck = player2Deck[1:]
		if len(player2Deck) == 0 {
			player1Wins = true
			break
		}
	}

	// Winner score
	res := 0
	if player1Wins {
		deckSize := len(player1Deck)
		for i, card := range player1Deck {
			res += (deckSize - i) * card
		}
	} else {
		deckSize := len(player2Deck)
		for i, card := range player2Deck {
			res += (deckSize - i) * card
		}
	}

	return res
}

func day22Part2(values []string) int {
	return 0
}
