package main

import (
	"strings"
	"testing"
)

func TestReceiveCard(t *testing.T) {

	testDeck := CreateDeck()
	cards := make([]string, 0, 5)

	testDeck, cards = ReceiveDealtCard(testDeck, cards)
	testDeck, cards = ReceiveDealtCard(testDeck, cards)
	testDeck, cards = ReceiveDealtCard(testDeck, cards)
	testDeck, cards = ReceiveDealtCard(testDeck, cards)
	testDeck, cards = ReceiveDealtCard(testDeck, cards)

	expected := "2D,3D,4D,5D,6D"
	actual := strings.Join(cards, ",")

	if actual != expected {
		t.Errorf("Expected %s got %s", expected, actual)
	}

}

func TestGenerateHand(t *testing.T) {

	testDeck := CreateDeck()

	testDeck, testHand := GenerateHand(testDeck)

	expected := "2D,3D,4D,5D,6D"
	actual := strings.Join(testHand, ",")

	if actual != expected {
		t.Errorf("Expected %s got %s", expected, actual)
	}

}
