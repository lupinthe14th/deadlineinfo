package server

import (
	"testing"

	"context"
	"github.com/lupinthe14th/deadlineinfo"
	"github.com/stretchr/testify/assert"
)

func TestGet(t *testing.T) {
	ctx := context.Background()
	req := &deadlineinfo.GetRequest{}

	res, err := cli.Get(ctx, req)
	assert.Nil(t, err)
	assert.NotNil(t, res)
}
