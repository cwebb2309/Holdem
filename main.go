package main

import (
	"fmt"
	"strings"
)

func main() {

	myDeck := CreateDeck()
	myDeck = ShuffleDeck(myDeck)

	myDeck, playerOne := DealHoleCards(myDeck)
	playerOneHole := strings.Join(playerOne, ",")
	fmt.Println(playerOneHole)

	myDeck, sharedCards := DealFlop(myDeck)
	sharedCardsString := strings.Join(sharedCards, ",")
	fmt.Println(sharedCardsString)

	myDeck, sharedCards = DealTurnOrRiver(myDeck, sharedCards)
	sharedCardsString = strings.Join(sharedCards, ",")
	fmt.Println(sharedCardsString)

	myDeck, sharedCards = DealTurnOrRiver(myDeck, sharedCards)
	sharedCardsString = strings.Join(sharedCards, ",")
	fmt.Println(sharedCardsString)
}
