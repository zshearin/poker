package poker

import (
	"testing"
)

func Test_compareStraightFlushes(t *testing.T) {

	cards1 := Cards{
		Card{Suit: "C", Value: "4"},
		Card{Suit: "C", Value: "5"},
		Card{Suit: "C", Value: "6"},
		Card{Suit: "C", Value: "7"},
		Card{Suit: "C", Value: "8"},
	}

	cards2 := Cards{
		Card{Suit: "C", Value: "5"},
		Card{Suit: "C", Value: "6"},
		Card{Suit: "C", Value: "7"},
		Card{Suit: "C", Value: "8"},
		Card{Suit: "C", Value: "9"},
	}

	cards3 := Cards{
		Card{Suit: "C", Value: "A"},
		Card{Suit: "C", Value: "2"},
		Card{Suit: "C", Value: "3"},
		Card{Suit: "C", Value: "4"},
		Card{Suit: "C", Value: "5"},
	}

	type args struct {
		firstFive  Cards
		secondFive Cards
	}
	tests := []struct {
		name string
		args args
		want int
	}{
		{"test second 5 better than first 5", args{firstFive: cards1, secondFive: cards2}, 2},
		{"test same 5", args{firstFive: cards1, secondFive: cards1}, 0},
		{"test first 5 better than second 5", args{firstFive: cards2, secondFive: cards1}, 1},
		{"test low straight treated as low", args{firstFive: cards2, secondFive: cards3}, 1},

		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := compareStraightFlushes(tt.args.firstFive, tt.args.secondFive); got != tt.want {
				t.Errorf("compareStraightFlushes() = %v, want %v", got, tt.want)
			}
		})
	}
}

func Test_compareQuads(t *testing.T) {

	cards1 := Cards{
		Card{Suit: "C", Value: "4"},
		Card{Suit: "S", Value: "4"},
		Card{Suit: "D", Value: "4"},
		Card{Suit: "H", Value: "4"},
		Card{Suit: "C", Value: "A"},
	}

	cards2 := Cards{
		Card{Suit: "C", Value: "4"},
		Card{Suit: "S", Value: "4"},
		Card{Suit: "D", Value: "4"},
		Card{Suit: "H", Value: "4"},
		Card{Suit: "C", Value: "K"},
	}

	cards3 := Cards{
		Card{Suit: "C", Value: "5"},
		Card{Suit: "S", Value: "5"},
		Card{Suit: "D", Value: "5"},
		Card{Suit: "H", Value: "5"},
		Card{Suit: "C", Value: "A"},
	}

	type args struct {
		firstFive  Cards
		secondFive Cards
	}
	tests := []struct {
		name string
		args args
		want int
	}{
		{"test same quads higher card first one", args{firstFive: cards1, secondFive: cards2}, 1},
		{"test same quads higher card second one", args{firstFive: cards2, secondFive: cards1}, 2},
		{"test higher quads first one", args{firstFive: cards3, secondFive: cards2}, 1},
		{"test higher quads second one", args{firstFive: cards2, secondFive: cards3}, 2},
		{"test same exact hand", args{firstFive: cards1, secondFive: cards1}, 0},

		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := compareQuads(tt.args.firstFive, tt.args.secondFive); got != tt.want {
				t.Errorf("compareQuads() = %v, want %v", got, tt.want)
			}
		})
	}
}

func Test_compareFullHouses(t *testing.T) {

	cards1 := Cards{
		Card{Suit: "C", Value: "4"},
		Card{Suit: "S", Value: "4"},
		Card{Suit: "D", Value: "4"},
		Card{Suit: "H", Value: "5"},
		Card{Suit: "C", Value: "5"},
	}
	cards2 := Cards{
		Card{Suit: "C", Value: "5"},
		Card{Suit: "S", Value: "5"},
		Card{Suit: "D", Value: "6"},
		Card{Suit: "H", Value: "6"},
		Card{Suit: "C", Value: "6"},
	}
	cards3 := Cards{
		Card{Suit: "C", Value: "8"},
		Card{Suit: "S", Value: "8"},
		Card{Suit: "D", Value: "6"},
		Card{Suit: "H", Value: "6"},
		Card{Suit: "C", Value: "6"},
	}

	type args struct {
		firstFive  Cards
		secondFive Cards
	}
	tests := []struct {
		name string
		args args
		want int
	}{
		{"test same exact hand ", args{firstFive: cards1, secondFive: cards1}, 0},
		{"test same pair, higher trips on first", args{firstFive: cards2, secondFive: cards1}, 1},
		{"test same pair, higher trips on second", args{firstFive: cards1, secondFive: cards2}, 2},
		{"test same trips, higher pair on first", args{firstFive: cards3, secondFive: cards2}, 1},
		{"test same trips, higher pair on second", args{firstFive: cards2, secondFive: cards3}, 2},
		{"test higher trips and higher pair on first", args{firstFive: cards3, secondFive: cards1}, 1},
		{"test higher trips and higher pair on second", args{firstFive: cards1, secondFive: cards3}, 2},

		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := compareFullHouses(tt.args.firstFive, tt.args.secondFive); got != tt.want {
				t.Errorf("compareFullHouses() = %v, want %v", got, tt.want)
			}
		})
	}
}

