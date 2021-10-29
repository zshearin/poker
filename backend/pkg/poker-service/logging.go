package pokerservice

import (
	"context"

	"github.com/go-kit/log"
	p_v1alpha1 "github.com/zshearin/poker/backend/api/v1alpha1"
)

type loggingService struct {
	logger log.Logger
	Service
}

func NewLoggingService(logger log.Logger, s Service) Service {
	return &loggingService{logger, s}
}

func (s *loggingService) GetGame(ctx context.Context, req *p_v1alpha1.GetGameRequest) (*p_v1alpha1.GetGameResponse, error) {
	return s.Service.GetGame(ctx, req)
}
