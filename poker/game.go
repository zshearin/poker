package poker

import (
	"fmt"
)

//Game is the cards for the flop, turn, river and hands dealt to each player
type Game struct {
	Hands        Hands
	Flop         Cards
	Turn         Cards
	River        Cards
	BestFiveList []Cards
}

//PrintBoardAndHands prints the board and the hands
func (g *Game) PrintBoardAndHands() {
	g.PrintBoard()
	g.PrintHands()
}

//PrintBoard prints the board for a game
func (g *Game) PrintBoard() {

	var board Cards

	for _, card := range g.Flop {
		board = append(board, card)
	}

	for _, card := range g.Turn {
		board = append(board, card)
	}

	for _, card := range g.River {
		board = append(board, card)
	}

	board.Print("Board")
}

//PrintHands prints the hands for a game
func (g *Game) PrintHands() {
	g.Hands.Print()
}

//PrintBestFive prints the all the cards to be evaluated for a hand
func (g *Game) PrintBestFive() {
	for _, val := range g.BestFiveList {
		printBestFive(val)
	}

}

func printBestFive(cards Cards) {

	bestFiveCards := cards.GetFiveBest(true)

	for _, card := range bestFiveCards {
		val := card.Value
		if card.Value == "1" {
			val = "A"
		}
		fmt.Printf(val + card.Suit + " ")
	}
	fmt.Printf("\n\n")
}

//GetGame deals hands and returns a game object
func (d *Deck) GetGame(players int) Game {

	hands := d.Deal(players, 2)
	flop := d.GetFlop()
	turn := d.GetTurn()
	river := d.GetRiver()

	var bestFiveCardsList []Cards

	for _, curCards := range hands {

		var curCardList Cards

		curCardList = append(curCardList, curCards...)
		curCardList = append(curCardList, flop...)
		curCardList = append(curCardList, turn...)
		curCardList = append(curCardList, river...)

		bestFiveCards := curCardList.GetFiveBest(false)

		bestFiveCardsList = append(bestFiveCardsList, bestFiveCards)
	}

	game := Game{
		Hands:        hands,
		Flop:         flop,
		Turn:         turn,
		River:        river,
		BestFiveList: bestFiveCardsList,
	}
	return game
}
