package main

// ReceiveDealtCard takes a hand and a deck and deals a card into the hand from the deck
func ReceiveDealtCard(deck []string, hand []string) ([]string, []string) {

	dealtCard, remainingDeck := DealCard(deck)
	newHand := append(hand, dealtCard)

	return remainingDeck, newHand
}
