package local

import (
	job "github.com/AgentCoop/go-work"
	i "github.com/go-serv/service/internal"
	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials/insecure"
)

func (c *localClient) ConnectTask(j job.JobInterface) (job.Init, job.Run, job.Finalize) {
	init := func(task job.TaskInterface) {
		creds := insecure.NewCredentials()
		c.WithDialOption(grpc.WithTransportCredentials(creds))
	}
	run := func(task job.TaskInterface) {
		v := j.GetValue()
		addr := "unix-abstract:///tmp/." + c.Endpoint().Address()
		conn, err := grpc.Dial(addr, c.DialOptions()...)
		task.Assert(err)
		v.(i.ClientInterface).NewClient(conn)
		task.Done()
	}
	return init, run, nil
}