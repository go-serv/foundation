package client

import (
	job "github.com/AgentCoop/go-work"
	net_client "github.com/mesh-master/foundation/client/net"
	"github.com/mesh-master/foundation/pkg/z"
)

func NewClient(e z.EndpointInterface) *client {
	c := new(client)
	c.NetworkClientInterface = net_client.NewClient(serviceName, e)
	return c
}

func WithNewFtpSessionTimeout(c *client, tmMs int) {
	c.FtpNewSessionOptions.TimeoutMs = tmMs
}

func WithNewFtpSessionLifetime(c *client, lifetime uint32) {
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
