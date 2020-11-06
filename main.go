package main

import (
	"fmt"

	"github.com/zshearin/poker-go/poker"
)

func main() {
	deck := getShuffledDeck()

	game1 := deck.GetGame(5)
	game1.PrintBoard()
	game1.PrintHands()

	//	game1.PrintEvalCards()
	for _, value := range game1.CardsForEval {

		printBestFive(value)
	}

	/*
		cards := getCardsToEvaluate()

		printBestFive(cards)
	*/
}

func printBestFive(cards poker.Cards) {

	bestFiveCards := cards.Evaluate()
	for _, card := range bestFiveCards {
		val := card.Value
		if card.Value == "1" {
			val = "A"
		}
		fmt.Printf(val + card.Suit + " ")
	}
	fmt.Printf("\n\n")
}

//function implemented for testing - should create unit tests instead
func getCardsToEvaluate() poker.Cards {
	cards := poker.Cards{
		/*
			poker.Card{
				Suit:   "S",
				Value:  "5",
				Number: 5,
			},
		*/
		poker.Card{
			Suit:   "H",
			Value:  "J",
			Number: 11,
		},
		poker.Card{
			Suit:   "D",
			Value:  "6",
			Number: 6,
		},
		/*
			poker.Card{
				Suit:   "C",
				Value:  "6",
				Number: 6,
			},
		*/
		/*
			poker.Card{
				Suit:   "C",
				Value:  "4",
				Number: 4,
			},
		*/
		/*
			poker.Card{
				Suit:   "S",
				Value:  "3",
				Number: 3,
			},
			poker.Card{
				Suit:   "H",
				Value:  "3",
				Number: 3,
			},
		*/

		poker.Card{
			Suit:   "D",
			Value:  "5",
			Number: 5,
		},
		/*
			poker.Card{
				Suit:   "D",
				Value:  "A",
				Number: 14,
			},
		*/
		poker.Card{
			Suit:   "H",
			Value:  "K",
			Number: 13,
		},
		poker.Card{
			Suit:   "D",
			Value:  "A",
			Number: 14,
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

 */
