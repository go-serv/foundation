package z

import job "github.com/AgentCoop/go-work"

type EventHandlerFn func(...any) bool

type AppInterface interface {
	Job() job.JobInterface
	FetchConfig()
	AddService(ServiceInterface)
	//Services() AppInterface
	Start()
	Stop(reason any)
}
