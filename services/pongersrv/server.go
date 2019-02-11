package pongersrv

import (
	"context"
	"time"

	"github.com/iheanyi/go-tracing-example/rpc/ponger"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
)

type pongerServer struct{}

func New() ponger.PongerServer {
	return &pongerServer{}
}

func (s *pongerServer) Pong(ctx context.Context, req *ponger.PongRequest) (*ponger.PongResponse, error) {
	if req.Message == "" {
		return nil, status.Errorf(codes.InvalidArgument, "message cannot be blank")
	}
	if req.Delay > 0 {
		time.Sleep(time.Duration(req.Delay) * time.Second)
	}
	return &ponger.PongResponse{
		Body: "Ping!",
	}, nil
}
