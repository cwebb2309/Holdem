package main

import "fmt"

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
