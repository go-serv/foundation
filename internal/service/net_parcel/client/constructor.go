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

func NewClientJob(endpoint z.EndpointInterface) job.JobInterface {
	c := NewClient(endpoint)
	cj := job.NewJob(c)
	cj.AddOneshotTask(c.ConnectTask)
	return cj
}
