package main

import (
	"fmt"
	"strings"
)

func main() {

	myDeck := CreateDeck()
	fmt.Println("Unshuffled")
	fmt.Println(strings.Join(myDeck, ", "))

	myDeck = ShuffleDeck(myDeck)
	fmt.Println("Shuffled")
	fmt.Println(strings.Join(myDeck, ", "))

}
