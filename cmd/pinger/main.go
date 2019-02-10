package main

import (
	"log"
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
	if err != nil {
		log.Fatalf("error instantiating tracer: %v", err)
	}
	defer closer.Close()

	hooks := ottwirp.NewOpenTracingHooks(tracer)
	twirpHandler := pinger.NewPingerServer(server, hooks)
	log.Fatal(http.ListenAndServe(":8082", twirpHandler))
}
