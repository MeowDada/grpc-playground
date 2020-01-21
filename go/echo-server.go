package main

import (
	"log"
	"net"
	"time"

	pb "github.com/meowdada/grpc-playground/go/proto/echo"
	context "golang.org/x/net/context"
	"google.golang.org/grpc"
	"google.golang.org/grpc/reflection"
)

type EchoServer struct{}

func (e *EchoServer) Echo(ctx context.Context, req *pb.EchoRequest) (resp *pb.EchoReply, err error) {

	log.Printf("receive client request, client send:%s\n", req.Msg)
	return &pb.EchoReply{
		Msg: req.Msg,
		Unixtime: time.Now().Unix(),
	}, nil
}

func main() {
	apiListener, err := net.Listen("tcp", ":9999")
	if err != nil {
		log.Println(err)
		return
	}

	es := &EchoServer{}

	grpc := grpc.NewServer()
	pb.RegisterEchoServer(grpc, es)

	reflection.Register(grpc)
	if err := grpc.Serve(apiListener); err != nil {
		log.Fatal(" grpc.Server Error: ", err)
		return
	}
}