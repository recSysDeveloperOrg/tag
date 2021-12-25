package service

import (
	"context"
	"github.com/stretchr/testify/assert"
	"tag/constant"
	"tag/idl/gen/tag"
	"testing"
)

func TestQueryMovieTagService_Query(t *testing.T) {
	ctx := NewQueryMovieTagContext(context.Background(), &tag.QueryMovieTagReq{
		UserId:  testUser,
		MovieId: "920002603369950176332508",
	})
	NewQueryMovieTagService().Query(ctx)
	assert.Equal(t, constant.RetSuccess.Code, ctx.Resp.BaseResp.ErrNo)
	assert.NotZero(t, len(ctx.Resp.Tags))
	for _, tg := range ctx.Resp.Tags {
		t.Logf("%+v", tg)
	}
}
