package app

import (
	job "github.com/AgentCoop/go-work"
	"github.com/go-serv/foundation/internal/grpc/middleware"
	net_mw "github.com/go-serv/foundation/internal/middleware/net"
	"github.com/go-serv/foundation/internal/middleware/session"
	"github.com/go-serv/foundation/internal/web/dashboard"
	mwPkg "github.com/go-serv/foundation/pkg/z"
	"math/rand"
	"time"
)

func NewServerApp(wp mwPkg.WebProxyInterface) *server {
	rand.Seed(time.Now().UnixNano())
	srv := new(server)
	srv.wp = wp
	srv.mainJob = job.NewJob(nil)
	srv.middleware = middleware.NewServerMiddleware()
	srv.middleware.Append(mwPkg.SessionMwKey, session.ServerRequestSessionHandler, session.ServerResponseSessionHandler)
	srv.middleware.Append(mwPkg.NetworkMwKey, net_mw.ServerRequestNetHandler, net_mw.ServerResponseNetHandler)
	return srv
}

func NewDashboard() mwPkg.DashboardInterface {
	return dashboard.NewDashboard()
}
