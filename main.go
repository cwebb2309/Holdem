package main

import (
	"fmt"
	"strings"
)

func main() {

	myDeck := CreateDeck()
	myDeck = ShuffleDeck(myDeck)

	myDeck, myHand := GenerateHand(myDeck)

	myDeckString := strings.Join(myHand, ",")

	fmt.Println(myDeckString)

}
