package poker

//DealHoldEm deals 2 cards to the number of hands passed in
//eg: if 6 is passed in, it will return 6 hands of 2
func (d *Deck) DealHoldEm(numHands int) Hands {
	return d.Deal(numHands, 2)
}

//GetFlop burns a card and returns the 3 after
func (d *Deck) GetFlop() Cards {
	return d.BurnAndFlip(3)
}

//GetTurn burns a card and returns the card after
func (d *Deck) GetTurn() Cards {
	return d.BurnAndFlip(1)
}

//GetRiver burns  a card and returns the card after
func (d *Deck) GetRiver() Cards {
	return d.BurnAndFlip(1)
}

//Reset resets the deck index to 0
func (d *Deck) Reset() {
	d.NextCardIndex = 0
}

//GetCard gets the next card in a deck and updates the index of the deck object
func (d *Deck) GetCard() Card {
	var card Card

	card = d.Cards[d.NextCardIndex]
	d.NextCardIndex++
	return card
}

//BurnAndFlip takes in a parameter that is the number of cards to get
//It first skips a card then grabs the desired number of cards
func (d *Deck) BurnAndFlip(numCards int) Cards {

	var cards Cards

	//Burn a card:
	d.GetCard()

	for i := 0; i < numCards; i++ {

		nextCard := d.GetCard()
		cards = append(cards, nextCard)
	}
	return cards

}

//Deal takes in the number of hands and the number of cards per hand
//and creates the hands from the deck, updating the next card as it goes
func (d *Deck) Deal(numHands, numCards int) Hands {

	var hands Hands

	for i := 0; i < numHands; i++ {
		hand := []Card{}
		hands = append(hands, hand)
	}

	for i := 0; i < numCards; i++ {
		for j := 0; j < numHands; j++ {
			curHand := hands[j]
			nextCard := d.GetCard()
			curHand = append(curHand, nextCard)
			//curHand = append(curHand, d.Cards[d.NextCardIndex])
			hands[j] = curHand

			//d.NextCardIndex++
		}
	}

	return hands
}
