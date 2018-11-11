package main

import (
	"fmt"
	"strings"
)

func main() {

	myDeck := CreateDeck()
	myDeck = ShuffleDeck(myDeck)

	myDeck, players := DealHoleCards(myDeck, 3)

	for i, player := range players {
		playerCards := strings.Join(player, ",")
		fmt.Printf("PLayer %d hole cards : %s\n", i, playerCards)

	}

	myDeck, sharedCards := DealFlop(myDeck)
	sharedCardsString := strings.Join(sharedCards, ",")
	fmt.Println("Flop :", sharedCardsString)

	myDeck, sharedCards = DealTurnOrRiver(myDeck, sharedCards)
	sharedCardsString = strings.Join(sharedCards, ",")
	fmt.Println("Turn :", sharedCardsString)

	myDeck, sharedCards = DealTurnOrRiver(myDeck, sharedCards)
	sharedCardsString = strings.Join(sharedCards, ",")
	fmt.Println("River :", sharedCardsString)
}
