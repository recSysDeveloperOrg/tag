package service

import (
	"context"
	"strings"
	"sync"
	. "tag/constant"
	"tag/idl/gen/tag"
	"tag/model"
)

type QueryMovieTagService struct {
}

type QueryMovieTagContext struct {
	Ctx     context.Context
	Req     *tag.QueryMovieTagReq
	Resp    *tag.QueryMovieTagResp
	ErrCode *ErrorCode

	Tags []*model.TagUser
}

var queryMovieTagService *QueryMovieTagService
var queryMovieTagServiceOnce sync.Once

func NewQueryMovieTagService() *QueryMovieTagService {
	queryMovieTagServiceOnce.Do(func() {
		queryMovieTagService = &QueryMovieTagService{}
	})

	return queryMovieTagService
}

func NewQueryMovieTagContext(ctx context.Context, req *tag.QueryMovieTagReq) *QueryMovieTagContext {
	return &QueryMovieTagContext{
		Ctx: ctx,
		Req: req,
		Resp: &tag.QueryMovieTagResp{
			BaseResp: &tag.BaseResp{},
		},
	}
}

func (s *QueryMovieTagService) Query(ctx *QueryMovieTagContext) {
	defer s.buildResponse(ctx)
	if s.checkParams(ctx); ctx.ErrCode != nil {
		return
	}
	s.query(ctx)
}

func (*QueryMovieTagService) checkParams(ctx *QueryMovieTagContext) {
	if strings.TrimSpace(ctx.Req.MovieId) == "" {
		ctx.ErrCode = BuildErrCode("empty movie id", RetParamsErr)
		return
	}
	if strings.TrimSpace(ctx.Req.UserId) == "" {
		ctx.ErrCode = BuildErrCode("empty user id", RetParamsErr)
	}
}

func (*QueryMovieTagService) query(ctx *QueryMovieTagContext) {
	tags, err := model.NewTagDao().QueryMovieTag(ctx.Ctx, ctx.Req.UserId, ctx.Req.MovieId)
	if err != nil {
		ctx.ErrCode = BuildErrCode(err, RetReadRepoErr)
		return
	}
	ctx.Tags = tags
}

func (*QueryMovieTagService) buildResponse(ctx *QueryMovieTagContext) {
	errCode := RetSuccess
	if ctx.ErrCode != nil {
		errCode = ctx.ErrCode
	}

	ctx.Resp.BaseResp.ErrNo, ctx.Resp.BaseResp.ErrMsg = errCode.Code, errCode.Msg
	for _, t := range ctx.Tags {
		ctx.Resp.Tags = append(ctx.Resp.Tags, &tag.Tag{
			Id:        t.TagID,
			Content:   t.Content,
			UpdatedAt: t.UpdatedAt,
			NTimes:    int64(len(t.MovieIDs)),
		})
	}
}
