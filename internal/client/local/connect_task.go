package local

import (
	job "github.com/AgentCoop/go-work"
	"github.com/go-serv/service/pkg/z"
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
		addr := "unix-abstract://" + c.Endpoint().Address()
		//addr := "@/tmp/." + c.Endpoint().Address()
		conn, err := grpc.Dial(addr, c.DialOptions()...)
		task.Assert(err)
		v.(z.ClientInterface).NewClient(conn)
		task.Done()
	}
	return init, run, nil
}
