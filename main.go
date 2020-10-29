package main

import (
	"fmt"

	"github.com/zshearin/poker-go/deal"
)

func main() {
	deck := deal.GetDeck()

	deck.PrintOrder()
	//deck.Shuffle()
	//deck.Shuffle()
	//deck.PrintOrder()

	cards := deal.Cards{
		deal.Card{
			Suit:   "H",
			Value:  "A",
			Number: 14,
		},
		deal.Card{
			Suit:   "H",
			Value:  "2",
			Number: 2,
		},
		deal.Card{
			Suit:   "H",
			Value:  "3",
			Number: 3,
		},
		deal.Card{
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
		deal.Card{
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
