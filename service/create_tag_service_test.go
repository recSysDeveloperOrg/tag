package service

import (
	"context"
	"github.com/stretchr/testify/assert"
	"tag/constant"
	"tag/idl/gen/tag"
	"testing"
)

func TestCreateTagService_Create(t *testing.T) {
	movieID := generateStr()
	ctx := NewCreateTagContext(context.Background(), &tag.CreateTagReq{
		UserId:     testUser,
		MovieId:    movieID,
		TagContent: "HelloWorld",
	})
	NewCreateTagService().Create(ctx)
	assert.Equal(t, constant.RetSuccess.Code, ctx.Resp.BaseResp.ErrNo)
	NewCreateTagService().Create(ctx)
	assert.Equal(t, constant.RetSuccess.Code, ctx.Resp.BaseResp.ErrNo)
}
