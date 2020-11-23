package main

import (
	game "github.com/zshearin/poker/game"
)

func main() {
	deck := getShuffledDeck()

	game1 := deck.GetGame(5)
	game1.PrintBoardAndHands()

	game1.PrintBestFive()

}

//function implemented for testing - should create unit tests instead
func getCardsToEvaluate() game.Cards {
	cards := game.Cards{
		game.Card{Value: "J", Suit: "H"},
		game.Card{Value: "6", Suit: "D"},
		game.Card{Value: "5", Suit: "D"},
		game.Card{Value: "K", Suit: "H"},
		game.Card{Value: "A", Suit: "D"},
	}
	return cards
}

func getShuffledDeck() game.Deck {
	deck := game.GetDeck()

	//	deck.PrintOrder()
	deck.Shuffle()
	deck.Shuffle()
	deck.PrintOrder()
	return deck
}
