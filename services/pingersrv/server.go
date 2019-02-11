package pingersrv

import (
	"context"
	"log"

	"github.com/iheanyi/go-tracing-example/rpc/pinger"
	"github.com/iheanyi/go-tracing-example/rpc/ponger"
	opentracing "github.com/opentracing/opentracing-go"
	"github.com/twitchtv/twirp"
)

type pingerServer struct {
	pongClient ponger.Ponger
}

func New(pongClient ponger.Ponger) pinger.Pinger {
	return &pingerServer{
		pongClient: pongClient,
	}
}

func (s *pingerServer) Ping(ctx context.Context, req *pinger.PingRequest) (*pinger.PingResponse, error) {
	if req.Message == "" {
		return nil, twirp.InvalidArgumentError("message", "this can't be blank")
	}

	childSpanDemo(ctx)
	return &pinger.PingResponse{
		Body: "Pong!",
	}, nil
}

func (s *pingerServer) PingPong(ctx context.Context, req *pinger.PingPongRequest) (*pinger.PingPongResponse, error) {

	// Do something here.
	pongReq := &ponger.PongRequest{
		Message: req.Message,
		Delay:   req.Delay,
	}

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
	log.Println("We in a child span.")
}
