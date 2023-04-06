package main

import (
	"context"
	"flag"
	"fmt"
	pb "giao/grpc_demo/proto"
	"google.golang.org/grpc"
	"net"
)

var port string

func init() {
	flag.StringVar(&port, "p", "9999", "启动端口号")
	flag.Parse()
}

type GreeterServer struct{}

func (s *GreeterServer) SayList(r *pb.HelloRequest, stream pb.Greeter_SayListServer) error {
	for n := 0; n <= 6; n++ {
		_ = stream.Send(&pb.HelloReply{Message: "hello.list"})
	}

	return nil
}

func (s *GreeterServer) SayHello(ctx context.Context, r *pb.HelloRequest) (*pb.HelloReply, error) {
	fmt.Println("server :", r.Name)
	return &pb.HelloReply{Message: "hello world"}, nil
}
func main() {
	server := grpc.NewServer()
	pb.RegisterGreeterServer(server, &GreeterServer{})
	listen, _ := net.Listen("tcp", ":"+port)
	err := server.Serve(listen)
	if err != nil {
		fmt.Println("Server err:", err.Error())
		return
	}
}
