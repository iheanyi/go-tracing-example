package main

import (
	"context"
	"fmt"
	"net/http"
	"os"

	"github.com/iheanyi/go-tracing-example/rpc/pinger"
)

func main() {
	client := pinger.NewPingerProtobufClient("http://localhost:8082", &http.Client{})
	msg := ""
	if len(os.Args) > 1 {
		msg = os.Args[1]
	}

	res, err := client.Ping(context.Background(), &pinger.PingRequest{
		Message: msg,
	})
	if err != nil {
		fmt.Printf("oh no: %v", err)
		os.Exit(1)
	}

	fmt.Printf("Response successfully received: %v", res.Body)
}
