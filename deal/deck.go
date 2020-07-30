package main

import (
	"math/rand"
	"time"
	"bytes"
	"fmt"
)

//Card is a single element in the deck
type Card struct {
	Suit  string
	Value string
}

//Deck is a set of cards
type Deck []Card

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
			deck = append(deck, currentCard)
		}
	}

	return deck

}

//Shuffle shuffles a deck of cards
func (d Deck) Shuffle() {
	rand.Seed(time.Now().UnixNano())
	rand.Shuffle(len(d), func(i, j int) { d[i], d[j] = d[j], d[i] })
}

//PrintOrder prints the order of the deck
func (d Deck) PrintOrder() {
	var b bytes.Buffer
	for index, card := range d {
		b.WriteString(card.Value)
		b.WriteString(card.Suit)
		
		//b.WriteString("\n")
		if (index +1) % 13 == 0 {
			b.WriteString("\n")
		} else {
			b.WriteString(", ")
		}
	}

	fmt.Printf(b.String())
}
