package poker

import (
	"fmt"
	"sort"
	"strconv"
)

//Evaluate evaluates the hand and prints out what it is
// For now I'm going to print out what it is but later should return something
func (c *Cards) Evaluate() Cards {

	sort.Sort(ByNumber(*c))

	cards := *c

	suitMap := getSuitToCardsMap(cards)

	isStraightFlush, straightFlushCards := checkForStraightFlush(suitMap)
	if isStraightFlush {
		fmt.Println("straight flush")
		return straightFlushCards
	}

	isQuads, cardsFound := checkForQuads(cards)
	if isQuads {
		fmt.Println("quads")

		cardSet := getCardSet(cards)
		//1 remove quad value from list of cards
		newCardSet := removeCardsFromCardSet(cardSet, cardsFound)

		//2 grab higest value card out of rest
		highCard := getHighestCard(newCardSet)
		cardsFound = append(cardsFound, highCard)

		return cardsFound
	}

	//TODO - Add processing full house here (make sure to consider best full house, not just any - rare but can happen)
	//scenarios that there are multiple:
	//	1. person has 3 of one card and 3 of another (ie board: T T J J K, hand: J T -> best would be JJJTT but could have JJTTT)
	//  2. person has 3 of one card and 2 of 2 different cards (ie  T T J J K, hand: K J -> best would be JJJKK but could have JJJTT)

	isFlush, flushCards := checkForFlush(suitMap)
	if isFlush {
		fmt.Println("flush - suit is: " + flushCards[0].Suit)
		return flushCards

	}

	//=======================================================================
	// below this line, suit no longer matters (but would like to return winning hand with suit)

	isStraight, straightCards := checkForFiveInARow(cards)
	if isStraight {
		fmt.Println("is straight: ")
		return straightCards
	}

	//TODO - ADD PROCESSING THREE OF A KIND

	//TODO - ADD PROCESSING TWO PAIR

	//TODO - ADD PROCESSING PAIR

	//TODO - ADD PROCESSING HIGH CARD - MAY WANT TO START WITH THIS ONE - WILL HELP WITH THE OTHERS

	return Cards{}
}

//TODO - ADD CHECK TO SEE IF THERE ARE ANY CARDS LEFT AND RETURN AN ERROR IF SO
func getHighestCard(cardSet map[Card]bool) Card {
	var highCard Card
	curVal := 0
	for card, exists := range cardSet {
		if exists && card.Number > curVal {
			curVal = card.Number
			highCard = card
		}
	}

	return highCard
}

func removeCardsFromCardSet(cardSet map[Card]bool, cardsToRemove Cards) map[Card]bool {

	for _, curCard := range cardsToRemove {
		cardSet[curCard] = false
	}

	return cardSet

}

func printCardSet(cardSet map[Card]bool) {

	for card, value := range cardSet {
		if value == true {
			fmt.Println(card.Suit + card.Value + " ")
		}
	}
}

//create set out of all the cards

func getCardSet(cards Cards) map[Card]bool {

	cardSet := make(map[Card]bool)

	for _, currentCard := range cards {
		cardSet[currentCard] = true
	}

	return cardSet

}

func checkForFlush(suitMap map[string]Cards) (bool, Cards) {
	for _, value := range suitMap {
		if len(value) >= 5 {
			sort.Sort(ByNumber(value))
			return true, value[:5]
		}
	}
	return false, Cards{}
}

func checkForQuads(cards Cards) (bool, Cards) {
	//Next - four of a kind
	numMap := cards.getCardValues()

	var fourOfAKindCards Cards

	for key, value := range numMap {
		if value == 4 {

			for _, curCard := range cards {

				if curCard.Value == key {
					fourOfAKindCards = append(fourOfAKindCards, curCard)
				}

			}
			return true, fourOfAKindCards

		}
	}
	return false, fourOfAKindCards
}

