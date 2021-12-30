package service

import (
	"context"
	"strings"
	"sync"
	. "tag/constant"
	"tag/idl/gen/tag"
	"tag/model"
)

type QueryMovieTopNTagsService struct {
}

type QueryMovieTopNTagsContext struct {
	Ctx     context.Context
	Req     *tag.QueryMovieTopNTagsReq
	Resp    *tag.QueryMovieTopNTagsResp
	ErrCode *ErrorCode

	Tags []*model.TagMovie
}

var queryMovieTopNTagsService *QueryMovieTopNTagsService
var queryMovieTopNTagsServiceOnce sync.Once

func NewQueryMovieTopNTagsService() *QueryMovieTopNTagsService {
	queryMovieTopNTagsServiceOnce.Do(func() {
		queryMovieTopNTagsService = &QueryMovieTopNTagsService{}
	})

	return queryMovieTopNTagsService
}

func NewQueryMovieTopNTagsContext(ctx context.Context, req *tag.QueryMovieTopNTagsReq) *QueryMovieTopNTagsContext {
	return &QueryMovieTopNTagsContext{
		Ctx: ctx,
		Req: req,
		Resp: &tag.QueryMovieTopNTagsResp{
			BaseResp: &tag.BaseResp{},
		},
	}
}

func (s *QueryMovieTopNTagsService) Query(ctx *QueryMovieTopNTagsContext) {
	defer s.buildResponse(ctx)
	if s.checkParams(ctx); ctx.ErrCode != nil {
		return
	}
	s.query(ctx)
}

func (*QueryMovieTopNTagsService) checkParams(ctx *QueryMovieTopNTagsContext) {
	if strings.TrimSpace(ctx.Req.MovieId) == "" {
		ctx.ErrCode = BuildErrCode("empty movie id", RetParamsErr)
		return
	}
	if ctx.Req.NTags < 0 {
		ctx.ErrCode = BuildErrCode("negative tag number", RetParamsErr)
	}
}

func (*QueryMovieTopNTagsService) query(ctx *QueryMovieTopNTagsContext) {
	tagMovies, err := model.NewTagDao().QueryTopKTagMovies(ctx.Ctx, ctx.Req.MovieId, ctx.Req.NTags)
	if err != nil {
		ctx.ErrCode = BuildErrCode(err, RetReadRepoErr)
		return
	}
	ctx.Tags = tagMovies
}

func (*QueryMovieTopNTagsService) buildResponse(ctx *QueryMovieTopNTagsContext) {
	errCode := RetSuccess
	if ctx.ErrCode != nil {
		errCode = ctx.ErrCode
	}

	ctx.Resp.BaseResp.ErrNo, ctx.Resp.BaseResp.ErrMsg = errCode.Code, errCode.Msg
	for _, tagMovie := range ctx.Tags {
		ctx.Resp.Tags = append(ctx.Resp.Tags, &tag.Tag{
			Id:        tagMovie.TagID,
			Content:   tagMovie.Content,
			UpdatedAt: tagMovie.UpdatedAt,
			NTimes:    tagMovie.TaggedTimes,
		})
	}
}
