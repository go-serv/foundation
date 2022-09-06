package z

import job "github.com/AgentCoop/go-work"

type EventHandlerFn func(...any) bool

type AppInterface interface {
	Job() job.JobInterface
	Platform() PlatformInterface
	FetchConfig()
	//RegisterResolver(key any, resolver MemoizerInterface)
	//Resolve(key any, args ...any) (v any, err error)
	//RegisterEventHandler(eventType any, h EventHandlerFn)
	//TriggerEvent(event any, extraArgs ...any)
	AddService(ServiceInterface)
	//Services() AppInterface
	Start()
	Stop(reason any)
}
