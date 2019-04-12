package server

import (
	"context"

	"github.com/lupinthe14th/deadlineinfo"
)

func (s DeadlineinfoServer) Get(ctx context.Context, r *deadlineinfo.GetRequest) (*deadlineinfo.GetResponse, error) {
	return &deadlineinfo.GetResponse{Date: r.GetDate()}, nil
}
