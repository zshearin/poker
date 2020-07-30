package main

import (
	"fmt"
)

func main() {
	deck := GetDeck()
	fmt.Println("Unshuffled deck:")
	deck.PrintOrder()

	deck.Shuffle()
	fmt.Println("\nShuffled deck:")
	deck.PrintOrder()

	hands := deck.Deal(5, 2)
	

	hands.PrintHands()
	deck.PrintRemainingCards()
}
