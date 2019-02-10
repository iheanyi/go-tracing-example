package pingersrv

import (
	"context"

	"github.com/iheanyi/go-tracing-example/rpc/pinger"
	opentracing "github.com/opentracing/opentracing-go"
	"github.com/twitchtv/twirp"
)

type pingerServer struct{}

func New() pinger.Pinger {
	return &pingerServer{}
}

func (s *pingerServer) Ping(ctx context.Context, req *pinger.PingRequest) (*pinger.PingResponse, error) {
	if req.Message == "" {
		return nil, twirp.InvalidArgumentError("message", "This can't be blank!")
	}

	childSpanDemo(ctx)
	return &pinger.PingResponse{
		Body: "Pong!",
	}, nil
}

func childSpanDemo(ctx context.Context) {
	span, _ := opentracing.StartSpanFromContext(ctx, "ChildSpanCall")
	defer span.Finish()
}
