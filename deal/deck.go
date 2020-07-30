package main

import (
	"math/rand"
	"time"
	"bytes"
	"fmt"
	"strconv"
)

//Card is a single element in the deck
type Card struct {
	Suit  string
	Value string
}


type Deck struct {
	Cards []Card
	NextCardIndex int
}

//GetDeck returns a sorted deck of cards
func GetDeck() Deck {
	/*
	suits := []string{"Spades", "Hearts", "Clubs", "Diamonds"}
	values := []string{"Two", "Three", "Four", "Five", "Six", "Seven", "Eight",
		"Nine", "Ten", "Jack", "Queen", "King", "Ace"}
	*/

	suits := []string{"S", "H", "C", "D"}
	values := []string{"2","3","4","5","6","7","8","9","T","J","Q","K", "A"}
	var deck Deck

	for _, suit := range suits {
		for _, value := range values {
			currentCard := Card{
				Suit:  suit,
				Value: value,
			}
			deck.Cards = append(deck.Cards, currentCard)
		}
	}
	deck.NextCardIndex = 0
	return deck

}

//Shuffle shuffles a deck of cards
func (d *Deck) Shuffle() {
	rand.Seed(time.Now().UnixNano())
	rand.Shuffle(len(d.Cards), func(i, j int) { d.Cards[i], d.Cards[j] = d.Cards[j], d.Cards[i] })
}

//PrintOrder prints the order of the deck
func (d Deck) PrintOrder() {
	var b bytes.Buffer
	for index, card := range d.Cards {
		b.WriteString(card.Value)
		b.WriteString(card.Suit)

		//b.WriteString("\n")
		if (index +1) % 13 == 0 {
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
		if (index +1 - d.NextCardIndex) % mod == 0 {
                        b.WriteString("\n")
                } else {
                        b.WriteString(", ")
                }
        }
        b.WriteString("\n")
        fmt.Printf(b.String())
}


func (d *Deck) DealHoldem(numHands int) Hands {
	return d.Deal(numHands, 2)
}

type Hand []Card
type Hands []Hand

func (h Hands) PrintHands() {
	var b bytes.Buffer
	for index, hand := range h {
		        b.WriteString("Hand ")
                        b.WriteString(strconv.Itoa(index +1))
                        b.WriteString(": ")

		for i, card := range hand {
			b.WriteString(card.Value)
			b.WriteString(card.Suit)
			if i != len(hand) - 1 {
				b.WriteString(", ")
			}
		}
		b.WriteString("\n")
	}
	b.WriteString("\n")

	fmt.Printf(b.String())
}

func (d *Deck) Deal(numHands, numCards int) Hands {

	//1 - create DS to store hands in 

	//2 - iterate through and add cards to those hands


	var hands Hands

	for i := 0; i < numHands; i++ {

		hand := []Card{}
		hands = append(hands, hand)
	}

	for i := 0; i < numCards; i++ {

		for j := 0; j < numHands; j++ {
			curHand := hands[j]
			curHand = append(curHand, d.Cards[d.NextCardIndex])

			hands[j] = curHand 
			d.NextCardIndex++
		}
	}


	return hands
}

