package server

import (
	"io"
	"net"

	"google.golang.org/grpc"
	pb "github.com/meowdada/grpc-playground/examples/streaming-file/pb"
	cli "github.com/urfave/cli/v2"
)

type Server struct {}

func (s Server) Upload (stream pb.Uploader_UploadServer) error {

	var processedBytes int64 = 0

	for {
		chunk, err := stream.Recv()
		if err == io.EOF {
			return stream.SendAndClose(&pb.UploadResponse {
				Bytes:   processedBytes,
				Success: true,
			})
		}
		if err != nil {
			return stream.SendAndClose(&pb.UploadResponse {
				Bytes:   processedBytes,
				Success: false,
			})
		}
		processedBytes += int64(len(chunk.Data))
	}
}

func newServer() Server {
	return Server{}
}

func Run(c *cli.Context) error {

	address := c.String("address")
	root := c.String("root")

	if err := createRootdir(root); err != nil {
		return err
	}

	lis, err := net.Listen("tcp", address)
	if err != nil {
		return err
	}

	grpcServer := grpc.NewServer()
	pb.RegisterUploaderServer(grpcServer, newServer())
	grpcServer.Serve(lis)

	return nil
}

func createRootdir(root string) error {
	return os.MkdirAll(root, os.ModePerm)
}