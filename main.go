package main

import (
	"fmt"
	"strings"
)

func main() {

	myDeck := CreateDeck()
	myDeck = ShuffleDeck(myDeck)

	myDeck, myHand := DealCards(myDeck, 5)

	myDeckString := strings.Join(myHand, ",")

	fmt.Println(myDeckString)

}
