package service

import (
	"context"
	"github.com/stretchr/testify/assert"
	"tag/constant"
	"tag/idl/gen/tag"
	"testing"
)

func TestQueryTagRecordsService_Query(t *testing.T) {
	ctx := NewQueryTagRecordsContext(context.Background(), &tag.QueryTagRecordsReq{
		UserId: testUser,
		Page:   0,
		Offset: 100,
	})
	NewQueryTagRecordsService().Query(ctx)
	assert.Equal(t, constant.RetSuccess.Code, ctx.Resp.BaseResp.ErrNo)
	assert.NotZero(t, len(ctx.Resp.Tags))
	assert.NotZero(t, ctx.Resp.NRecords)
	for _, tg := range ctx.Resp.Tags {
		t.Logf("%+v", tg)
	}
}
