package server

import (
	"log"
	"path/filepath"
	"io"
	"net"
	"os"

	"google.golang.org/grpc"
	pb "github.com/meowdada/grpc-playground/examples/streaming-file/pb"
	cli "github.com/urfave/cli/v2"
)

type Server struct {
	root string
}

func (s Server) Upload (stream pb.Uploader_UploadServer) error {

	var processedBytes int64 = 0
	var f *os.File

	defer func() {
		if f != nil {
			f.Close()
		}
	}()

	for {
		chunk, err := stream.Recv()
		if err == io.EOF {
			return stream.SendAndClose(&pb.UploadResponse {
				Bytes:   processedBytes,
				Success: true,
			})
		}
		if err != nil {
			log.Println(err)
			return stream.SendAndClose(&pb.UploadResponse {
				Bytes:   processedBytes,
				Success: false,
			})
		}
		if f == nil {
			f, err = os.Create(filepath.Join(s.root, chunk.Filename))
			if err != nil {
				log.Println(err)
				return err
			}
		}
		n, err := f.Write(chunk.Data)
		if err != nil {
			log.Println(err)
			return err
		}
		processedBytes += int64(n)
	}
}

func newServer(root string) Server {
	return Server{root}
}

func Run(c *cli.Context) error {

	address := c.String("address")
	root := c.String("root")

	if err := createRootdir(root); err != nil {
		log.Println(err)
		return err
	}

	lis, err := net.Listen("tcp", address)
	if err != nil {
		log.Println(err)
		return err
	}

	grpcServer := grpc.NewServer()
	pb.RegisterUploaderServer(grpcServer, newServer(root))
	grpcServer.Serve(lis)

	return nil
}

func createRootdir(root string) error {
	return os.MkdirAll(root, os.ModePerm)
}