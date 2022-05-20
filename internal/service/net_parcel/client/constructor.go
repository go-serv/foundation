package client

import (
	job "github.com/AgentCoop/go-work"
	i "github.com/go-serv/service/internal"
	cc "github.com/go-serv/service/internal/client"
)

func NewClient(e i.EndpointInterface) *client {
	c := new(client)
	c.NetworkClientInterface = cc.NewNetClient(serviceName, e)
	return c
}

func NewClientJob(e i.EndpointInterface) job.JobInterface {
	c := NewClient(e)
	cj := job.NewJob(c)
	cj.AddOneshotTask(c.Client_ConnectTask)
	return cj
}
