package poker

import (
	"reflect"
	"testing"
)

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

//High card
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
			Card{
				Suit:   "H",
				Value:  "10",
				Number: 10,
			},
		poker.Card{
			Suit:   "S",
			Value:  "5",
			Number: 5,
		},

		poker.Card{
			Suit:   "S",
			Value:  "10",
			Number: 10,
		},

		poker.Card{
			Suit:   "D",
			Value:  "J",
			Number: 11,
		},

		poker.Card{
			Suit:   "D",
			Value:  "Q",
			Number: 12,
		},
				poker.Card{
					Suit:   "H",
					Value:  "K",
					Number: 13,
				},
			poker.Card{
				Suit:   "D",
				Value:  "A",
				Number: 14,
			},
			Card{
				Suit:   "H",
				Value:  "5",
				Number: 5,
			},
	}

*/

func TestCards_GetFiveBest(t *testing.T) {
	type args struct {
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
			got, got1 := tt.c.GetFiveBest(tt.args.printValue)
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("Cards.GetFiveBest() got = %v, want %v", got, tt.want)
			}
			if got1 != tt.want1 {
				t.Errorf("Cards.GetFiveBest() got1 = %v, want %v", got1, tt.want1)
			}
		})
	}
}
