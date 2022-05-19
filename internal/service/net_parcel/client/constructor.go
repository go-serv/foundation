package client

import (
	job "github.com/AgentCoop/go-work"
	cc "github.com/go-serv/service/internal/client"
	"github.com/go-serv/service/internal/server"
)

func NewClient(e server.EndpointInterface) *client {
	c := new(client)
	c.NetworkClientInterface = cc.NewNetClient(e)
	return c
}

func NewClientJob(e server.EndpointInterface) job.JobInterface {
	c := NewClient(e)
	cj := job.NewJob(c)
	cj.AddOneshotTask(c.Client_ConnectTask)
	return cj
}
