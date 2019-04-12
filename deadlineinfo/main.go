package main

import (
	_ "net/http/pprof"

	"github.com/lileio/fromenv"
	"github.com/lileio/lile"
	"github.com/lileio/logr"
	"github.com/lileio/pubsub"
	"github.com/lileio/pubsub/middleware/defaults"
	"github.com/lupinthe14th/deadlineinfo"
	"github.com/lupinthe14th/deadlineinfo/deadlineinfo/cmd"
	"github.com/lupinthe14th/deadlineinfo/server"
	"google.golang.org/grpc"
)

func main() {
	logr.SetLevelFromEnv()
	s := &server.DeadlineinfoServer{}

	lile.Name("deadlineinfo")
	lile.Server(func(g *grpc.Server) {
		deadlineinfo.RegisterDeadlineinfoServer(g, s)
	})

	pubsub.SetClient(&pubsub.Client{
		ServiceName: lile.GlobalService().Name,
		Provider:    fromenv.PubSubProvider(),
		Middleware:  defaults.Middleware,
	})

	cmd.Execute()
}
