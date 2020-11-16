package main

import (
	"github.com/zshearin/poker-go/poker"
)

func main() {
	deck := getShuffledDeck()

	game1 := deck.GetGame(5)
	game1.PrintBoardAndHands()

	game1.PrintBestFive()

}

//function implemented for testing - should create unit tests instead
func getCardsToEvaluate() poker.Cards {
	cards := poker.Cards{
		poker.Card{Value: "J", Suit: "H"},
		poker.Card{Value: "6", Suit: "D"},
		poker.Card{Value: "5", Suit: "D"},
		poker.Card{Value: "K", Suit: "H"},
		poker.Card{Value: "A", Suit: "D"},
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
