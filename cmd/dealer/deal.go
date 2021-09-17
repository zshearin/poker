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
	Board       Cards
	HandResults []HandResult
}

//HandResult is the player number and the hand that they had
//This keeps track of the relative rank between players and the type of hand
//that they have
type HandResult struct {
	Player           Player
	RelativeHandRank int
}

//GetDeal deals hands and returns a deal object
func (d *Deck) GetDeal(numPlayers int) Deal {

	hands := d.Deal(numPlayers, 2)
	flop := d.GetFlop()
	turn := d.GetTurn()
	river := d.GetRiver()

	var board Cards
	board = append(board, flop...)
	board = append(board, turn...)
	board = append(board, river...)

	var players Players

	for i, curCards := range hands {

		var curCardList Cards

		curCardList = append(curCardList, curCards...)
		curCardList = append(curCardList, board...)

		bestFiveCards, rank := GetFiveBest(curCardList)

		curPlayer := Player{
			Num:      i + 1,
			BestFive: bestFiveCards,
			HandName: getStringForRank(rank),
		}

		players = append(players, curPlayer)
	}

	sortedPlayers := sortPlayers(players)

	winnerMap := getRankOrderMap(sortedPlayers)
	handResults := formatHandResults(winnerMap)

	deal := Deal{
		Hands:       hands,
		Board:       board,
		HandResults: handResults,
	}
	return deal
}

//sort players sorts the list of players to an ordered list based on ranking
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
func formatHandResults(p map[int]Players) []HandResult {
	numRanks := len(p)

	var handResults []HandResult

	for i := 0; i < numRanks; i++ {
		curRank := i + 1
		curPlayerList := p[curRank]

		for _, curPlayer := range curPlayerList {
			handResult := HandResult{
				Player:           curPlayer,
				RelativeHandRank: curRank,
			}
			handResults = append(handResults, handResult)
		}

	}
	return handResults

}

func getRankOrderMap(p Players) map[int]Players {
	var curPList Players
	curWinner := 1
	curPList = append(curPList, p[0])
	winnerMap := map[int]Players{curWinner: curPList}

	//start at the second element
	for i := 1; i < len(p); i++ {
		//keep track of cur rank being used
		curList := winnerMap[curWinner]
		//previous
		curBestFive1 := p[i-1].BestFive
		//current
		curBestFive2 := p[i].BestFive

		winner, err := CompareTwoBestFive(curBestFive1, curBestFive2)
		check(err)

		if winner == 0 {
			curList = append(curList, p[i])
			winnerMap[curWinner] = curList
		} else {
			curWinner = curWinner + 1
			var newList Players
			newList = append(newList, p[i])
			winnerMap[curWinner] = newList
		}
	}
	return winnerMap
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
	card := d.Cards[d.NextCardIndex]
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
			hands[j] = curHand
		}
	}

	return hands
}
