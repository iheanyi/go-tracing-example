package main

import (
	"log"
	"net"
	"os"

	grpc_opentracing "github.com/grpc-ecosystem/go-grpc-middleware/tracing/opentracing"

	"github.com/iheanyi/go-tracing-example/rpc/pinger"
	"github.com/iheanyi/go-tracing-example/rpc/ponger"
	"github.com/iheanyi/go-tracing-example/services/pingersrv"
	opentracing "github.com/opentracing/opentracing-go"
	"github.com/uber/jaeger-client-go/config"
	jaegerlog "github.com/uber/jaeger-client-go/log"
	"google.golang.org/grpc"
)

func main() {
	os.Setenv("JAEGER_SERVICE_NAME", "pingersrv")
	os.Setenv("JAEGER_ENDPOINT", "http://localhost:14268/api/traces")
	os.Setenv("JAEGER_REPORTER_LOG_SPANS", "true")
	conn, err := grpc.Dial("127.0.0.1:8083", grpc.WithInsecure(), grpc.WithUnaryInterceptor(grpc_opentracing.UnaryClientInterceptor()))
	if err != nil {
		log.Fatalf("failed to dial: %v", err)
	}
	defer conn.Close()

	client := ponger.NewPongerClient(conn)
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

	lis, err := net.Listen("tcp", "localhost:8082")
	if err != nil {
		log.Fatalf("failed to create listener: %v", err)
	}
	grpcServer := grpc.NewServer(grpc.UnaryInterceptor(grpc_opentracing.UnaryServerInterceptor()))
	pinger.RegisterPingerServer(grpcServer, pingersrv.New(client))
	grpcServer.Serve(lis)
}
