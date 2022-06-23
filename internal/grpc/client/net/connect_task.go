package net

import (
	job "github.com/AgentCoop/go-work"
	"github.com/go-serv/service/pkg/z"
	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials/insecure"
)

func (c *netClient) ConnectTask(j job.JobInterface) (job.Init, job.Run, job.Finalize) {
	init := func(task job.TaskInterface) {
		if c.insecure {
			creds := insecure.NewCredentials()
			c.WithDialOption(grpc.WithTransportCredentials(creds))
		}
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
