package server

import (
	"context"
	"encoding/json"
	"giao/src/tour/tag_service/pkg"
	"giao/src/tour/tag_service/pkg/errcode"
	"giao/src/tour/tag_service/proto"
)

func NewTagServe() *TagServer {
	return &TagServer{}
}

type TagServer struct {
}

func (s *TagServer) GetTagList(ctx context.Context, r *proto.GetTagListRequest) (*proto.GetTagListReply, error) {
	api := pkg.NewAPI("http://127.0.0.1:9999")
	list, err := api.GetTagList(ctx, r.GetName())
	if err != nil {
		return nil, err
	}
	tagList := proto.GetTagListReply{}
	err = json.Unmarshal(list, &tagList)
	if err != nil {
		return nil, errcode.ToRPCError(errcode.Fail)
	}
	return &tagList, nil
}
