package client

import (
	job "github.com/AgentCoop/go-work"
	net_client "github.com/mesh-master/foundation/client/net"
	"github.com/mesh-master/foundation/pkg/z"
)

func NewClient(ep z.EndpointInterface) (*client, job.JobInterface) {
	c := new(client)
	c.NetworkClientInterface, _ = net_client.NewClient(serviceName, ep)

	cj := job.NewJob(c)
	cj.AddOneshotTask(c.ConnectTask)
	return c, cj
}
