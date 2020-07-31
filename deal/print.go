package main

import (
	"strconv"
	"bytes"
	"fmt"
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


func (h Hands) Print() {
        var b bytes.Buffer
        for index, hand := range h {
                b.WriteString("Hand ")
                b.WriteString(strconv.Itoa(index + 1))
                b.WriteString(": ")

                for i, card := range hand {
                        b.WriteString(card.Value)
                        b.WriteString(card.Suit)
                        if i != len(hand)-1 {
                                b.WriteString(", ")
                        }
                }
                b.WriteString("\n")
        }
        b.WriteString("\n")

        fmt.Printf(b.String())
}


func (c Cards) Print(cardCategory string) {
        var b bytes.Buffer
        b.WriteString(cardCategory)
        b.WriteString("\n")
        for index, card := range c {
                b.WriteString(card.Value)
                b.WriteString(card.Suit)
                if index != len(c)-1 {
                        b.WriteString("    ")
                }

        }
        b.WriteString("\n")
        fmt.Printf(b.String())

}

