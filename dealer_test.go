package main

import (
	"strings"
	"testing"
)

func TestReceiveCard(t *testing.T) {

	testDeck := CreateDeck()
	cards := make([]string, 0, 5)

	testDeck, cards = DealSingleCard(testDeck, cards)
	testDeck, cards = DealSingleCard(testDeck, cards)
	testDeck, cards = DealSingleCard(testDeck, cards)
	testDeck, cards = DealSingleCard(testDeck, cards)
	testDeck, cards = DealSingleCard(testDeck, cards)

	expected := "2D,3D,4D,5D,6D"
	actual := strings.Join(cards, ",")

	if actual != expected {
		t.Errorf("Expected %s got %s", expected, actual)
	}

}

func TestDealFlop(t *testing.T) {

	testDeck := CreateDeck()

	testDeck, testHand := DealFlop(testDeck)

	expected := "2D,3D,4D"
	actual := strings.Join(testHand, ",")

	if actual != expected {
		t.Errorf("Expected %s got %s", expected, actual)
	}
}

func TestDealTurnOrRiver(t *testing.T) {
	testDeck := CreateDeck()

	testDeck, testHand := DealFlop(testDeck)
	testDeck, testHand = DealTurnOrRiver(testDeck, testHand)

	expected := "2D,3D,4D,5D"
	actual := strings.Join(testHand, ",")

	if actual != expected {
		t.Errorf("Expected %s got %s", expected, actual)
	}
}
