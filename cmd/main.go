package main

import (
	poker "github.com/zshearin/poker/cmd/poker"
)

func main() {
	deal1 := shuffleAndDeal(4)

	deal1.PrintBoardAndHands()
	deal1.Players.Print()
}

func shuffleAndDeal(players int) poker.Deal {

	deck := poker.GetDeck()
	deck.Shuffle()
	deck.Shuffle()
	//deck.PrintOrder()
	return deck.GetDeal(players)
}
