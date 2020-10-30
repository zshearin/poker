package poker

import (
	"math/rand"
	"time"
)

//Card is a single element in the deck
type Card struct {
	Suit   string
	Value  string
	Number int //for evaluation purposes - 2,3,4,5,6,7,8,9,10,11,12,13,14 (ace is 14)
}

//Cards is a list of Card objects.  It can be used as a deck or a hand
type Cards []Card

//Hands is a list of Cards objects.
type Hands []Cards

//Deck is a Cards object and the next card to use
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
	numbers := []int{2, 3, 4, 5, 6, 7, 8, 9, 10, 11, 12, 13, 14}
	var deck Deck

	for _, suit := range suits {
		for index, value := range values {
			currentCard := Card{
				Suit:   suit,
				Value:  value,
				Number: numbers[index],
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

//ByNumber implemented for sort function to sort cards by number
type ByNumber Cards

func (n ByNumber) Len() int           { return len(n) }
func (n ByNumber) Less(i, j int) bool { return n[i].Number > n[j].Number }
func (n ByNumber) Swap(i, j int)      { n[i], n[j] = n[j], n[i] }

//Sort highest to lowest based on "number"
func (c *Cards) Sort() {

}
