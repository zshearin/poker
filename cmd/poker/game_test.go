package poker

/*

	testGame := poker.Game{
		Hands: []poker.Cards{
			//hand 1
			{
				poker.Card{
					Suit:   "C",
					Value:  "4",
					Number: 4,
				},
				poker.Card{
					Suit:   "D",
					Value:  "9",
					Number: 9,
				},
			},
			//hand 2
			{
				poker.Card{
					Suit:   "C",
					Value:  "9",
					Number: 9,
				},
				poker.Card{
					Suit:   "H",
					Value:  "3",
					Number: 3,
				},
			},
			//hand 3
			{
				poker.Card{
					Suit:   "C",
					Value:  "8",
					Number: 8,
				},
				poker.Card{
					Suit:   "D",
					Value:  "3",
					Number: 3,
				},
			},
			//hand 4
			{
				poker.Card{
					Suit:   "H",
					Value:  "J",
					Number: 11,
				},
				poker.Card{
					Suit:   "S",
					Value:  "T",
					Number: 10,
				},
			},
			//hand 5
			{
				poker.Card{
					Suit:   "D",
					Value:  "T",
					Number: 10,
				},
				poker.Card{
					Suit:   "H",
					Value:  "A",
					Number: 14,
				},
			},
		},
		Flop: poker.Cards{

			poker.Card{
				Suit:   "D",
				Value:  "5",
				Number: 5,
			},
			poker.Card{
				Suit:   "D",
				Value:  "2",
				Number: 2,
			},
			poker.Card{
				Suit:   "D",
				Value:  "K",
				Number: 13,
			},
		},
		Turn: poker.Cards{
			poker.Card{
				Suit:   "H",
				Value:  "T",
				Number: 10,
			},
		},
		River: poker.Cards{
			poker.Card{
				Suit:   "C",
				Value:  "T",
				Number: 10,
			},
		},
	}

	var evalCards []poker.Cards

	for _, curCards := range testGame.Hands {

		var curCardList poker.Cards

		curCardList = append(curCardList, curCards...)
		curCardList = append(curCardList, testGame.Flop...)
		curCardList = append(curCardList, testGame.Turn...)
		curCardList = append(curCardList, testGame.River...)

		evalCards = append(evalCards, curCardList)
	}
	testGame.CardsForEval = evalCards

	for _, value := range testGame.CardsForEval {

		printBestFive(value)
	}


*/
