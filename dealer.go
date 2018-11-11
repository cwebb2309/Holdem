package main

const maxPlayers = 8

// DealSingleCard takes a hand and a deck and deals a card into the hand from the deck
func DealSingleCard(deck []string, hand []string) ([]string, []string) {

	dealtCard, remainingDeck := DealCard(deck)
	newHand := append(hand, dealtCard)

	return remainingDeck, newHand
}

// DealCards takes a deck and pulls a number of cards hand from it
func DealCards(deck []string, hand []string, cardCount int) ([]string, []string) {

	for i := 0; i < cardCount; i++ {
		deck, hand = DealSingleCard(deck, hand)
	}

	return deck, hand
}

// DealOnePlayerHoleCards deals hole cards to one player
func DealOnePlayerHoleCards(deck []string) ([]string, []string) {

	hand := make([]string, 0, 7)
	deck, hand = DealCards(deck, hand, 2)
	return deck, hand
}

// DealFlop deals 3 cards for flop
func DealFlop(deck []string) ([]string, []string) {

	hand := make([]string, 0, 5)
	deck, hand = DealCards(deck, hand, 3)
	return deck, hand
}

// DealTurnOrRiver deals the 4th or 5th shared card
func DealTurnOrRiver(deck []string, hand []string) ([]string, []string) {
	deck, hand = DealCards(deck, hand, 1)
	return deck, hand
}

// DealHoleCards deals hole cards to players
func DealHoleCards(deck []string, playerCount int) ([]string, [][]string) {

	var hand []string
	playerHands := make([][]string, 0, maxPlayers)

	for i := 0; i < playerCount; i++ {

		deck, hand = DealOnePlayerHoleCards(deck)
		playerHands = append(playerHands, hand)
	}

	return deck, playerHands
}
