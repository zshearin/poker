package main

import (
	"deck"
	"fmt"
)

func main() {
	deck := deck.GetDeck()
	fmt.Println("Unshuffled deck:")
	deck.PrintOrder()

	deck.Shuffle()
	fmt.Println("\nShuffled deck:")
	deck.PrintOrder()
}
