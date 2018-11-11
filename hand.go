package main

// ReceiveDealtCard takes a hand and a deck and deals a card into the hand from the deck
func ReceiveDealtCard(deck []string, hand []string) ([]string, []string) {

	dealtCard, remainingDeck := DealCard(deck)
	newHand := append(hand, dealtCard)

	return remainingDeck, newHand
}

// DealCards takes a deck and pulls a number of cards hand from it
func DealCards(deck []string, cardCount int) ([]string, []string) {

	hand := make([]string, 0, 9)

	for i := 0; i < cardCount; i++ {
		deck, hand = ReceiveDealtCard(deck, hand)
	}

	return deck, hand
}
