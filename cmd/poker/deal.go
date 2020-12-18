package poker

import (
	"fmt"
	"strconv"
)

func getStringForRank(rank int) string {

	switch rank {
	case 1:
		return "straight flush"
	case 2:
		return "quads"
	case 3:
		return "full house"
	case 4:
		return "flush"
	case 5:
		return "straight"
	case 6:
		return "three of a kind"
	case 7:
		return "two pair"
	case 8:
		return "pair"
	case 9:
		return "high card"

	default:
		return "there's an error"
	}

}

type Player struct {
	Num      int
	BestFive Cards
}

//Deal is the cards for the flop, turn, river and hands dealt to each player
type Deal struct {
	Hands        Hands
	Flop         Cards
	Turn         Cards
	River        Cards
	BestFiveList []Cards //
	Players      []Player
}

//PrintBoardAndHands prints the board and the hands
func (d *Deal) PrintBoardAndHands() {
	d.PrintBoard()
	d.PrintHands()
}

//PrintBoard prints the board for a game
func (d *Deal) PrintBoard() {

	var board Cards

	for _, card := range d.Flop {
		board = append(board, card)
	}

	for _, card := range d.Turn {
		board = append(board, card)
	}

	for _, card := range d.River {
		board = append(board, card)
	}

	board.Print("Board")
}

//PrintHands prints the hands for a game
func (d *Deal) PrintHands() {
	d.Hands.Print()
}

//PrintBestFive prints the all the cards to be evaluated for a hand
func (d *Deal) PrintBestFive() {
	for i, val := range d.Players {
		fmt.Printf("hand " + strconv.Itoa(i+1) + ": ")
		printBestFive(val.BestFive)
	}
}

func printBestFive(cards Cards) {

	bestFiveCards, rank := GetFiveBest(cards)

	rankStr := getStringForRank(rank)

	fmt.Printf(rankStr + " (")
	for i, card := range bestFiveCards {
		val := card.Value
		if card.Value == "1" {
			val = "A"
		}
		fmt.Printf(val + card.Suit)
		if i != len(bestFiveCards)-1 {
			fmt.Printf(", ")
		}
	}

	fmt.Printf(")\n")
}

//GetDeal deals hands and returns a deal object
func (d *Deck) GetDeal(numPlayers int) Deal {

	hands := d.Deal(numPlayers, 2)
	flop := d.GetFlop()
	turn := d.GetTurn()
	river := d.GetRiver()

	var players []Player

	for i, curCards := range hands {

		var curCardList Cards

		curCardList = append(curCardList, curCards...)
		curCardList = append(curCardList, flop...)
		curCardList = append(curCardList, turn...)
		curCardList = append(curCardList, river...)

		bestFiveCards, _ := GetFiveBest(curCardList)

		curPlayer := Player{Num: i + 1, BestFive: bestFiveCards}
		players = append(players, curPlayer)
		//		bestFiveCardsList = append(bestFiveCardsList, bestFiveCards)
	}

	deal := Deal{
		Hands:   hands,
		Flop:    flop,
		Turn:    turn,
		River:   river,
		Players: players,
	}
	return deal
}

//DealHoldEm deals 2 cards to the number of hands passed in
//eg: if 6 is passed in, it will return 6 hands of 2
func (d *Deck) DealHoldEm(numHands int) Hands {
	return d.Deal(numHands, 2)
}

//GetFlop burns a card and returns the 3 after
func (d *Deck) GetFlop() Cards {
	return d.BurnAndFlip(3)
}

//GetTurn burns a card and returns the card after
func (d *Deck) GetTurn() Cards {
	return d.BurnAndFlip(1)
}

//GetRiver burns  a card and returns the card after
func (d *Deck) GetRiver() Cards {
	return d.BurnAndFlip(1)
}

//Reset resets the deck index to 0
func (d *Deck) Reset() {
	d.NextCardIndex = 0
}

//GetCard gets the next card in a deck and updates the index of the deck object
func (d *Deck) GetCard() Card {
	var card Card

	card = d.Cards[d.NextCardIndex]
	d.NextCardIndex++
	return card
}

//BurnAndFlip takes in a parameter that is the number of cards to get
//It first skips a card then grabs the desired number of cards
func (d *Deck) BurnAndFlip(numCards int) Cards {

	var cards Cards

	//Burn a card:
	d.GetCard()

	for i := 0; i < numCards; i++ {

		nextCard := d.GetCard()
		cards = append(cards, nextCard)
	}
	return cards

}

//Deal takes in the number of hands and the number of cards per hand
//and creates the hands from the deck, updating the next card as it goes
func (d *Deck) Deal(numHands, numCards int) Hands {

	var hands Hands

	for i := 0; i < numHands; i++ {
		hand := []Card{}
		hands = append(hands, hand)
	}

	for i := 0; i < numCards; i++ {
		for j := 0; j < numHands; j++ {
			curHand := hands[j]
			nextCard := d.GetCard()
			curHand = append(curHand, nextCard)
			//curHand = append(curHand, d.Cards[d.NextCardIndex])
			hands[j] = curHand

			//d.NextCardIndex++
		}
	}

	return hands
}
