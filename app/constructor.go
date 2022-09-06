package app

import (
	job "github.com/AgentCoop/go-work"
	"github.com/go-serv/foundation/internal/platform"
	"math/rand"
	"time"
)

func NewApp() *app {
	rand.Seed(time.Now().UnixNano())
	baseApp := new(app)
	baseApp.mainJob = job.NewJob(nil)
	baseApp.platform = platform.NewPlatform(0)
	return baseApp
}
