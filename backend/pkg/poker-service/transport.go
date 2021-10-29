package pokerservice

import (
	"context"
	"errors"

	"github.com/go-kit/kit/transport"
	"github.com/go-kit/kit/transport/grpc"
	"github.com/go-kit/log"

	p_v1alpha1 "github.com/zshearin/poker/backend/api/v1alpha1"
)

type grpcServer struct {
	getGame grpc.Handler
	p_v1alpha1.UnimplementedPokerAPIServer
}

func NewGrpcTransport(ep *Endpoints, logger log.Logger) p_v1alpha1.PokerAPIServer {
	errorHandler := transport.NewLogErrorHandler(logger)
	options := []grpc.ServerOption{
		grpc.ServerErrorHandler(errorHandler),
	}

	server := &grpcServer{
		getGame: grpc.NewServer(
			ep.getGame,
			decodeGetGameRequest,
			encodeGetGameResponse,
			options...,
		),
	}

	return server
}

func (s *grpcServer) GetGame(ctx context.Context, req *p_v1alpha1.GetGameRequest) (*p_v1alpha1.GetGameResponse, error) {
	_, resp, err := s.getGame.ServeGRPC(ctx, req)
	if err != nil {
		return nil, err
	}
	return resp.(*p_v1alpha1.GetGameResponse), nil
}

func decodeGetGameRequest(_ context.Context, r interface{}) (interface{}, error) {
	req, ok := r.(*p_v1alpha1.GetGameRequest)
	if !ok {
		return nil, errors.New("cannot process request")
	}

	return req, nil
}

func encodeGetGameResponse(_ context.Context, r interface{}) (interface{}, error) {
	return r.(*p_v1alpha1.GetGameResponse), nil
}
