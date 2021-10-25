package pokerservice

import (
	"context"

	"github.com/go-kit/log"
	p_v1alpha1 "github.com/zshearin/poker/api/v1alpha1"
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

	/*
		message GetGameResponse {
		    Deal deal = 1;
		}

		message Deal {
		    repeated Hand hands = 1;
		    repeated Card board = 2;
		    repeated HandResult hand_result = 3;
		}

		message HandResult {
		    Player player = 1;
		    int32 relative_hand_rank = 2;
		}
		message Player {
		    int32 number = 1;
		    repeated Card best_five = 2;
		    string hand_name = 3;
		}

		message Card {
		    string suit = 1;
		    string value= 2;
		}

		message Hand {
		    repeated Card cards = 1;
		}


	*/

	card := &p_v1alpha1.Card{
		Suit:  "this is the suit",
		Value: "this is the value",
	}

	var hands []*p_v1alpha1.Hand
	var board []*p_v1alpha1.Card
	var handResults []*p_v1alpha1.HandResult

	board = append(board, card)
	board = append(board, card)

	return &p_v1alpha1.GetGameResponse{
		Deal: &p_v1alpha1.Deal{
			Hands:      hands,
			Board:      board,
			HandResult: handResults,
		},
	}, nil
}
