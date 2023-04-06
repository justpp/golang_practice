package main

import (
	"context"
	"flag"
	"fmt"
	pb "giao/grpc_demo/proto"
	"google.golang.org/grpc"
	"io"
	"log"
)

var port string

func init() {
	flag.StringVar(&port, "p", "9999", "启动端口号")
	flag.Parse()
}

func SayHello(client pb.GreeterClient) error {
	resp, _ := client.SayHello(context.Background(), &pb.HelloRequest{Name: "ei hei"})
	log.Printf("clent say hello :%s", resp.Message)
	return nil
}

func SayList(client pb.GreeterClient, r *pb.HelloRequest) error {
	stream, _ := client.SayList(context.Background(), r)
	for {
		resp, err := stream.Recv()
		if err == io.EOF {
			break
		}
		if err != nil {
			return err
		}

		log.Printf("resp: %v", resp)
	}

	return nil
}

func main() {
	conn, _ := grpc.Dial(":"+port, grpc.WithInsecure())
	defer conn.Close()

	client := pb.NewGreeterClient(conn)
	err := SayHello(client)
	if err != nil {
		fmt.Println(err)
	}
	_ = SayList(client, &pb.HelloRequest{Name: "ei hhhhh"})
}
