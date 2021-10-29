package pokerservice

import (
	"context"
	"errors"

	"github.com/go-kit/log"
	p_v1alpha1 "github.com/zshearin/poker/backend/api/v1alpha1"

	dealer "github.com/zshearin/poker/backend/pkg/dealer"
)

type Service interface {
	GetGame(ctx context.Context, req *p_v1alpha1.GetGameRequest) (*p_v1alpha1.GetGameResponse, error)
}

type service struct {
	logger log.Logger
}

func NewService(
	l log.Logger,
) Service {
	return &service{
		logger: l,
	}
}

func (s *service) GetGame(ctx context.Context, req *p_v1alpha1.GetGameRequest) (*p_v1alpha1.GetGameResponse, error) {

	if req.Hands > 10 {
		return nil, errors.New("maximum number of hands to deal in is 10")
	} else if req.Hands < 2 {
		return nil, errors.New("minimum number of hands to deal in is 2")
	}

	print := true

	deck := dealer.GetDeck()
	deck.Shuffle()
	deck.Shuffle()
	//deck.PrintOrder()

	game := deck.GetDeal(int(req.Hands))
	if print {
		game.PrintBoardAndHands()
		game.PrintRanksAndBestFive()
	}

	deal := converGameToDeal(game)
	return &p_v1alpha1.GetGameResponse{Deal: deal}, nil

}

func converGameToDeal(game dealer.Deal) *p_v1alpha1.Deal {
	var newHandResults []*p_v1alpha1.HandResult

	for _, handResult := range game.HandResults {

		newBestFive := []*p_v1alpha1.Card{}
		for _, card := range handResult.Player.BestFive {
			newCard := &p_v1alpha1.Card{
				Suit:  card.Suit,
				Value: card.Value,
			}
			newBestFive = append(newBestFive, newCard)
		}

		newPlayer := &p_v1alpha1.Player{
			BestFive: newBestFive,
			Number:   int32(handResult.Player.Num),
			HandName: handResult.Player.HandName,
		}

		newHandResult := &p_v1alpha1.HandResult{
			RelativeHandRank: int32(handResult.RelativeHandRank),
			Player:           newPlayer,
		}

		newHandResults = append(newHandResults, newHandResult)
	}

	var newBoard []*p_v1alpha1.Card

	for _, card := range game.Board {

		newCard := &p_v1alpha1.Card{
			Suit:  card.Suit,
			Value: card.Value,
		}
		newBoard = append(newBoard, newCard)
	}

	allHands := []*p_v1alpha1.Hand{}
	for _, hand := range game.Hands {
		newCardsList := []*p_v1alpha1.Card{}
		for _, card := range hand {
			newCard := &p_v1alpha1.Card{
				Suit:  card.Suit,
				Value: card.Value,
			}
			newCardsList = append(newCardsList, newCard)
		}
		newHand := &p_v1alpha1.Hand{
			Cards: newCardsList,
		}
		allHands = append(allHands, newHand)
	}

	return &p_v1alpha1.Deal{
		Hands:      allHands,
		Board:      newBoard,
		HandResult: newHandResults,
	}
}
