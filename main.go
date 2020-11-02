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
	fmt.Println("=================")
	fmt.Println("Five best cards (won't print unless quads/straight flush observed):")
	for _, card := range bestFiveCards {
		val := card.Value
		if card.Value == "1" {
			val = "A"
		}
		fmt.Printf(val + card.Suit + " ")
	}
	fmt.Println("\n=================")
	fmt.Printf("\n")
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

		poker.Card{
			Suit:   "S",
			Value:  "10",
			Number: 10,
		},

		poker.Card{
			Suit:   "D",
			Value:  "J",
			Number: 11,
		},

		poker.Card{
			Suit:   "D",
			Value:  "Q",
			Number: 12,
		},
		/*
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
		*/

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
