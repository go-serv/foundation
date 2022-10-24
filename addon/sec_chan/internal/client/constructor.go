package client

import (
	job "github.com/AgentCoop/go-work"
	net_client "github.com/go-serv/foundation/internal/client/net"
	"github.com/go-serv/foundation/pkg/z"
)

func NewClient(ep z.EndpointInterface) (*client, job.JobInterface) {
	c := new(client)
	c.NetworkClientInterface = net_client.NewClient(serviceName, ep)

	// Set client to use the custom gs-proto-enc codec.
	//c.WithDialOption(grpc.WithDefaultCallOptions(
	//	grpc.ForceCodec(codec.NewCodec())))
	// Create a new client job.
	cj := job.NewJob(c)
	// Add a dial task.
	cj.AddOneshotTask(c.ConnectTask)
	return c, cj
}
