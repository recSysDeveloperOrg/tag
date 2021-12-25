package service

import (
	"context"
	"strings"
	"sync"
	. "tag/constant"
	"tag/idl/gen/tag"
	"tag/model"
)

type QueryUserTagCloudService struct {
}

type QueryUserTagCloudContext struct {
	Ctx     context.Context
	Req     *tag.QueryUserTagCloudReq
	Resp    *tag.QueryUserTagCloudResp
	ErrCode *ErrorCode

	Tags []*model.TagUser
}

var queryUserTagCloudService *QueryUserTagCloudService
var queryUserTagCloudServiceOnce sync.Once

func NewQueryUserTagCloudService() *QueryUserTagCloudService {
	queryUserTagCloudServiceOnce.Do(func() {
		queryUserTagCloudService = &QueryUserTagCloudService{}
	})

	return queryUserTagCloudService
}

func NewQueryUserTagCloudContext(ctx context.Context, req *tag.QueryUserTagCloudReq) *QueryUserTagCloudContext {
	return &QueryUserTagCloudContext{
		Ctx: ctx,
		Req: req,
		Resp: &tag.QueryUserTagCloudResp{
			BaseResp: &tag.BaseResp{},
		},
	}
}

func (s *QueryUserTagCloudService) Query(ctx *QueryUserTagCloudContext) {
	defer s.buildResponse(ctx)
	if s.checkParams(ctx); ctx.ErrCode != nil {
		return
	}
	s.query(ctx)
}

func (*QueryUserTagCloudService) checkParams(ctx *QueryUserTagCloudContext) {
	if strings.TrimSpace(ctx.Req.UserId) == "" {
		ctx.ErrCode = BuildErrCode("empty user id", RetParamsErr)
		return
	}
	if ctx.Req.NTags < 0 {
		ctx.ErrCode = BuildErrCode("cannot query negative number of tags", RetParamsErr)
	}
}

func (*QueryUserTagCloudService) query(ctx *QueryUserTagCloudContext) {
	tagUsers, err := model.NewTagDao().QueryTagUsersSortByUseTimes(ctx.Ctx, ctx.Req.UserId, ctx.Req.NTags)
	if err != nil {
		ctx.ErrCode = BuildErrCode(err, RetReadRepoErr)
		return
	}
	ctx.Tags = tagUsers
}

func (*QueryUserTagCloudService) buildResponse(ctx *QueryUserTagCloudContext) {
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
}
