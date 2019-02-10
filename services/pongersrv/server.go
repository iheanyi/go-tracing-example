package pongersrv

import (
	"context"
	"time"

	"github.com/iheanyi/go-tracing-example/rpc/ponger"
)

type pongerServer struct {
}

func New() ponger.Ponger {
	return &pongerServer{}
}

func (s *pongerServer) Pong(ctx context.Context, req *ponger.PongRequest) (*ponger.PongResponse, error) {
	if req.Delay > 0 {
		time.Sleep(time.Duration(req.Delay) * time.Second)
	}
	return &ponger.PongResponse{
		Body: "Ping!",
	}, nil
}
