package main

import (
	"strconv"

	poker "github.com/zshearin/poker/cmd/poker"
)

func check(err error) {
	if err != nil {
		panic(err)
	}

}

func main() {
	deal1 := shuffleAndDeal(4)

	deal1.PrintBoardAndHands()
	//	deal1.PrintBestFive()

	deal1.PrintBestFive()

	playersList := deal1.Players
	/*
		for i := 0; i < len(bestFive); i++ {

			curPlayer := Player{num: i + 1, bestFive: bestFive[i]}
			playersList = append(playersList, curPlayer)

		}
	*/
	//n is at most 10 - bubble sort - yikes

	for i := 0; i < len(playersList)-1; i++ {
		for j := 0; j < len(playersList)-i-1; j++ {

			curBestFive1 := playersList[j].BestFive
			curBestFive2 := playersList[j+1].BestFive

			winner, err := poker.CompareTwoBestFive(curBestFive1, curBestFive2)
			check(err)

			if winner == 2 {
				playersList[j], playersList[j+1] = playersList[j+1], playersList[j]
			}
		}
	}

	for i := 0; i < len(playersList); i++ {
		playerNum := playersList[i].Num

		playersList[i].BestFive.Print("Player " + strconv.Itoa(playerNum))
	}

}

func shuffleAndDeal(players int) poker.Deal {

	deck := poker.GetDeck()
	deck.Shuffle()
	deck.Shuffle()
	//deck.PrintOrder()
	return deck.GetDeal(players)
}
