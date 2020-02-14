package main

import (
	"log"
	"os"

	"github.com/meowdada/grpc-playground/examples/streaming-file/server"
	cli "github.com/urfave/cli/v2"
)

func main() {
	app := &cli.App {
		Name:   "File Server",
		Usage:  "Server that receive file from client through gRPC",
		Action: server.Run,
		Flags: []cli.Flag {
			&cli.StringFlag {
				Name:  "address",
				Value: "127.0.0.1:8888",
				Usage: "address to communicate via gRPC",
			},
		},
	}

	err := app.Run(os.Args)
	if err != nil {
		log.Fatal(err)
	}
}