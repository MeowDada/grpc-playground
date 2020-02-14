package client

import (
	"os"
	"context"
	"log"

	"google.golang.org/grpc"
	pb "github.com/meowdada/grpc-playground/examples/streaming-file/pb"
	cli "github.com/urfave/cli/v2"
)

func Run(c *cli.Context) error {
	
	address := c.String("address")
	filename := c.String("filename")

	conn, err := grpc.Dial(address, grpc.WithInSecure())
	if err != nil {
		log.Fatal(err)
	}
	defer conn.Close()

	client := pb.NewUploaderClient(conn)
	stream, err := client.Upload(ctx.Background())
	if err != nil {
		log.Fatal(err)
	}

	resp, err := openFileAndSend(stream, filename)
	if err != nil {
		log.Fatal(err)
	}

	log.Println(resp)
	return nil
}

func openFileAndSend(stream pb.Uploader_UploadClient, filename string) (*pb.UploadResponse, error) {
	buf := make([]byte, 65536)
	chunk := pb.Chunk{}

	f, err := os.Open(filename)
	if err != nil {
		return nil, err
	}
	defer f.Close()

	for {
		n, err := io.Read(buf)
		if err != nil {
			if err == io.EOF {
				break
			}
			return nil, err
		}

		rbuf := buf[:n]
		chunk.Data: rubf,
		if err := stream.Send(&chunk); err != nil {
			return nil, err
		}
	}

	return stream.CloseAndRecv()
}