func Test_compareFlushes(t *testing.T) {
	cards1 := Cards{
		Card{Suit: "C", Value: "4"},
		Card{Suit: "C", Value: "5"},
		Card{Suit: "C", Value: "A"},
		Card{Suit: "C", Value: "7"},
		Card{Suit: "C", Value: "8"},
	}

	cards2 := Cards{
		Card{Suit: "C", Value: "J"},
		Card{Suit: "C", Value: "6"},
		Card{Suit: "C", Value: "7"},
		Card{Suit: "C", Value: "8"},
		Card{Suit: "C", Value: "9"},
	}

	cards3 := Cards{
		Card{Suit: "C", Value: "A"},
		Card{Suit: "C", Value: "K"},
		Card{Suit: "C", Value: "9"},
		Card{Suit: "C", Value: "5"},
		Card{Suit: "C", Value: "6"},
	}
	cards4 := Cards{
		Card{Suit: "C", Value: "A"},
		Card{Suit: "C", Value: "K"},
		Card{Suit: "C", Value: "9"},
		Card{Suit: "C", Value: "4"},
		Card{Suit: "C", Value: "6"},
	}
	type args struct {
		firstFive  Cards
		secondFive Cards
	}
	tests := []struct {
		name string
		args args
		want int
	}{
		{"test first card in flush is higher first one", args{firstFive: cards1, secondFive: cards2}, 1},
		{"test first card in flush is higher second one", args{firstFive: cards2, secondFive: cards1}, 2},

		{"test second card in flush is higher first one", args{firstFive: cards3, secondFive: cards1}, 1},
		{"test second card in flush is higher second one", args{firstFive: cards1, secondFive: cards3}, 2},

		{"test last card in first flush is higher - rest same", args{firstFive: cards3, secondFive: cards4}, 1},
		{"test last card in second flush is higher - rest same", args{firstFive: cards4, secondFive: cards3}, 2},

		{"test same flush 1", args{firstFive: cards1, secondFive: cards1}, 0},
		{"test same flush 2", args{firstFive: cards2, secondFive: cards2}, 0},
		{"test same flush 3", args{firstFive: cards3, secondFive: cards3}, 0},
		{"test same flush 4", args{firstFive: cards4, secondFive: cards4}, 0},

		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := compareFlushes(tt.args.firstFive, tt.args.secondFive); got != tt.want {
				t.Errorf("compareFlushes() = %v, want %v", got, tt.want)
			}
		})
	}
}

