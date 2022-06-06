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
		conn, err := grpc.Dial(c.Endpoint().Address(), c.DialOptions()...)
		task.Assert(err)
		v.(z.NetworkClientInterface).NewClient(conn)
		task.Done()
	}
	return init, run, nil
}
