package main
/*
import (
	"fmt"
)
*/

func main() {
	deck := GetDeck()
	//deck.PrintOrder()
	deck.Shuffle()
	deck.Shuffle()
	//deck.PrintOrder()

	game1 := deck.GetGame(5)
	game1.PrintBoard()
	game1.PrintHands()

}
