package client

import (
	job "github.com/AgentCoop/go-work"
	i "github.com/go-serv/service/internal"
	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials/insecure"
)

func (c *client) ConnectTask(j job.JobInterface) (job.Init, job.Run, job.Finalize) {
	init := func(task job.TaskInterface) {
		if c.insecure {
			creds := insecure.NewCredentials()
			c.dialOpts = append(c.dialOpts, grpc.WithTransportCredentials(creds))
		}
	}
	run := func(task job.TaskInterface) {
		v := j.GetValue()
		conn, err := grpc.Dial(c.Endpoint().Address(), c.dialOpts...)
		task.Assert(err)
		v.(i.NetworkClientInterface).NewClient(conn)
		task.Done()
	}
	return init, run, nil
}
