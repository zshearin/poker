package pokerservice

import (
	"context"
	//"log"

	"github.com/go-kit/kit/endpoint"
	"github.com/go-kit/kit/log"

	p_v1alpha1 "github.com/zshearin/poker/api/v1alpha1"
)

type Endpoints struct {
	getGame endpoint.Endpoint
}

func CreateEndpoints(s Service, logger log.Logger) *Endpoints {

	var getGame endpoint.Endpoint
	{
		getGame = makeGetGameEndpoint(s)
	}

	return &Endpoints{
		getGame: getGame,
	}
}

func makeGetGameEndpoint(s Service) endpoint.Endpoint {
	return func(ctx context.Context, r interface{}) (interface{}, error) {
		req := r.(*p_v1alpha1.GetGameRequest)
		return s.GetGame(ctx, req)
	}
}
