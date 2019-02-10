package main

import (
	"net/http"

	"github.com/iheanyi/go-tracing-example/rpc/pinger"
	"github.com/iheanyi/go-tracing-example/services/pingersrv"
	ottwirp "github.com/iheanyi/twirp-opentracing"
	"github.com/uber/jaeger-client-go/config"
)

func main() {
	server := pingersrv.New()

	tracer, closer, err := config.Configuration{
		ServiceName: "pingersrv",
	}.NewTracer()
	defer closer.Close()

	hooks := ottwirp.NewOpenTracingHooks(tracer)

	// TODO: Add opentracing hooks here
	twirpHandler := pinger.NewPingerServer(server, nil)

	http.ListenAndServe(":8082", twirpHandler)
}
