package main

import (
	"giao/src/tour/tag_service/proto"
	"giao/src/tour/tag_service/server"
	"google.golang.org/grpc"
	"google.golang.org/grpc/reflection"
	"log"
	"net"
)

func main() {
	s := grpc.NewServer()
	proto.RegisterTagServiceServer(s, server.NewTagServe())

	reflection.Register(s)

	lis, err := net.Listen("tcp", ":9991")
	if err != nil {
		log.Fatalf("net listen err:%s", err)
	}

	log.Println("serve start")
	err = s.Serve(lis)
	if err != nil {
		log.Fatalf("server.Serve err: %s", err)
	}
}
