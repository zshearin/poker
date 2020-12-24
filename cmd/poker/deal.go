package poker

func check(err error) {
	if err != nil {
		panic(err)
	}
}

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

//Players is a list of Player objects
type Players []Player

//Player is a player that is part of a deal
type Player struct {
	Num      int
	BestFive Cards
	HandName string
}

//Deal is the cards for the flop, turn, river and hands dealt to each player
type Deal struct {
	Hands       Hands
	Flop        Cards
	Turn        Cards
	River       Cards
	Players     Players
	HandResults []HandResult
}

//HandResult is the player number and the hand that they had
//This keeps track of the relative rank between players and the type of hand
//that they have
type HandResult struct {
	PlayerNumber int
	HandName     string
}

//PrintBoardAndHands prints the board and the hands
func (d *Deal) PrintBoardAndHands() {
	d.PrintBoard()
	d.PrintHands()
}

//GetBoard gets the board by appending the flop, turn and river
func (d *Deal) GetBoard() Cards {

	var board Cards

	board = append(board, d.Flop...)
	board = append(board, d.Turn...)
	board = append(board, d.River...)

	return board
}

//PrintBoard prints the board for a game
func (d *Deal) PrintBoard() {

	board := d.GetBoard()

	board.Print("Board", "")
}

//PrintHands prints the hands for a game
func (d *Deal) PrintHands() {
	d.Hands.Print()
}

//GetDeal deals hands and returns a deal object
func (d *Deck) GetDeal(numPlayers int) Deal {

	hands := d.Deal(numPlayers, 2)
	flop := d.GetFlop()
	turn := d.GetTurn()
	river := d.GetRiver()

	var players Players

	for i, curCards := range hands {

		var curCardList Cards

		curCardList = append(curCardList, curCards...)
		curCardList = append(curCardList, flop...)
		curCardList = append(curCardList, turn...)
		curCardList = append(curCardList, river...)

		bestFiveCards, rank := GetFiveBest(curCardList)

		curPlayer := Player{
			Num:      i + 1,
			BestFive: bestFiveCards,
			HandName: getStringForRank(rank),
		}

		players = append(players, curPlayer)
		//		bestFiveCardsList = append(bestFiveCardsList, bestFiveCards)
	}

	sortedPlayers := sortPlayers(players)

	handResults := getRankOrderList(sortedPlayers)

	deal := Deal{
		Hands:       hands,
		Flop:        flop,
		Turn:        turn,
		River:       river,
		Players:     sortedPlayers,
		HandResults: handResults,
	}
	return deal
}

func sortPlayers(pList Players) Players {

	playersList := make(Players, len(pList))
	copy(playersList, pList)

	for i := 0; i < len(playersList)-1; i++ {
		for j := 0; j < len(playersList)-i-1; j++ {

			curBestFive1 := playersList[j].BestFive
			curBestFive2 := playersList[j+1].BestFive

			winner, err := CompareTwoBestFive(curBestFive1, curBestFive2)
			check(err)

			if winner == 2 {
				playersList[j], playersList[j+1] = playersList[j+1], playersList[j]
			}
		}
	}

	return playersList

}

func getRankOrderList(p Players) []HandResult {
	/*

		type HandResult struct {
			PlayerNumber int
			HandName     string
		}
	*/
	var handResults []HandResult
	for _, val := range p {

		handResult := HandResult{PlayerNumber: val.Num, HandName: val.HandName}

		handResults = append(handResults, handResult)
	}

	return handResults
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
