package poker

import (
	"errors"
	"sort"
)

//CompareTwoBestFive compares two hands
//returns 1 if first hand best, 2 if second hand best and
//0 if hands are the same (in evaluation, not necessarily identical)
//-1 if error
func CompareTwoBestFive(firstFive, secondFive Cards) (int, error) {

	if len(firstFive) != 5 || len(secondFive) != 5 {
		return -1, errors.New("bad input - both card sets need to be of length 5")
	}

	_, rank1 := firstFive.GetFiveBest(false)
	_, rank2 := secondFive.GetFiveBest(false)

	if rank1 < rank2 {
		return 1, nil
	} else if rank2 < rank1 {
		return 2, nil
	}

	//have same level of hand, need to evaluate further
	if rank1 == 1 {
		return compareStraightFlushes(firstFive, secondFive), nil
	} else if rank1 == 2 {
		compareQuads(firstFive, secondFive)
	} else if rank1 == 3 {

	}

	//check first for straight flush:

	//if one straight flush and other not, return that one as better

	//if neither, move on to next eval

	//check for quads

	//if one has and other doesn't, return one as better

	//check for full house:

	//compareFullHouses()

	return 0, nil
}

func compareStraightFlushes(firstFive Cards, secondFive Cards) int {
	sort.Sort(ByNumber(firstFive))

	sort.Sort(ByNumber(secondFive))

	return compareCard(firstFive[0], secondFive[0])
}

func compareQuads(firstFive Cards, secondFive Cards) {
	//TODO:

	//if both quads:
	//check if one's quads are higher than the other
	//if one higher, return that one as better

	//if same, check high card
	//if one has higher high card, return that one as better

	//else return -1

}

//compareCard iterates through the orderOfHighest list
//returns 1 if first card earlier in this list, 2 if second card earlier in the list
//and 0 if they appear in the same spot as this list
func compareCard(card1 Card, card2 Card) int {

	//orderOfHighest = []string{"A", "K", "Q", "J", "T", "9", "8", "7", "6", "5", "4", "3", "2"}

	var index1, index2 int

	for i, val := range orderOfHighest {
		if val == card1.Value {
			index1 = i
		}

		if val == card2.Value {
			index2 = i
		}
	}

	if index1 < index2 {
		return 1
	} else if index2 < index1 {
		return 2
	} else {
		return 0
	}

}
