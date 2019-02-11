package main

import (
	"log"
	"net"
	"os"

	"github.com/grpc-ecosystem/grpc-opentracing/go/otgrpc"
	"github.com/iheanyi/go-tracing-example/rpc/ponger"
	"github.com/iheanyi/go-tracing-example/services/pongersrv"
	opentracing "github.com/opentracing/opentracing-go"
	"github.com/uber/jaeger-client-go/config"
	jaegerlog "github.com/uber/jaeger-client-go/log"
	"google.golang.org/grpc"
)

func main() {
	os.Setenv("JAEGER_SERVICE_NAME", "pongersrv")
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

	lis, err := net.Listen("tcp", "localhost:8083")
	if err != nil {
		log.Fatalf("failed to create listener: %v", err)
	}
	grpcServer := grpc.NewServer(grpc.UnaryInterceptor(otgrpc.OpenTracingServerInterceptor(tracer)))
	ponger.RegisterPongerServer(grpcServer, pongersrv.New())
	grpcServer.Serve(lis)
}
