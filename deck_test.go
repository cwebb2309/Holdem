package main

import (
	"strings"
	"testing"
)

func TestCreateDeck(t *testing.T) {

	testDeck := CreateDeck()

	actual := strings.Join(testDeck, ",")
	expected := "2D,3D,4D,5D,6D,7D,8D,9D,10D,JD,QD,KD,AD"
	expected += ",2C,3C,4C,5C,6C,7C,8C,9C,10C,JC,QC,KC,AC"
	expected += ",2H,3H,4H,5H,6H,7H,8H,9H,10H,JH,QH,KH,AH"
	expected += ",2S,3S,4S,5S,6S,7S,8S,9S,10S,JS,QS,KS,AS"

	if actual != expected {
		t.Errorf("Expected %s got %s", expected, actual)
	}
}

func TestDealCard(t *testing.T) {

	testDeck := CreateDeck()

	// test first card dealt
	expectedCard := "2D"
	expectedDeck := "3D,4D,5D,6D,7D,8D,9D,10D,JD,QD,KD,AD"
	expectedDeck += ",2C,3C,4C,5C,6C,7C,8C,9C,10C,JC,QC,KC,AC"
	expectedDeck += ",2H,3H,4H,5H,6H,7H,8H,9H,10H,JH,QH,KH,AH"
	expectedDeck += ",2S,3S,4S,5S,6S,7S,8S,9S,10S,JS,QS,KS,AS"

	actualCard, testDeck := DealCard(testDeck)
	actualDeck := strings.Join(testDeck, ",")

	if actualCard != expectedCard {
		t.Errorf("Expected %s got %s", expectedCard, actualCard)
	}

	if actualDeck != expectedDeck {
		t.Errorf("Expected %s got %s", expectedDeck, actualDeck)
	}

	// now test a second card
	expectedCard = "3D"
	expectedDeck = "4D,5D,6D,7D,8D,9D,10D,JD,QD,KD,AD"
	expectedDeck += ",2C,3C,4C,5C,6C,7C,8C,9C,10C,JC,QC,KC,AC"
	expectedDeck += ",2H,3H,4H,5H,6H,7H,8H,9H,10H,JH,QH,KH,AH"
	expectedDeck += ",2S,3S,4S,5S,6S,7S,8S,9S,10S,JS,QS,KS,AS"

	actualCard, testDeck = DealCard(testDeck)
	actualDeck = strings.Join(testDeck, ",")

	if actualCard != expectedCard {
		t.Errorf("Expected %s got %s", expectedCard, actualCard)
	}

	if actualDeck != expectedDeck {
		t.Errorf("Expected %s got %s", expectedDeck, actualDeck)
	}
}

func TestShuffleDeck(t *testing.T) {

	testDeck := CreateDeck()
	testDeck = ShuffleDeck(testDeck)

	actual := strings.Join(testDeck, ",")
	expected := "2D,3D,4D,5D,6D,7D,8D,9D,10D,JD,QD,KD,AD"
	expected += ",2C,3C,4C,5C,6C,7C,8C,9C,10C,JC,QC,KC,AC"
	expected += ",2H,3H,4H,5H,6H,7H,8H,9H,10H,JH,QH,KH,AH"
	expected += ",2S,3S,4S,5S,6S,7S,8S,9S,10S,JS,QS,KS,AS"

	if actual == expected {
		t.Errorf("Expected %s got %s", expected, actual)
	}
}
