package main

// ReceiveDealtCard takes a hand and a deck and deals a card into the hand from the deck
func ReceiveDealtCard(deck []string, hand []string) ([]string, []string) {

	dealtCard, remainingDeck := DealCard(deck)
	newHand := append(hand, dealtCard)

	return remainingDeck, newHand
}

// GenerateHand takes a deck and pulls a 5 card hand from it
func GenerateHand(deck []string) ([]string, []string) {

	hand := make([]string, 0, 5)
	deck, hand = ReceiveDealtCard(deck, hand)
	deck, hand = ReceiveDealtCard(deck, hand)
	deck, hand = ReceiveDealtCard(deck, hand)
	deck, hand = ReceiveDealtCard(deck, hand)
	deck, hand = ReceiveDealtCard(deck, hand)

	return deck, hand
}
