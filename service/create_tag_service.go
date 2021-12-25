package service

import (
	"context"
	"errors"
	"strings"
	"sync"
	. "tag/constant"
	"tag/idl/gen/tag"
	"tag/model"
)

type CreateTagService struct {
}

type CreateTagContext struct {
	Ctx     context.Context
	Req     *tag.CreateTagReq
	Resp    *tag.CreateTagResp
	ErrCode *ErrorCode
}

var createTagService *CreateTagService
var createTagServiceOnce sync.Once

func NewCreateTagService() *CreateTagService {
	createTagServiceOnce.Do(func() {
		createTagService = &CreateTagService{}
	})

	return createTagService
}

func NewCreateTagContext(ctx context.Context, req *tag.CreateTagReq) *CreateTagContext {
	return &CreateTagContext{
		Ctx: ctx,
		Req: req,
		Resp: &tag.CreateTagResp{
			BaseResp: &tag.BaseResp{},
		},
	}
}

func (s *CreateTagService) Create(ctx *CreateTagContext) {
	defer s.buildResponse(ctx)
	if s.checkParams(ctx); ctx.ErrCode != nil {
		return
	}
	s.create(ctx)
}

func (*CreateTagService) checkParams(ctx *CreateTagContext) {
	if strings.TrimSpace(ctx.Req.UserId) == "" {
		ctx.ErrCode = BuildErrCode("empty user id", RetParamsErr)
		return
	}
	if strings.TrimSpace(ctx.Req.MovieId) == "" {
		ctx.ErrCode = BuildErrCode("empty movie id", RetParamsErr)
		return
	}
	if strings.TrimSpace(ctx.Req.TagContent) == "" {
		ctx.ErrCode = BuildErrCode("empty tag content", RetParamsErr)
	}
}

func (*CreateTagService) create(ctx *CreateTagContext) {
	if err := model.NewTagDao().InsertTag(ctx.Ctx, ctx.Req.TagContent); err != nil {
		// 非tag重复的错误直接返回
		if !errors.Is(err, model.ErrDuplicateTag) {
			ctx.ErrCode = BuildErrCode(err, RetWriteRepoErr)
			return
		}
	}

	tagID, err := model.NewTagDao().QueryTagID(ctx.Ctx, ctx.Req.TagContent)
	if err != nil {
		ctx.ErrCode = BuildErrCode(err, RetReadRepoErr)
		return
	}
	if tagID == nil {
		ctx.ErrCode = BuildErrCode("no corresponding tag id", RetSysErr)
		return
	}

	shouldIncTagTimes, err := model.NewTagDao().InsertUserTags(ctx.Ctx, ctx.Req.UserId, *tagID, ctx.Req.MovieId)
	if err != nil {
		ctx.ErrCode = BuildErrCode(err, RetWriteRepoErr)
		return
	}
	if shouldIncTagTimes {
		if err := model.NewTagDao().IncMovieTag(ctx.Ctx, ctx.Req.MovieId, *tagID); err != nil {
			ctx.ErrCode = BuildErrCode(err, RetWriteRepoErr)
		}
	}
}

func (*CreateTagService) buildResponse(ctx *CreateTagContext) {
	errCode := RetSuccess
	if ctx.ErrCode != nil {
		errCode = ctx.ErrCode
	}

	ctx.Resp.BaseResp.ErrNo, ctx.Resp.BaseResp.ErrMsg = errCode.Code, errCode.Msg
}
