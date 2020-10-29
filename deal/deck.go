package deal

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

/*
type Person struct {
    Name string
    Age  int
}

// ByAge implements sort.Interface based on the Age field.
type ByAge []Person

func (a ByAge) Len() int           { return len(a) }
func (a ByAge) Less(i, j int) bool { return a[i].Age < a[j].Age }
func (a ByAge) Swap(i, j int)      { a[i], a[j] = a[j], a[i] }

func main() {
    family := []Person{
        {"Alice", 23},
        {"Eve", 2},
        {"Bob", 25},
    }
    sort.Sort(ByAge(family))
    fmt.Println(family) // [{Eve 2} {Alice 23} {Bob 25}]
}
*/

//ByNumber implemented for sort function to sort cards by number
type ByNumber Cards

func (n ByNumber) Len() int           { return len(n) }
func (n ByNumber) Less(i, j int) bool { return n[i].Number > n[j].Number }
func (n ByNumber) Swap(i, j int)      { n[i], n[j] = n[j], n[i] }

//Sort highest to lowest based on "number"
func (c *Cards) Sort() {

}
