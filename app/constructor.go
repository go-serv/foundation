package app

import (
	job "github.com/AgentCoop/go-work"
	"github.com/mesh-master/foundation/internal/grpc/middleware"
	net_mw "github.com/mesh-master/foundation/internal/middleware/net"
	"github.com/mesh-master/foundation/internal/middleware/session"
	"github.com/mesh-master/foundation/internal/web/dashboard"
	mwPkg "github.com/mesh-master/foundation/pkg/z"
	"math/rand"
	"time"
)

func NewServerApp(wp mwPkg.WebProxyInterface) *netServer {
	rand.Seed(time.Now().UnixNano())
	srv := new(netServer)
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
