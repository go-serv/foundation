package client

import (
	job "github.com/AgentCoop/go-work"
	net_client "github.com/go-serv/service/internal/client/net"
	"github.com/go-serv/service/pkg/z"
)

func NewClient(e z.EndpointInterface) *client {
	c := new(client)
	c.NetworkClientInterface = net_client.NewClient(serviceName, e)
	return c
}

func NewClientJob(e z.EndpointInterface) job.JobInterface {
	c := NewClient(e)
	cj := job.NewJob(c)
	cj.AddOneshotTask(c.ConnectTask)
	return cj
}
