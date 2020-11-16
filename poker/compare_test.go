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
