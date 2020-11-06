package poker

import (
	"fmt"
	"strconv"
)

//Game is the cards for the flop, turn, river and hands dealt to each player
type Game struct {
	Hands        Hands
	Flop         Cards
	Turn         Cards
	River        Cards
	CardsForEval []Cards
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

//PrintEvalCards prints the all the cards to be evaluated for a hand
func (g *Game) PrintEvalCards() {
	for i, cardsForEval := range g.CardsForEval {
		fmt.Printf("Hand " + strconv.Itoa(i+1) + ": ")

		for _, val := range cardsForEval {
			fmt.Printf(val.Value + val.Suit + " ")
		}
		fmt.Printf("\n")
	}
}

//GetGame deals hands and returns a game object
func (d *Deck) GetGame(players int) Game {

	hands := d.Deal(players, 2)
	flop := d.GetFlop()
	turn := d.GetTurn()
	river := d.GetRiver()

	var evalCards []Cards

	for _, curCards := range hands {

		var curCardList Cards

		curCardList = append(curCardList, curCards...)
		curCardList = append(curCardList, flop...)
		curCardList = append(curCardList, turn...)
		curCardList = append(curCardList, river...)

		evalCards = append(evalCards, curCardList)
	}

	game := Game{
		Hands:        hands,
		Flop:         flop,
		Turn:         turn,
		River:        river,
		CardsForEval: evalCards,
	}
	return game
}
