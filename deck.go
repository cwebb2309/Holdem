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

//DealCard takes a deck of cards and returns the top card and a new deack with the top card removed
func DealCard(d []string) (string, []string) {
	dealt := d[0]
	remaining := d[1:len(d)]

	return dealt, remaining
}
