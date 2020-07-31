package main

import (
	"math/rand"
	"time"
)

//Card is a single element in the deck
type Card struct {
	Suit  string
	Value string
}

type Cards []Card
type Hands []Cards

type Deck struct {
	Cards         []Card
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
	values := []string{"2", "3", "4", "5", "6", "7", "8", "9", "T", "J", "Q", "K", "A"}
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
//	fmt.Println("Shuffling Deck\n")
	rand.Seed(time.Now().UnixNano())
	rand.Shuffle(len(d.Cards), func(i, j int) { d.Cards[i], d.Cards[j] = d.Cards[j], d.Cards[i] })
}

func (d *Deck) DealHoldem(numHands int) Hands {
	return d.Deal(numHands, 2)
}

func (d *Deck) GetFlop() Cards {
	return d.BurnAndFlip(3)
}

func (d *Deck) GetTurn() Cards {
	return d.BurnAndFlip(1)
}

func (d *Deck) GetRiver() Cards {
	return d.BurnAndFlip(1)
}

func (d *Deck) Reset() {
	d.NextCardIndex = 0
}

func (d *Deck) GetCard() Card {
	var card Card

	card = d.Cards[d.NextCardIndex]
	d.NextCardIndex++
	return card
}

func (d *Deck) BurnAndFlip(numCards int) Cards{

	var cards Cards

	//Burn a card:
	d.GetCard()

	for i := 0; i < numCards; i++ {

		nextCard := d.GetCard()
		cards = append(cards, nextCard)
	}
	return cards

}

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
