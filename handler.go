package main

import (
	"context"
	"tag/idl/gen/tag"
)

type Handler struct {
	*tag.UnimplementedTagServiceServer
}

func (*Handler) CreateTag(ctx context.Context, req *tag.CreateTagReq) (*tag.CreateTagResp, error) {
	return nil, nil
}

func (*Handler) QueryMovieTag(ctx context.Context, req *tag.QueryMovieTagReq) (*tag.QueryMovieTagResp, error) {
	return nil, nil
}

func (*Handler) QueryUserTagCloud(ctx context.Context, req *tag.QueryUserTagCloudReq) (*tag.QueryUserTagCloudResp, error) {
	return nil, nil
}

func (*Handler) QueryTagRecords(ctx context.Context, req *tag.QueryTagRecordsReq) (*tag.QueryTagRecordsResp, error) {
	return nil, nil
}

func (*Handler) QueryMovieTopNTags(ctx context.Context, req *tag.QueryMovieTopNTagsReq) (*tag.QueryMovieTopNTagsResp, error) {
	return nil, nil
}
