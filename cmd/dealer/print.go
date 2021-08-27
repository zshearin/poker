package poker

import (
	"bytes"
	"fmt"
	"strconv"
	"strings"
)

var (
	MIN_SPACE = 16
)

//PrintOrder prints the order of the deck
func (d Deck) PrintOrder() {
	var b bytes.Buffer
	b.WriteString("Card Order:\n")
	for index, card := range d.Cards {
		b.WriteString(card.Value)
		b.WriteString(card.Suit)

		//b.WriteString("\n")
		if (index+1)%13 == 0 {
			b.WriteString("\n")
		} else {
			b.WriteString(", ")
		}
	}
	b.WriteString("\n")
	fmt.Printf(b.String())
}

//PrintRemainingCards prints the remaining cards in the deck
func (d Deck) PrintRemainingCards() {
	var b bytes.Buffer
	b.WriteString("Remaining Cards:\n")
	remainingCards := 64 - d.NextCardIndex

	mod := remainingCards / 6
	for index, card := range d.Cards {
		if index < d.NextCardIndex {
			continue
		}
		b.WriteString(card.Value)
		b.WriteString(card.Suit)

		//make the next card the "0" placement
		if (index+1-d.NextCardIndex)%mod == 0 {
			b.WriteString("\n")
		} else {
			b.WriteString(", ")
		}
	}
	b.WriteString("\n")
	fmt.Printf(b.String())
}

//Print prints hands
func (h Hands) Print() {
	for index, hand := range h {
		playerStr := strconv.Itoa(index + 1)

		//Add space to make formatting prettier if 10 handed - 2 chars vs 1 in digit
		if len(playerStr) == 1 {
			playerStr = " " + playerStr
		}
		hand.Print("Hand "+playerStr, "")
	}
	fmt.Printf("\n")
}

//Print prints cards
func (c Cards) Print(beforeStr, afterStr string) {
	var b bytes.Buffer
	b.WriteString(beforeStr)
	b.WriteString(": ")
	for index, card := range c {
		if card.Value == "1" {
			b.WriteString("A")
		} else {
			b.WriteString(card.Value)
		}
		b.WriteString(card.Suit)
		if index != len(c)-1 {
			b.WriteString(", ")
		}

	}
	b.WriteString(afterStr)
	b.WriteString("\n")
	fmt.Printf(b.String())

}

func (c Cards) BasicPrint() {
	var b bytes.Buffer
	for index, card := range c {
		if card.Value == "1" {
			b.WriteString("A")
		} else {
			b.WriteString(card.Value)
		}
		b.WriteString(card.Suit)
		if index != len(c)-1 {
			b.WriteString(" ")
		}

	}
	fmt.Printf(b.String())
}

//Print prints players
func (p Players) Print() {
	for i := 0; i < len(p); i++ {
		playerNum := p[i].Num

		p[i].BestFive.Print("Player "+strconv.Itoa(playerNum), " ("+p[i].HandName+")")
	}

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

func (d *Deal) PrintRanksAndBestFive() {
	fmt.Println("Results:")

	fmt.Println("| Player | Rank |    Best Five   |     Hand Name   |")
	fmt.Println("----------------------------------------------------")

	for _, handResult := range d.HandResults {
		curPlayer := handResult.Player

		playerNumStr := strconv.Itoa(curPlayer.Num)
		//Add space to make formatting prettier if 10 handed - 2 chars vs 1 in digit
		if len(playerNumStr) == 1 {
			playerNumStr = playerNumStr + " "
		}

		handRankStr := strconv.Itoa(handResult.RelativeHandRank)
		if len(handRankStr) == 1 {
			handRankStr = " " + handRankStr
		}

		fmt.Printf("|   %s   |", playerNumStr)
		fmt.Printf("  %s  | ", handRankStr)
		curPlayer.BestFive.BasicPrint()
		fmt.Printf(" | ")

		numSpaces := MIN_SPACE - len(curPlayer.HandName)
		spacesBefore := strings.Repeat(" ", numSpaces/2)

		fmt.Printf(spacesBefore)
		fmt.Printf(curPlayer.HandName)
		//spaces := strings.Repeat(" ", MIN_SPACE-len(curPlayer.HandName))

		var spacesAfter string
		if numSpaces%2 == 0 {
			spacesAfter = strings.Repeat(" ", numSpaces/2)
		} else {
			spacesAfter = strings.Repeat(" ", numSpaces/2+1)

		}
		fmt.Printf("%s|\n", spacesAfter)

		/*		fmt.Printf("Player %s rank: %d. Hand name: %s.", playerNumStr, handResult.RelativeHandRank, curPlayer.HandName)

				spaces := strings.Repeat(" ", MIN_SPACE-len(curPlayer.HandName))
				fmt.Printf("%s Best Five: ", spaces)
				curPlayer.BestFive.BasicPrint()
				fmt.Printf("\n")
		*/
	}
}

//PrintBoardAndHands prints the board and the hands
func (d *Deal) PrintBoardAndHands() {
	d.PrintBoard()
	d.PrintHands()
}
