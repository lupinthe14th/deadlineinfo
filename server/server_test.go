package server

import (
	"os"
	"testing"

	"google.golang.org/grpc"

	"github.com/lileio/lile"
	"github.com/lupinthe14th/deadlineinfo"
)

var s = DeadlineinfoServer{}
var cli deadlineinfo.DeadlineinfoClient

func TestMain(m *testing.M) {
	impl := func(g *grpc.Server) {
		deadlineinfo.RegisterDeadlineinfoServer(g, s)
	}

	gs := grpc.NewServer()
	impl(gs)

	addr, serve := lile.NewTestServer(gs)
	go serve()

	cli = deadlineinfo.NewDeadlineinfoClient(lile.TestConn(addr))

	os.Exit(m.Run())
}
