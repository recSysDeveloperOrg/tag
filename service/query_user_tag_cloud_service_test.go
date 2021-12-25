package service

import (
	"context"
	"github.com/stretchr/testify/assert"
	"tag/constant"
	"tag/idl/gen/tag"
	"testing"
)

func TestQueryUserTagCloudService_Query(t *testing.T) {
	ctx := NewQueryUserTagCloudContext(context.Background(), &tag.QueryUserTagCloudReq{
		UserId: testUser,
		NTags:  100,
	})
	NewQueryUserTagCloudService().Query(ctx)
	assert.Equal(t, constant.RetSuccess.Code, ctx.Resp.BaseResp.ErrNo)
	assert.NotZero(t, len(ctx.Resp.Tags))
	for _, tg := range ctx.Resp.Tags {
		t.Logf("%+v", tg)
	}
}
