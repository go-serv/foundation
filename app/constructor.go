package app

import (
	"fmt"
	job "github.com/AgentCoop/go-work"
	"github.com/go-serv/foundation/internal/grpc/middleware"
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
	fmt.Println("append session")
	srv.middleware.Append(mwPkg.SessionMwKey, session.ServerRequestSessionHandler, nil)
	return srv
}

func NewDashboard() mwPkg.DashboardInterface {
	return dashboard.NewDashboard()
}
