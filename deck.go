package main

import (
	"fmt"
	"math/rand"
	"time"
)

// CreateDeck creates an unshuffled deck of cards
func CreateDeck() []string {
	cards := make([]string, 0, 51)
	suits := []string{"D", "C", "H", "S"}
	ranks := []string{"2", "3", "4", "5", "6", "7", "8", "9", "10", "J", "Q", "K", "A"}

	p := 0
	for _, suit := range suits {
		for _, rank := range ranks {
			x := fmt.Sprintf("%s%s", rank, suit)
			cards = append(cards, x)
			p++
		}
	}
	return cards
}

//DealCard takes a deck of cards and returns the top card and a new deck with the top card removed
func DealCard(d []string) (string, []string) {
	dealt := d[0]
	remaining := d[1:len(d)]

	return dealt, remaining
}

// ShuffleDeck takes a deck and shuffles it
func ShuffleDeck(unshuffled []string) []string {

	randomizer := rand.New(rand.NewSource(time.Now().Unix()))
	noCards := len(unshuffled)
	shuffled := make([]string, noCards)

	for i := 0; i < noCards; i++ {
		randIndex := randomizer.Intn(len(unshuffled))
		shuffled[i] = unshuffled[randIndex]
		unshuffled = append(unshuffled[:randIndex], unshuffled[randIndex+1:]...)
	}

	return shuffled
}
