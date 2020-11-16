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

/*
func TestCompareTwoBestFive(t *testing.T) {

	cards1 := Cards{
		Card{Suit: "C", Value: "4", Number: 4},
		Card{Suit: "C", Value: "5", Number: 5},
		Card{Suit: "C", Value: "6", Number: 6},
		Card{Suit: "C", Value: "7", Number: 7},
		Card{Suit: "C", Value: "8", Number: 8},
	}

	cards2 := Cards{
		Card{Suit: "C", Value: "5", Number: 5},
		Card{Suit: "C", Value: "6", Number: 6},
		Card{Suit: "C", Value: "7", Number: 7},
		Card{Suit: "C", Value: "8", Number: 8},
		Card{Suit: "C", Value: "9", Number: 9},
	}

	type args struct {
		firstFive  Cards
		secondFive Cards
	}
	tests := []struct {
		name    string
		args    args
		want    int
		wantErr bool
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, err := CompareTwoBestFive(tt.args.firstFive, tt.args.secondFive)
			if (err != nil) != tt.wantErr {
				t.Errorf("CompareTwoBestFive() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if got != tt.want {
				t.Errorf("CompareTwoBestFive() = %v, want %v", got, tt.want)
			}
		})
	}
}
*/
