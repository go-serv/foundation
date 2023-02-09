package net

import (
	job "github.com/AgentCoop/go-work"
	"github.com/mesh-master/foundation/internal/runtime"
	"github.com/mesh-master/foundation/pkg/z"
	"github.com/mesh-master/foundation/pkg/z/event"
	"google.golang.org/grpc"
)

func (c *netClient) ConnectTask(j job.JobInterface) (job.Init, job.Run, job.Finalize) {
	init := func(task job.TaskInterface) {
		ep := c.Endpoint().(z.NetEndpointInterface)
		transCreds := ep.TransportCredentials()
		c.WithDialOption(grpc.WithTransportCredentials(transCreds))
		tlsEnabled := ep.TlsConfig() != nil
		args := event.NetClientBeforeDialArgs{Client: c, TlsEnabled: tlsEnabled}
		runtime.Runtime().TriggerEvent(event.NetClientBeforeDial, args)
	}
	run := func(task job.TaskInterface) {
		v := j.GetValue()
		opts := append(c.DialOptions(), grpc.WithChainUnaryInterceptor(c.Middleware().UnaryClientInterceptor()))
		conn, err := grpc.Dial(c.Endpoint().Address(), opts...)
		task.Assert(err)
		v.(z.NetworkClientInterface).OnConnect(conn)
		task.Done()
	}
	return init, run, nil
}
