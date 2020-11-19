package poker

import (
	"fmt"
	"sort"
	"strconv"
)

var orderOfHighest []string

func init() {
	//Ace represented twice - A for high ace and 1 for low ace
	orderOfHighest = []string{"A", "K", "Q", "J", "T", "9", "8", "7", "6", "5", "4", "3", "2", "1"}

}

//straight flush  - 1
//quads           - 2
//full house      - 3
//flush           - 4
//straight        - 5
//three of a kind - 6
//two pair        - 7
//pair            - 8
//high card       - 9

//GetFiveBest evaluates the hand and prints out what it is
//First return param: the 5 best cards
//Second return param: the ranking of the 5 best cards.  Rankings can be found above
func (c *Cards) GetFiveBest(printValue bool) (Cards, int) {

	sort.Sort(ByNumber(*c))

	cards := *c

	//Get suit map - creates four subsets of all the cards for eval
	suitMap := getSuitToCardsMap(cards)

	//Check straight flush
	isStraightFlush, straightFlushCards := checkForStraightFlush(suitMap)
	if isStraightFlush {
		if printValue {
			fmt.Println("straight flush")
		}
		return straightFlushCards, 1
	}

	//Check quads
	isQuads, cardsFound := checkHighestCardForQuantity(cards, 4)

	if isQuads {

		if printValue {
			fmt.Println("quads")
		}

		cards.Remove(cardsFound)
		highCards := getNumHighCards(cards, 1)
		cards.Add(cardsFound)
		cardsFound = append(cardsFound, highCards...)

		return cardsFound, 2
	}

	//Check full house
	isThreeOfAKind, foundCards := checkHighestCardForQuantity(cards, 3)
	if isThreeOfAKind {

		cards.Remove(foundCards)

		isPair, foundPair := checkHighestCardForQuantity(cards, 2)

		//if not a full house, add back 3 of a kind for later eval
		cards.Add(foundCards)

		if isPair {
			if printValue {
				fmt.Println("full house")
			}
			foundCards = append(foundCards, foundPair...)
			return foundCards, 3
		}

	}

	//Check flush
	isFlush, flushCards := checkForFlush(suitMap)
	if isFlush {
		if printValue {
			fmt.Println("flush - suit is: " + flushCards[0].Suit)
		}
		return flushCards, 4

	}

	//=======================================================================
	// below this line, suit no longer matters (but would like to return winning hand with suit)

	//Check straight
	isStraight, straightCards := checkForFiveInARow(cards)
	if isStraight {
		if printValue {
			fmt.Println("straight")
		}
		return straightCards, 5
	}

	//TODO - ADD PROCESSING THREE OF A KIND
	isThreeOfAKind, foundCards = checkHighestCardForQuantity(cards, 3)
	if isThreeOfAKind {
		if printValue {
			fmt.Println("three of a kind")
		}

		cards.Remove(foundCards)

		twoHighCards := getNumHighCards(cards, 2)
		cards.Add(foundCards)

		foundCards = append(foundCards, twoHighCards...)

		return foundCards, 6

	}

	//Processes Pair and Two Pair
	isPair, foundCards := checkHighestCardForQuantity(cards, 2)
	if isPair {

		cards.Remove(foundCards)
		isTwoPair, secondPair := checkHighestCardForQuantity(cards, 2)

		if isTwoPair {
			if printValue {
				fmt.Println("two pair")
			}

			cards.Remove(secondPair)
			foundCards = append(foundCards, secondPair...)

			highCards := getNumHighCards(cards, 1)
			cards.Remove(highCards)
			foundCards = append(foundCards, highCards...)

			cards.Add(foundCards)
			return foundCards, 7
		}

		if printValue {
			fmt.Println("pair")
		}
		highCards := getNumHighCards(cards, 3)

		cards.Remove(highCards)
		foundCards = append(foundCards, highCards...)

		cards.Add(foundCards)
		return foundCards, 8

	}

	//just a high card:
	if printValue {
		fmt.Println("high cards")
	}
	highCards := getNumHighCards(cards, 5)

	return highCards, 9
}

//Add adds cards to a Cards object.  This is intended to be used for failed multistep checks (fullhouse, two pair)
func (c *Cards) Add(cardsToAdd Cards) {

	for _, cardToAdd := range cardsToAdd {
		(*c) = append((*c), cardToAdd)
	}

}

//Remove removes cards from a Cards object
func (c *Cards) Remove(cardsToRemove Cards) {

	for _, cardToRemove := range cardsToRemove {

		for i, curCard := range *c {

			if cardToRemove.Suit == curCard.Suit &&
				cardToRemove.Value == curCard.Value {
				(*c)[i] = (*c)[len(*c)-1]
				(*c)[len(*c)-1] = Card{}
				*c = (*c)[:len(*c)-1]
			}

		}
	}
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

func getNumHighCards(cards Cards, highCardsNeeded int) Cards {

	var foundCards Cards

	for i := 0; i < highCardsNeeded; i++ {
		foundHighCard, highCard := checkHighestCardForQuantity(cards, 1)

		if foundHighCard {
			foundCards = append(foundCards, highCard[0])
		} else {
			fmt.Println("oof somethings broken")
		}
		cards.Remove(highCard)

	}

	return foundCards
}

func checkHighestCardForQuantity(cards Cards, cardsNeeded int) (bool, Cards) {
	numMap := cards.getCardValues()

	var highCard string
	for _, value := range orderOfHighest {

		if numMap[value] >= cardsNeeded {
			highCard = value
			break
		}
	}

	if highCard == "" {
		return false, Cards{}
	}

	var foundCards Cards

	for _, curCard := range cards {
		if curCard.Value == highCard {

			foundCards = append(foundCards, curCard)

			if len(foundCards) == cardsNeeded {
				break
			}

		}
	}

	return true, foundCards
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

	for _, card := range c {
		var curVal int

		curVal = getNumberValue(card)
		values = append(values, curVal)
	}

	sort.Sort(sort.Reverse(sort.IntSlice(values)))

	return values, nil
}

func getNumberValue(card Card) int {
	var curVal int
	var err error
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
			fmt.Println("error converting card to number value: " + card.Value + card.Suit + ".  Error: " + err.Error())
		}
	}
	return curVal
}

//if 5 in a row, returns true and the number
//if not, returns false and 0
func checkForFiveInARow(cards Cards) (bool, Cards) {

	var fiveInARow Cards
	sort.Sort(ByNumber(cards))
	//add value for 1 for an ace (to check for low straight - ace can be high or low)
	if cards[0].Value == "A" {
		newCard := Card{
			Suit:  cards[0].Suit,
			Value: "1",
		}
		cards = append(cards, newCard)
	}

	fiveInARow = append(fiveInARow, cards[0])

	for i := 0; i < len(cards)-1; i++ {
		card1 := getNumberValue(cards[i])
		card2 := getNumberValue(cards[i+1])

		//if sequential values are 1 apart, add card to array
		if card1-1 == card2 {
			fiveInARow = append(fiveInARow, cards[i+1])

			//if values are the same, don't add but also dont reset
		} else if card1 == card2 {
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
