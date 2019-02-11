package pongersrv

import (
	"context"
	"time"

	"github.com/iheanyi/go-tracing-example/rpc/ponger"
	opentracing "github.com/opentracing/opentracing-go"
	"github.com/opentracing/opentracing-go/ext"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
)

type pongerServer struct{}

func New() ponger.PongerServer {
	return &pongerServer{}
}

func (s *pongerServer) Pong(ctx context.Context, req *ponger.PongRequest) (*ponger.PongResponse, error) {
	span := opentracing.SpanFromContext(ctx)
	ext.SamplingPriority.Set(span, 1)
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
