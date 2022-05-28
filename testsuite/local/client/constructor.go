package client

import (
	job "github.com/AgentCoop/go-work"
	i "github.com/go-serv/service/internal"
	"github.com/go-serv/service/internal/client/local"
)

func NewClient(e i.EndpointInterface) *client {
	c := new(client)
	c.LocalClientInterface = local.NewClient(serviceName, e)
	return c
}

func NewClientJob(e i.EndpointInterface) job.JobInterface {
	c := NewClient(e)
	cj := job.NewJob(c)
	cj.AddOneshotTask(c.ConnectTask)
	return cj
}
