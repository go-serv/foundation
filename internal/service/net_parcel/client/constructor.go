package client

import (
	job "github.com/AgentCoop/go-work"
	net_client "github.com/go-serv/service/internal/client/net"
	"github.com/go-serv/service/pkg/z"
)

func NewClient(e z.EndpointInterface) *client {
	c := new(client)
	c.NetworkClientInterface = net_client.NewClient(serviceName, e)
	c.FtpNewSessionOptions.TimeoutMs = 3000
	return c
}

func WithFtpSessionLifetime(c *client, lifetime uint32) {
	lv := new(uint32)
	*lv = lifetime
	c.FtpNewSessionOptions.Lifetime = lv
}

func NewClientJob(endpoint z.EndpointInterface) job.JobInterface {
	c := NewClient(endpoint)
	cj := job.NewJob(c)
	cj.AddOneshotTask(c.ConnectTask)
	return cj
}
