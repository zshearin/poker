package main

type Game struct {
        Hands Hands
        Flop Cards
        Turn Cards
        River Cards
}

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

func (g *Game) PrintHands() {
        g.Hands.Print()
}

func (d *Deck) GetGame(players int) Game {

        hands := d.Deal(players, 2)
        flop := d.GetFlop()
        turn := d.GetTurn()
        river := d.GetRiver()

        game := Game{
                Hands: hands,
                Flop: flop,
                Turn: turn,
                River: river,
        }
        return game
}

