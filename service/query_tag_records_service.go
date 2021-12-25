package service

import (
	"context"
	"strings"
	"sync"
	. "tag/constant"
	"tag/idl/gen/tag"
	"tag/model"
)

type QueryTagRecordsService struct {
}

type QueryTagRecordsContext struct {
	Ctx     context.Context
	Req     *tag.QueryTagRecordsReq
	Resp    *tag.QueryTagRecordsResp
	ErrCode *ErrorCode

	Tags  []*model.TagUser
	NTags int64
}

var queryTagRecordsService *QueryTagRecordsService
var queryTagRecordsServiceOnce sync.Once

func NewQueryTagRecordsService() *QueryTagRecordsService {
	queryTagRecordsServiceOnce.Do(func() {
		queryTagRecordsService = &QueryTagRecordsService{}
	})

	return queryTagRecordsService
}

func NewQueryTagRecordsContext(ctx context.Context, req *tag.QueryTagRecordsReq) *QueryTagRecordsContext {
	return &QueryTagRecordsContext{
		Ctx: ctx,
		Req: req,
		Resp: &tag.QueryTagRecordsResp{
			BaseResp: &tag.BaseResp{},
		},
	}
}

func (s *QueryTagRecordsService) Query(ctx *QueryTagRecordsContext) {
	defer s.buildResponse(ctx)
	if s.checkParams(ctx); ctx.ErrCode != nil {
		return
	}
	s.query(ctx)
}

func (*QueryTagRecordsService) checkParams(ctx *QueryTagRecordsContext) {
	if strings.TrimSpace(ctx.Req.UserId) == "" {
		ctx.ErrCode = BuildErrCode("empty user id", RetParamsErr)
		return
	}
	if ctx.Req.Page < 0 {
		ctx.ErrCode = BuildErrCode("negative page", RetParamsErr)
		return
	}
	if ctx.Req.Offset < 0 {
		ctx.ErrCode = BuildErrCode("negative offset", RetParamsErr)
	}
}

func (*QueryTagRecordsService) query(ctx *QueryTagRecordsContext) {
	records, err := model.NewTagDao().QueryTagUsersSortByTime(ctx.Ctx, ctx.Req.UserId, ctx.Req.Page, ctx.Req.Offset)
	if err != nil {
		ctx.ErrCode = BuildErrCode(err, RetReadRepoErr)
		return
	}
	ctx.Tags = records

	nTags, err := model.NewTagDao().CountTagUsersByUserID(ctx.Ctx, ctx.Req.UserId)
	if err != nil {
		ctx.ErrCode = BuildErrCode(err, RetReadRepoErr)
		return
	}
	ctx.NTags = nTags
}

func (*QueryTagRecordsService) buildResponse(ctx *QueryTagRecordsContext) {
	errCode := RetSuccess
	if ctx.ErrCode != nil {
		errCode = ctx.ErrCode
	}

	ctx.Resp.BaseResp.ErrNo, ctx.Resp.BaseResp.ErrMsg = errCode.Code, errCode.Msg
	for _, tagUser := range ctx.Tags {
		ctx.Resp.Tags = append(ctx.Resp.Tags, &tag.Tag{
			Id:        tagUser.TagID,
			Content:   tagUser.Content,
			UpdatedAt: tagUser.UpdatedAt,
			NTimes:    int64(len(tagUser.MovieIDs)),
		})
	}
	ctx.Resp.NRecords = ctx.NTags
}
