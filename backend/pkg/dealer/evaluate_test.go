package poker

import (
	"reflect"
	"sort"
	"testing"
)

/*
func TestCards_GetFiveBest(t *testing.T) {
	type args struct {
		cards      Cards
		printValue bool
	}
	tests := []struct {
		name  string
		c     *Cards
		args  args
		want  Cards
		want1 int
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, got1 := tt.c.GetFiveBest(tt.args.cards, tt.args.printValue)
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("Cards.GetFiveBest() got = %v, want %v", got, tt.want)
			}
			if got1 != tt.want1 {
				t.Errorf("Cards.GetFiveBest() got1 = %v, want %v", got1, tt.want1)
			}
		})
	}
}
*/
func Test_checkForFiveInARow(t *testing.T) {

	cards1 := Cards{
		Card{Value: "9", Suit: "H"},
		Card{Value: "8", Suit: "D"},
		Card{Value: "7", Suit: "D"},
		Card{Value: "6", Suit: "H"},
		Card{Value: "5", Suit: "D"},
	}

	cards2 := Cards{
		Card{Value: "9", Suit: "H"},
		Card{Value: "8", Suit: "D"},
		Card{Value: "7", Suit: "D"},
		Card{Value: "6", Suit: "H"},
	}

	cards3 := Cards{
		Card{Value: "9", Suit: "H"},
		Card{Value: "9", Suit: "S"},
		Card{Value: "8", Suit: "D"},
		Card{Value: "7", Suit: "D"},
		Card{Value: "6", Suit: "H"},
	}

	//Test 4: low straight
	cards4 := Cards{
		Card{Value: "A", Suit: "H"},
		Card{Value: "2", Suit: "S"},
		Card{Value: "4", Suit: "D"},
		Card{Value: "5", Suit: "D"},
		Card{Value: "3", Suit: "H"},
	}

	cards4ExpectedOutput := Cards{
		Card{Value: "1", Suit: "H"},
		Card{Value: "2", Suit: "S"},
		Card{Value: "4", Suit: "D"},
		Card{Value: "5", Suit: "D"},
		Card{Value: "3", Suit: "H"},
	}
	sort.Sort(ByNumber(cards4ExpectedOutput))

	//Test 5 setup: high straight
	cards5 := Cards{
		Card{Value: "Q", Suit: "H"},
		Card{Value: "K", Suit: "S"},
		Card{Value: "J", Suit: "D"},
		Card{Value: "A", Suit: "D"},
		Card{Value: "T", Suit: "H"},
	}

	type args struct {
		cards Cards
	}
	tests := []struct {
		name  string
		args  args
		want  bool
		want1 Cards
	}{
		{"test basic straight", args{cards: cards1}, true, cards1},
		{"test only four", args{cards: cards2}, false, Cards{}},
		{"test two of same in a row, but no straight", args{cards: cards3}, false, Cards{}},
		{"test low straight", args{cards: cards4}, true, cards4ExpectedOutput},
		{"test high straight", args{cards: cards5}, true, cards5},

		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, got1 := checkForFiveInARow(tt.args.cards)
			if got != tt.want {
				t.Errorf("checkForFiveInARow() got = %v, want %v", got, tt.want)
			}
			if !reflect.DeepEqual(got1, tt.want1) {
				t.Errorf("checkForFiveInARow() got1 = %v, want %v", got1, tt.want1)
			}
		})
	}
}
