package main

import (
	"context"
	"log"
	"net/http"
	"os"

	kingpin "github.com/alecthomas/kingpin"
	"github.com/iheanyi/go-tracing-example/rpc/pinger"
	"github.com/iheanyi/go-tracing-example/rpc/ponger"
)

func main() {
	app := kingpin.New("pingerctl", "Tool to interact with ping and pong services")
	pingCmd(app.Command("ping", "ping pinger service"))
	pongCmd(app.Command("pong", "ping ponger service"))
	pingPongCmd(app.Command("ping-pong", "ping both the pinger and ponger services"))

	_, err := app.Parse(os.Args[1:])
	if err != nil {
		log.Fatalf("failed to run: %v", err)
	}
}

func pingCmd(cmd *kingpin.CmdClause) {
	cfg := struct {
		Message string
	}{}
	cmd.Arg("message", "message to send to the service").StringVar(&cfg.Message)

	cmd.Action(func(*kingpin.ParseContext) error {
		client := pinger.NewPingerProtobufClient("http://localhost:8082", &http.Client{})
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

func pongCmd(cmd *kingpin.CmdClause) {
	cfg := struct {
		Message string
		Delay   int64
	}{}
	cmd.Arg("message", "message to send to the service").StringVar(&cfg.Message)
	cmd.Flag("delay", "simulate latency to endpoint").Int64Var(&cfg.Delay)

	cmd.Action(func(*kingpin.ParseContext) error {
		client := ponger.NewPongerProtobufClient("http://localhost:8083", &http.Client{})
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

func pingPongCmd(cmd *kingpin.CmdClause) {
	cfg := struct {
		Message string
		Delay   int64
	}{}
	cmd.Arg("message", "message to send to the service").StringVar(&cfg.Message)
	cmd.Flag("delay", "delay to simulate latency in the pong service").Int64Var(&cfg.Delay)

	cmd.Action(func(*kingpin.ParseContext) error {
		client := pinger.NewPingerProtobufClient("http://localhost:8082", &http.Client{})
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