func getSuitToCardsMap(cards Cards) map[string]Cards {

	var suitMap map[string]Cards

	suitMap = make(map[string]Cards)

	suitMap["H"] = Cards{}
	suitMap["S"] = Cards{}
	suitMap["C"] = Cards{}
	suitMap["D"] = Cards{}

	for _, card := range cards {

		curList := suitMap[card.Suit]
		curList = append(curList, card)

		suitMap[card.Suit] = curList
	}
	return suitMap
}

func checkForStraightFlush(suitMap map[string]Cards) (bool, Cards) {

	for _, cards := range suitMap {

		//sort the cards of the same suit
		sort.Sort(ByNumber(cards))

		if len(cards) > 0 {
			straightFlushFound, straightFlushCards := checkForFiveInARow(cards)
			if straightFlushFound {
				return straightFlushFound, straightFlushCards
			}

		}
	}
	return false, Cards{}
}

func (c Cards) getCardValues() map[string]int {

	numMap := map[string]int{
		"A": 0,
		"K": 0,
		"Q": 0,
		"J": 0,
		"T": 0,
		"9": 0,
		"8": 0,
		"7": 0,
		"6": 0,
		"5": 0,
		"4": 0,
		"3": 0,
		"2": 0,
	}

	for _, card := range c {
		numMap[card.Value]++
	}
	return numMap
}

//GetNumberValues returns list of ints for
func (c Cards) getNumberValues() ([]int, error) {

	var values []int
	var err error

	for _, card := range c {
		var curVal int

		if card.Value == "T" {
			curVal = 10
		} else if card.Value == "J" {
			curVal = 11
		} else if card.Value == "Q" {
			curVal = 12
		} else if card.Value == "K" {
			curVal = 13
		} else if card.Value == "A" {
			curVal = 14
		} else {
			curVal, err = strconv.Atoi(card.Value)
			if err != nil {
				return values, err
			}
		}
		values = append(values, curVal)
	}
	/*
		for _, value := range values {
			fmt.Println(strconv.Itoa(value))
		}
	*/
	sort.Sort(sort.Reverse(sort.IntSlice(values)))

	/*
		fmt.Println("they've been sorted")

		for _, value := range values {
			fmt.Println(strconv.Itoa(value))
		}
	*/

	return values, nil
}

//if 5 in a row, returns true and the number
//if not, returns false and 0
func checkForFiveInARow(cards Cards) (bool, Cards) {

	var fiveInARow Cards

	//add value for 1 for an ace (to check for low straight - ace can be high or low)
	if cards[0].Number == 14 {
		newCard := Card{
			Suit:   cards[0].Suit,
			Value:  "1",
			Number: 1,
		}
		cards = append(cards, newCard)
	}

	fiveInARow = append(fiveInARow, cards[0])

	for i := 0; i < len(cards)-1; i++ {
		//if sequential values are 1 apart, add card to array
		if cards[i].Number-1 == cards[i+1].Number {
			fiveInARow = append(fiveInARow, cards[i+1])

			//if values are the same, don't add but also dont reset
		} else if cards[i].Number == cards[i+1].Number {
			continue
		} else {
			//reset to next value
			fiveInARow = fiveInARow[:0]
			fiveInARow = append(fiveInARow, cards[i+1])
		}
		if len(fiveInARow) == 5 {
			return true, fiveInARow
		}
	}

	return false, Cards{}
}

/*
//returns true and highest card if straight present, false if not
func checkForStraight(c Cards) (bool, Card) {

	//I'M GOING TO DO AN INEFFICIENT IMPLEMENTATION FIRST

	//get numbers in the


	//

	//A - 1
	//2 - 2
	//..
	//T - 10
	//J - 11
	//Q - 12
	//K - 13
	//A - 14



}
*/

//check straight flush - return highest card if yes, none if no
//check quads - return card if yes, none if no
//check full house - return cards that are in full house - the one of three and one of two
//check flush - return all cards in flush (to potentially compare lower cards to others)
//check straight - return highest card in straight
//check three of a kind - return card val in 3 of a kind, also return two kickers
//check two pair - return two card values and kicker value
//check pair - return card value and three kickers
//check high card - return 5 highest card values

//func EvaluatePreflop()
