package main

// HandRank is a pseudo enum ranking poker hands
type HandRank int

const (
	straightFlush HandRank = 8
	fourOfAKind   HandRank = 7
	fullHouse     HandRank = 6
	flush         HandRank = 5
	straight      HandRank = 4
	threeOfAKind  HandRank = 3
	twoPair       HandRank = 2
	onePair       HandRank = 1
	highCard      HandRank = 0
)
