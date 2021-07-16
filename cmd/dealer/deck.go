package poker

import (
	"math/rand"
	"time"
)

//Hands is a list of Cards objects.
type Hands []Cards

//Cards is a list of Card objects.  It can be used as a deck or a hand
type Cards []Card

//Card is a single element in the deck
type Card struct {
	Suit  string
	Value string
}

//Deck is a Cards object and the next card to use
type Deck struct {
	Cards         []Card
	NextCardIndex int
}

//GetDeck returns a sorted deck of cards
func GetDeck() Deck {
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
	rand.Seed(time.Now().UnixNano())
	rand.Shuffle(len(d.Cards), func(i, j int) { d.Cards[i], d.Cards[j] = d.Cards[j], d.Cards[i] })
}

//ByNumber implemented for sort function to sort cards by number
type ByNumber Cards

func (n ByNumber) Len() int { return len(n) }
func (n ByNumber) Less(i, j int) bool {

	var index1, index2 int

	for curIndex, val := range orderOfHighest {
		if val == n[i].Value {
			index1 = curIndex
		}

		if val == n[j].Value {
			index2 = curIndex
		}
	}

	return index1 < index2
}

func (n ByNumber) Swap(i, j int) { n[i], n[j] = n[j], n[i] }

//Sort highest to lowest based on Value param
func (c *Cards) Sort() {

}
