package main

import (
	"context"
	"fmt"
	"giao/pkg/tour/tag_service/proto"
	"google.golang.org/grpc"
)

func main() {
	ctx := context.Background()
	conn, _ := GetClientConn(ctx, "localhost:9991", nil)
	defer conn.Close()

	client := proto.NewTagServiceClient(conn)
	list, err := client.GetTagList(ctx, &proto.GetTagListRequest{Name: "gegegege"})
	if err != nil {
		fmt.Println(err)
		return
	}
	fmt.Println(list)
}

func GetClientConn(ctx context.Context, target string, ops []grpc.DialOption) (*grpc.ClientConn, error) {
	ops = append(ops, grpc.WithInsecure())
	return grpc.DialContext(ctx, target, ops...)
}
