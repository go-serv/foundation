package client

import (
	job "github.com/AgentCoop/go-work"
	i "github.com/go-serv/service/internal"
	net_client "github.com/go-serv/service/internal/client/net"
)

func NewClient(e i.EndpointInterface) *client {
	c := new(client)
	c.NetworkClientInterface = net_client.NewClient(serviceName, e)
	return c
}

func NewClientJob(e i.EndpointInterface) job.JobInterface {
	c := NewClient(e)
	cj := job.NewJob(c)
	cj.AddOneshotTask(c.ConnectTask)
	return cj
}
