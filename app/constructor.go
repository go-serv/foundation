package app

import (
	job "github.com/AgentCoop/go-work"
	"math/rand"
	"time"
)

func NewApp() *app {
	rand.Seed(time.Now().UnixNano())
	baseApp := new(app)
	baseApp.mainJob = job.NewJob(nil)
	return baseApp
}
