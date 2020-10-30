package poker

//Game is the cards for the flop, turn, river and hands dealt to each player
type Game struct {
	Hands Hands
	Flop  Cards
	Turn  Cards
	River Cards
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

//GetGame deals hands and returns a game object
func (d *Deck) GetGame(players int) Game {

	hands := d.Deal(players, 2)
	flop := d.GetFlop()
	turn := d.GetTurn()
	river := d.GetRiver()

	game := Game{
		Hands: hands,
		Flop:  flop,
		Turn:  turn,
		River: river,
	}
	return game
}
