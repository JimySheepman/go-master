package main

import (
	"fmt"
	"os"
)

func main() {
	deck := NewPlayingCardDeck()

	fmt.Printf("--- drawing playing card ---\n")
	card := deck.RandomCard()
	fmt.Printf("drew card: %s\n", card)

	playingCard, ok := card.(*PlayingCard)
	if !ok {
		fmt.Printf("card received wasn't a playing card!")
		os.Exit(1)
	}
	fmt.Printf("card suit: %s\n", playingCard.Suit)
	fmt.Printf("card rank: %s\n", playingCard.Rank)

	deck2 := NewPlayingCardDeck2()

	fmt.Printf("--- drawing playing card ---\n")
	playingCard2 := deck2.RandomCard()
	fmt.Printf("drew card: %s\n", playingCard2)
	// Code removed
	fmt.Printf("card suit: %s\n", playingCard2.Suit)
	fmt.Printf("card rank: %s\n", playingCard2.Rank)
}
