package main

import (
	"context"
	"log"
	"tag/idl/gen/tag"
	"tag/service"
)

type Handler struct {
	*tag.UnimplementedTagServiceServer
}

func (*Handler) CreateTag(ctx context.Context, req *tag.CreateTagReq) (*tag.CreateTagResp, error) {
	log.Printf("req:%+v", req)
	cCtx := service.NewCreateTagContext(ctx, req)
	service.NewCreateTagService().Create(cCtx)
	log.Printf("resp:%+v", cCtx.Resp)

	return cCtx.Resp, nil
}

func (*Handler) QueryMovieTag(ctx context.Context, req *tag.QueryMovieTagReq) (*tag.QueryMovieTagResp, error) {
	log.Printf("req:%+v", req)
	qCtx := service.NewQueryMovieTagContext(ctx, req)
	service.NewQueryMovieTagService().Query(qCtx)
	log.Printf("resp:%+v", qCtx.Resp)

	return qCtx.Resp, nil
}

func (*Handler) QueryUserTagCloud(ctx context.Context, req *tag.QueryUserTagCloudReq) (*tag.QueryUserTagCloudResp, error) {
	log.Printf("req:%+v", req)
	qCtx := service.NewQueryUserTagCloudContext(ctx, req)
	service.NewQueryUserTagCloudService().Query(qCtx)
	log.Printf("resp:%+v", qCtx.Resp)

	return qCtx.Resp, nil
}

func (*Handler) QueryTagRecords(ctx context.Context, req *tag.QueryTagRecordsReq) (*tag.QueryTagRecordsResp, error) {
	log.Printf("req:%+v", req)
	qCtx := service.NewQueryTagRecordsContext(ctx, req)
	service.NewQueryTagRecordsService().Query(qCtx)
	log.Printf("resp:%+v", qCtx.Resp)

	return qCtx.Resp, nil
}

func (*Handler) QueryMovieTopNTags(ctx context.Context, req *tag.QueryMovieTopNTagsReq) (*tag.QueryMovieTopNTagsResp, error) {
	log.Printf("req:%+v", req)
	qCtx := service.NewQueryMovieTopNTagsContext(ctx, req)
	service.NewQueryMovieTopNTagsService().Query(qCtx)
	log.Printf("resp:%+v", qCtx.Resp)

	return qCtx.Resp, nil
}
