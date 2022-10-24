package net

import (
	job "github.com/AgentCoop/go-work"
	"github.com/go-serv/foundation/internal/runtime"
	"github.com/go-serv/foundation/pkg/z"
	"github.com/go-serv/foundation/pkg/z/event"
	"google.golang.org/grpc"
)

func (c *netClient) ConnectTask(j job.JobInterface) (job.Init, job.Run, job.Finalize) {
	init := func(task job.TaskInterface) {
		ep := c.Endpoint().(z.NetEndpointInterface)
		transCreds := ep.TransportCredentials()
		c.WithDialOption(grpc.WithTransportCredentials(transCreds))
		//if ep.TlsConfig() == nil { // Use codec providing message encryption.
		//	c.WithDialOption(grpc.WithDefaultCallOptions(grpc.ForceCodec(y.NewCodec())))
		//}
		runtime.Runtime().TriggerEvent(event.NetClientBeforeDial, c, ep.TlsConfig() != nil)
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