func Test_compareStraight(t *testing.T) {

	eightHigh1 := Cards{
		Card{Suit: "C", Value: "4"},
		Card{Suit: "D", Value: "5"},
		Card{Suit: "C", Value: "6"},
		Card{Suit: "H", Value: "7"},
		Card{Suit: "S", Value: "8"},
	}
	eightHigh2 := Cards{
		Card{Suit: "D", Value: "4"},
		Card{Suit: "H", Value: "8"},
		Card{Suit: "S", Value: "5"},
		Card{Suit: "S", Value: "7"},
		Card{Suit: "H", Value: "6"},
	}

	nineHigh := Cards{
		Card{Suit: "C", Value: "5"},
		Card{Suit: "D", Value: "6"},
		Card{Suit: "C", Value: "7"},
		Card{Suit: "H", Value: "8"},
		Card{Suit: "S", Value: "9"},
	}

	fiveHigh := Cards{
		Card{Suit: "H", Value: "4"},
		Card{Suit: "D", Value: "2"},
		Card{Suit: "C", Value: "A"},
		Card{Suit: "C", Value: "3"},
		Card{Suit: "S", Value: "5"},
	}

	aceHigh := Cards{
		Card{Suit: "H", Value: "K"},
		Card{Suit: "D", Value: "J"},
		Card{Suit: "C", Value: "A"},
		Card{Suit: "C", Value: "Q"},
		Card{Suit: "S", Value: "T"},
	}
	kingHigh := Cards{
		Card{Suit: "H", Value: "K"},
		Card{Suit: "D", Value: "J"},
		Card{Suit: "C", Value: "9"},
		Card{Suit: "C", Value: "Q"},
		Card{Suit: "S", Value: "T"},
	}
	type args struct {
		firstFive  Cards
		secondFive Cards
	}
	tests := []struct {
		name string
		args args
		want int
	}{
		//first higher tests
		{"test one higher straight 1 - 1", args{firstFive: nineHigh, secondFive: eightHigh1}, 1},
		{"test one higher straight 2 - 1", args{firstFive: nineHigh, secondFive: eightHigh2}, 1},
		{"test highest higher than lowest - edge case both have ace - 1", args{firstFive: aceHigh, secondFive: fiveHigh}, 1},
		{"test highest higher than second highest - 1", args{firstFive: aceHigh, secondFive: kingHigh}, 1},

		//second higher tests
		{"test one higher straight 1 - 2", args{firstFive: eightHigh1, secondFive: nineHigh}, 2},
		{"test one higher straight 2 - 2", args{firstFive: eightHigh2, secondFive: nineHigh}, 2},
		{"test highest higher than lowest - edge case both have ace - 2", args{firstFive: fiveHigh, secondFive: aceHigh}, 2},
		{"test highest higher than second highest - 2", args{firstFive: kingHigh, secondFive: aceHigh}, 2},

		//same tests
		{"test same straight - dif suits", args{firstFive: eightHigh1, secondFive: eightHigh2}, 0},
		{"test same exact straight nine high", args{firstFive: nineHigh, secondFive: nineHigh}, 0},
		{"test same exact straight king high", args{firstFive: kingHigh, secondFive: kingHigh}, 0},
		{"test same exact straight five high", args{firstFive: fiveHigh, secondFive: fiveHigh}, 0},
		{"test same exact straight ace high", args{firstFive: aceHigh, secondFive: aceHigh}, 0},

		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := compareStraight(tt.args.firstFive, tt.args.secondFive); got != tt.want {
				t.Errorf("compareStraight() = %v, want %v", got, tt.want)
			}
		})
	}
}

func Test_compareThreeOfAKind(t *testing.T) {

	foursAKHigh := Cards{
		Card{Suit: "C", Value: "4"},
		Card{Suit: "S", Value: "4"},
		Card{Suit: "D", Value: "4"},
		Card{Suit: "H", Value: "A"},
		Card{Suit: "C", Value: "K"},
	}
	foursAQHigh := Cards{
		Card{Suit: "C", Value: "4"},
		Card{Suit: "S", Value: "4"},
		Card{Suit: "D", Value: "4"},
		Card{Suit: "H", Value: "A"},
		Card{Suit: "C", Value: "Q"},
	}

	sixesAKHigh := Cards{
		Card{Suit: "C", Value: "K"},
		Card{Suit: "S", Value: "A"},
		Card{Suit: "D", Value: "6"},
		Card{Suit: "H", Value: "6"},
		Card{Suit: "C", Value: "6"},
	}
	sixesJTHigh := Cards{
		Card{Suit: "C", Value: "J"},
		Card{Suit: "S", Value: "T"},
		Card{Suit: "D", Value: "6"},
		Card{Suit: "H", Value: "6"},
		Card{Suit: "C", Value: "6"},
	}
	type args struct {
		firstFive  Cards
		secondFive Cards
	}
	tests := []struct {
		name string
		args args
		want int
	}{

		//BETTER FIRST HAND
		{"test higher trips 1 - 1", args{firstFive: sixesJTHigh, secondFive: foursAKHigh}, 1},
		{"test higher trips 2 - 1", args{firstFive: sixesAKHigh, secondFive: foursAKHigh}, 1},
		{"same trips, higher first card - 1", args{firstFive: sixesAKHigh, secondFive: sixesJTHigh}, 1},
		{"same trips, higher second card - 1", args{firstFive: foursAKHigh, secondFive: foursAQHigh}, 1},

		//BETTER SECOND HAND
		{"test higher trips 1 - 2", args{firstFive: foursAKHigh, secondFive: sixesJTHigh}, 2},
		{"test higher trips 2 - 2", args{firstFive: foursAKHigh, secondFive: sixesAKHigh}, 2},
		{"same trips, higher first card - 2", args{firstFive: sixesJTHigh, secondFive: sixesAKHigh}, 2},
		{"same trips, higher second card - 2", args{firstFive: foursAQHigh, secondFive: foursAKHigh}, 2},

		//SAME VALUE HAND
		{"same value 1", args{firstFive: foursAKHigh, secondFive: foursAKHigh}, 0},
		{"same value 2", args{firstFive: foursAQHigh, secondFive: foursAQHigh}, 0},
		{"same value 3", args{firstFive: sixesAKHigh, secondFive: sixesAKHigh}, 0},
		{"same value 4", args{firstFive: sixesJTHigh, secondFive: sixesJTHigh}, 0},

		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := compareThreeOfAKind(tt.args.firstFive, tt.args.secondFive); got != tt.want {
				t.Errorf("compareThreeOfAKind() = %v, want %v", got, tt.want)
			}
		})
	}
}

