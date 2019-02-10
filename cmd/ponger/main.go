package main

import (
	"log"
	"net/http"
	"os"

	"github.com/iheanyi/go-tracing-example/rpc/ponger"
	"github.com/iheanyi/go-tracing-example/services/pongersrv"
	ottwirp "github.com/iheanyi/twirp-opentracing"
	opentracing "github.com/opentracing/opentracing-go"
	"github.com/uber/jaeger-client-go/config"
	jaegerlog "github.com/uber/jaeger-client-go/log"
)

func main() {
	os.Setenv("JAEGER_SERVICE_NAME", "pongersrv")
	os.Setenv("JAEGER_ENDPOINT", "http://localhost:14268/api/traces")
	os.Setenv("JAEGER_REPORTER_LOG_SPANS", "true")

	cfg, err := config.FromEnv()
	if err != nil {
		// parsing errors might happen here, such as when we get a string where we expect a number
		log.Printf("Could not parse Jaeger env vars: %s", err.Error())
		return
	}
	jLogger := jaegerlog.StdLogger

	tracer, closer, err := cfg.NewTracer(config.Logger(jLogger))
	if err != nil {
		log.Fatalf("error instantiating tracer: %v", err)
	}
	defer closer.Close()
	opentracing.SetGlobalTracer(tracer)

	hooks := ottwirp.NewOpenTracingHooks(tracer)
	server := pongersrv.New()
	twirpHandler := ponger.NewPongerServer(server, hooks)
	log.Fatal(http.ListenAndServe(":8083", twirpHandler))
}
