package server

import (
	"context"
	"encoding/json"
	"giao/pkg/tour/tag_service/pkg"
	"giao/pkg/tour/tag_service/pkg/errcode"
	"giao/pkg/tour/tag_service/proto"
)

func NewTagServe() *TagServer {
	return &TagServer{}
}

type TagServer struct {
}

func (s *TagServer) GetTagList(ctx context.Context, r *proto.GetTagListRequest) (*proto.GetTagListReply, error) {
	api := pkg.NewAPI(pkg.Target)
	list, err := api.GetTagList(ctx, r.GetName())
	if err != nil {
		return nil, errcode.ToRPCError(errcode.ErrorTagListFail)
	}
	tagList := proto.GetTagListReply{}
	err = json.Unmarshal(list, &tagList)
	if err != nil {
		return nil, errcode.ToRPCError(errcode.Fail)
	}
	return &tagList, nil
}
