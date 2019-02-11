package main

import (
	"log"
	"net/http"
	"os"

	"github.com/iheanyi/go-tracing-example/rpc/pinger"
	"github.com/iheanyi/go-tracing-example/rpc/ponger"
	"github.com/iheanyi/go-tracing-example/services/pingersrv"
	ottwirp "github.com/iheanyi/twirp-opentracing"
	opentracing "github.com/opentracing/opentracing-go"
	"github.com/uber/jaeger-client-go/config"
	jaegerlog "github.com/uber/jaeger-client-go/log"
)

func main() {
	os.Setenv("JAEGER_SERVICE_NAME", "pingersrv")
	os.Setenv("JAEGER_ENDPOINT", "http://localhost:14268/api/traces")
	os.Setenv("JAEGER_REPORTER_LOG_SPANS", "true")
	client := ponger.NewPongerProtobufClient("http://localhost:8083", &http.Client{})
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
	server := pingersrv.New(client)
	twirpHandler := pinger.NewPingerServer(server, hooks)
	log.Fatal(http.ListenAndServe(":8082", twirpHandler))
}
