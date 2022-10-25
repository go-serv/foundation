package client

import (
	job "github.com/AgentCoop/go-work"
	"github.com/go-serv/foundation/addon/sec_chan/internal/codec"
	net_client "github.com/go-serv/foundation/client/net"
	"github.com/go-serv/foundation/pkg/z"
	"google.golang.org/grpc"
)

func NewClient(ep z.EndpointInterface) (*client, job.JobInterface) {
	c := new(client)
	c.NetworkClientInterface = net_client.NewClient(serviceName, ep)

	// Set client to use the custom gs-proto-enc codec.
	c.WithDialOption(grpc.WithDefaultCallOptions(
		grpc.ForceCodec(codec.NewCodec())))
	// Create a new client job.
	cj := job.NewJob(c)
	// Add a dial task.
	cj.AddOneshotTask(c.ConnectTask)
	return c, cj
}
