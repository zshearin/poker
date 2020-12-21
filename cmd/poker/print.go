package poker

import (
	"bytes"
	"fmt"
	"strconv"
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
		hand.Print("Hand "+strconv.Itoa(index+1), "")
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
			b.WriteString("  ")
		}

	}
	b.WriteString(afterStr)
	b.WriteString("\n")
	fmt.Printf(b.String())

}
