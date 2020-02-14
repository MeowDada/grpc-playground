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
		Flag: []cli.Flag {
			&cli.StringFlag {
				Name:  "address",
				Value: "",
				Usage: "address to communicate via gRPC",
			},
			&cli.StringFlag {
				Name:  "filename",
				Value: "",
				Usage: "path to the file to be uploaded",
			},
		}
	}

	err := app.Run(os.Args)
	if err != nil {
		log.Fatal(err)
	}
}