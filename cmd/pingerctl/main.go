package main

import (
	"context"
	"log"
	"os"

	kingpin "github.com/alecthomas/kingpin"
	"github.com/grpc-ecosystem/grpc-opentracing/go/otgrpc"
	"github.com/iheanyi/go-tracing-example/rpc/pinger"
	"github.com/iheanyi/go-tracing-example/rpc/ponger"
	opentracing "github.com/opentracing/opentracing-go"
	"github.com/uber/jaeger-client-go/config"
	jaegerlog "github.com/uber/jaeger-client-go/log"
	"google.golang.org/grpc"
)

func main() {
	os.Setenv("JAEGER_SERVICE_NAME", "pingersrv")
	// os.Setenv("JAEGER_ENDPOINT", "http://localhost:14268/api/traces")

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

	app := kingpin.New("pingerctl", "Tool to interact with ping and pong services")
	pingCmd(app.Command("ping", "ping pinger service"), tracer)
	pongCmd(app.Command("pong", "ping ponger service"), tracer)
	pingPongCmd(app.Command("ping-pong", "ping both the pinger and ponger services"), tracer)

	_, err = app.Parse(os.Args[1:])
	if err != nil {
		log.Fatalf("failed to run: %v", err)
	}
}

func pingCmd(cmd *kingpin.CmdClause, tracer opentracing.Tracer) {
	cfg := struct {
		Message string
	}{}
	cmd.Arg("message", "message to send to the service").StringVar(&cfg.Message)

	cmd.Action(func(*kingpin.ParseContext) error {
		conn, err := grpc.Dial("127.0.0.1:8082", grpc.WithInsecure(), grpc.WithUnaryInterceptor(otgrpc.OpenTracingClientInterceptor(tracer)))
		if err != nil {
			log.Fatalf("failed to dial: %v", err)
		}
		defer conn.Close()

		client := pinger.NewPingerClient(conn)
		res, err := client.Ping(context.Background(), &pinger.PingRequest{
			Message: cfg.Message,
		})
		if err != nil {
			return err
		}
		log.Println(res.Body)
		return nil
	})
}

func pongCmd(cmd *kingpin.CmdClause, tracer opentracing.Tracer) {
	cfg := struct {
		Message string
		Delay   int64
	}{}
	cmd.Arg("message", "message to send to the service").StringVar(&cfg.Message)
	cmd.Flag("delay", "simulate latency to endpoint").Int64Var(&cfg.Delay)

	cmd.Action(func(*kingpin.ParseContext) error {
		conn, err := grpc.Dial("127.0.0.1:8083", grpc.WithInsecure(), grpc.WithUnaryInterceptor(otgrpc.OpenTracingClientInterceptor(tracer)))
		if err != nil {
			log.Fatalf("failed to dial: %v", err)
		}
		defer conn.Close()
		client := ponger.NewPongerClient(conn)
		res, err := client.Pong(context.Background(), &ponger.PongRequest{
			Message: cfg.Message,
			Delay:   cfg.Delay,
		})
		if err != nil {
			return err
		}
		log.Println(res.Body)
		return nil
	})
}

func pingPongCmd(cmd *kingpin.CmdClause, tracer opentracing.Tracer) {
	cfg := struct {
		Message string
		Delay   int64
	}{}
	cmd.Arg("message", "message to send to the service").StringVar(&cfg.Message)
	cmd.Flag("delay", "delay to simulate latency in the pong service").Int64Var(&cfg.Delay)

	cmd.Action(func(*kingpin.ParseContext) error {
		conn, err := grpc.Dial("127.0.0.1:8082", grpc.WithInsecure(), grpc.WithUnaryInterceptor(otgrpc.OpenTracingClientInterceptor(tracer)))
		if err != nil {
			log.Fatalf("failed to dial: %v", err)
		}
		defer conn.Close()

		client := pinger.NewPingerClient(conn)
		res, err := client.PingPong(context.Background(), &pinger.PingPongRequest{
			Message: cfg.Message,
			Delay:   cfg.Delay,
		})
		if err != nil {
			return err
		}
		log.Println(res.Body)
		return nil
	})
}
