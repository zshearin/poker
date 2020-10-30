package main

import (
	"fmt"

	"github.com/zshearin/poker-go/poker"
)

func main() {

	//deck := getShuffledDeck()
	getShuffledDeck()

	cards := getCardsToEvaluate()

	printBestFive(cards)

}

func printBestFive(cards poker.Cards) {

	bestFiveCards := cards.Evaluate()

	for _, card := range bestFiveCards {
		val := card.Value
		if card.Value == "1" {
			val = "A"
		}
		fmt.Printf(val + card.Suit + "  ")
	}

}

//function implemented for testing - should create unit tests instead
func getCardsToEvaluate() poker.Cards {
	cards := poker.Cards{
		poker.Card{
			Suit:   "S",
			Value:  "2",
			Number: 2,
		},
		poker.Card{
			Suit:   "H",
			Value:  "2",
			Number: 2,
		},
		poker.Card{
			Suit:   "D",
			Value:  "2",
			Number: 2,
		},
		poker.Card{
			Suit:   "C",
			Value:  "2",
			Number: 2,
		},
		/*
			Card{
				Suit:   "H",
				Value:  "10",
				Number: 10,
			},
		*/
		poker.Card{
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
	return cards
}

func getShuffledDeck() poker.Deck {
	deck := poker.GetDeck()

	//	deck.PrintOrder()
	deck.Shuffle()
	deck.Shuffle()
	//	deck.PrintOrder()
	return deck
}

/*
	game1 := deck.GetGame(5)
	game1.PrintBoard()
	game1.PrintHands()
*/
