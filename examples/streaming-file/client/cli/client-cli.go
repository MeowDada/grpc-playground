package main

import (
	"log"
	"os"

	"github.com/meowdada/grpc-playground/examples/streaming-file/client"
	cli "github.com/urfave/cli/v2"
)

func main() {
	app := &cli.App {
		Name:   "File uploader",
		Usage:  "Upload file through gRPC",
		Action: client.Run,
		Flags: []cli.Flag {
			&cli.StringFlag {
				Name:  "address",
				Aliases: []string{"a"},
				Value: "127.0.0.1:8888",
				Usage: "address to communicate via gRPC",
			},
			&cli.StringFlag {
				Name:  "filename, ",
				Aliases: []string{"f"},
				Value: "",
				Usage: "path to the file to be uploaded",
			},
		},
	}

	err := app.Run(os.Args)
	if err != nil {
		log.Fatal(err)
	}
}