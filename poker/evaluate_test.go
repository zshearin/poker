package poker

import (
	"reflect"
	"testing"
)

func TestCards_Evaluate(t *testing.T) {
	tests := []struct {
		name string
		c    *Cards
		want Cards
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := tt.c.Evaluate(); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("Cards.Evaluate() = %v, want %v", got, tt.want)
			}
		})
	}
}

//Straight flush edge case:
/*
	cards := poker.Cards{
		poker.Card{
			Suit:   "H",
			Value:  "A",
			Number: 14,
		},
		poker.Card{
			Suit:   "H",
			Value:  "2",
			Number: 2,
		},
		poker.Card{
			Suit:   "H",
			Value:  "3",
			Number: 3,
		},
		poker.Card{
			Suit:   "H",
			Value:  "4",
			Number: 4,
		},
		poker.Card{
			Suit:   "H",
			Value:  "5",
			Number: 5,
		},
	}
*/

//Straight edge case:
/*
	cards := poker.Cards{
		poker.Card{
			Suit:   "H",
			Value:  "A",
			Number: 14,
		},
		poker.Card{
			Suit:   "H",
			Value:  "2",
			Number: 2,
		},
		poker.Card{
			Suit:   "H",
			Value:  "3",
			Number: 3,
		},
		poker.Card{
			Suit:   "H",
			Value:  "4",
			Number: 4,
		},
		poker.Card{
			Suit:   "S",
			Value:  "5",
			Number: 5,
		},
	}
*/

//Quads
/*
	cards := poker.Cards{
		poker.Card{
			Suit:   "S",
			Value:  "2",
			Number: 2,
		},
		poker.Card{
			Suit:   "H",
			Value:  "2",
			Number: 2,
		},
		poker.Card{
			Suit:   "D",
			Value:  "2",
			Number: 2,
		},
		poker.Card{
			Suit:   "C",
			Value:  "2",
			Number: 2,
		},
		poker.Card{
			Suit:   "S",
			Value:  "5",
			Number: 5,
		},
	}
*/
