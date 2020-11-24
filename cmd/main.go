package main

import (
	poker "github.com/zshearin/poker/cmd/poker"
)

func main() {
	game1 := shuffleAndDeal(5)
	game1.PrintBoardAndHands()
	game1.PrintBestFive()

}

func shuffleAndDeal(players int) poker.Deal {

	deck := poker.GetDeck()
	deck.Shuffle()
	deck.Shuffle()
	//deck.PrintOrder()
	return deck.GetDeal(players)
}
