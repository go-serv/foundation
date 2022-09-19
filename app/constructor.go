package app

import (
	job "github.com/AgentCoop/go-work"
	"github.com/go-serv/foundation/internal/web/dashboard"
	"github.com/go-serv/foundation/pkg/z"
	"math/rand"
	"time"
)

func NewApp(wp z.WebProxyInterface) *app {
	rand.Seed(time.Now().UnixNano())
	baseApp := new(app)
	baseApp.wp = wp
	baseApp.mainJob = job.NewJob(nil)
	return baseApp
}

func NewDashboard() z.DashboardInterface {
	return dashboard.NewDashboard()
}
