package pingersrv

import (
	"context"
	"time"

	"github.com/iheanyi/go-tracing-example/rpc/pinger"
	"github.com/iheanyi/go-tracing-example/rpc/ponger"
	opentracing "github.com/opentracing/opentracing-go"
	"github.com/opentracing/opentracing-go/ext"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
)

type pingerServer struct {
	pongClient ponger.PongerClient
}

func New(pongClient ponger.PongerClient) pinger.PingerServer {
	return &pingerServer{
		pongClient: pongClient,
	}
}

func (s *pingerServer) Ping(ctx context.Context, req *pinger.PingRequest) (*pinger.PingResponse, error) {
	span := opentracing.SpanFromContext(ctx)
	ext.SamplingPriority.Set(span, 1)
	if req.Message == "" {
		return nil, status.Errorf(codes.InvalidArgument, "message cannot be blank")
	}

	childSpanDemo(ctx)
	return &pinger.PingResponse{
		Body: "Pong!",
	}, nil
}

func (s *pingerServer) PingPong(ctx context.Context, req *pinger.PingPongRequest) (*pinger.PingPongResponse, error) {
	span := opentracing.SpanFromContext(ctx)
	ext.SamplingPriority.Set(span, 1)
	// Do something here.
	pongReq := &ponger.PongRequest{
		Message: req.Message,
		Delay:   req.Delay,
	}

	childSpanDemo(ctx)
	_, err := s.pongClient.Pong(ctx, pongReq)
	if err != nil {
		return nil, err
	}

	return &pinger.PingPongResponse{
		Body: "Ping Pong!",
	}, nil
}

func childSpanDemo(ctx context.Context) {
	span, _ := opentracing.StartSpanFromContext(ctx, "ChildSpanCall")
	defer span.Finish()
	time.Sleep(2 * time.Second)
}
