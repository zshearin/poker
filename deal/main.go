package main

import "fmt"

/*
import (
	"fmt"
)
*/

func main() {
	deck := GetDeck()

	deck.PrintOrder()
	//deck.Shuffle()
	//deck.Shuffle()
	//deck.PrintOrder()

	cards := Cards{
		Card{
			Suit:   "H",
			Value:  "A",
			Number: 14,
		},
		Card{
			Suit:   "H",
			Value:  "2",
			Number: 2,
		},
		Card{
			Suit:   "H",
			Value:  "3",
			Number: 3,
		},
		Card{
			Suit:   "H",
			Value:  "4",
			Number: 4,
		},
		/*
			Card{
				Suit:   "H",
				Value:  "10",
				Number: 10,
			},
		*/
		Card{
			Suit:   "S",
			Value:  "5",
			Number: 5,
		},
		/*
			Card{
				Suit:   "H",
				Value:  "5",
				Number: 5,
			},
		*/
	}

	bestFiveCards := cards.Evaluate()

	for _, card := range bestFiveCards {
		val := card.Value
		if card.Value == "1" {
			val = "A"
		}
		fmt.Printf(val + card.Suit + "  ")
	}

	/*
		game1 := deck.GetGame(5)
		game1.PrintBoard()
		game1.PrintHands()
	*/
}