func Test_compareTwoPair(t *testing.T) {
	kingsAndFoursAHigh := Cards{
		Card{Suit: "C", Value: "4"},
		Card{Suit: "S", Value: "4"},
		Card{Suit: "D", Value: "K"},
		Card{Suit: "H", Value: "A"},
		Card{Suit: "C", Value: "K"},
	}
	acesAndFoursQHigh := Cards{
		Card{Suit: "C", Value: "4"},
		Card{Suit: "S", Value: "4"},
		Card{Suit: "D", Value: "A"},
		Card{Suit: "H", Value: "A"},
		Card{Suit: "C", Value: "Q"},
	}
	kingsAndSixesAHigh := Cards{
		Card{Suit: "C", Value: "K"},
		Card{Suit: "S", Value: "A"},
		Card{Suit: "D", Value: "K"},
		Card{Suit: "H", Value: "6"},
		Card{Suit: "C", Value: "6"},
	}
	kingsAndSixesQHigh := Cards{
		Card{Suit: "C", Value: "K"},
		Card{Suit: "S", Value: "K"},
		Card{Suit: "D", Value: "Q"},
		Card{Suit: "H", Value: "6"},
		Card{Suit: "C", Value: "6"},
	}
	type args struct {
		firstFive  Cards
		secondFive Cards
	}
	tests := []struct {
		name string
		args args
		want int
	}{
		//First hand better
		{"test higher first pair - 1", args{firstFive: acesAndFoursQHigh, secondFive: kingsAndSixesAHigh}, 1},
		{"test same first pair, higher second pair - 1", args{firstFive: kingsAndSixesQHigh, secondFive: kingsAndFoursAHigh}, 1},
		{"test same pairs, higher kicker - 1", args{firstFive: kingsAndSixesAHigh, secondFive: kingsAndSixesQHigh}, 1},

		//Second hand better
		{"test higher first pair - 2", args{firstFive: kingsAndSixesAHigh, secondFive: acesAndFoursQHigh}, 2},
		{"test same first pair, higher second pair - 2", args{firstFive: kingsAndFoursAHigh, secondFive: kingsAndSixesQHigh}, 2},
		{"test same pairs, higher kicker - 2", args{firstFive: kingsAndSixesQHigh, secondFive: kingsAndSixesAHigh}, 2},

		//Hands the same
		{"test same 1", args{firstFive: kingsAndFoursAHigh, secondFive: kingsAndFoursAHigh}, 0},
		{"test same 2", args{firstFive: acesAndFoursQHigh, secondFive: acesAndFoursQHigh}, 0},
		{"test same 3", args{firstFive: kingsAndSixesAHigh, secondFive: kingsAndSixesAHigh}, 0},
		{"test same 4", args{firstFive: kingsAndSixesQHigh, secondFive: kingsAndSixesQHigh}, 0},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := compareTwoPair(tt.args.firstFive, tt.args.secondFive); got != tt.want {
				t.Errorf("compareTwoPair() = %v, want %v", got, tt.want)
			}
		})
	}
}
