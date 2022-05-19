package client

import (
	job "github.com/AgentCoop/go-work"
)

func (c *client) ConnectTask(j job.JobInterface) (job.Init, job.Run, job.Finalize) {
	run := func(task job.TaskInterface) {
		//var opts = []grpc.DialOption{
		//	grpc.WithInsecure(),
		//	grpc.WithChainUnaryInterceptor(c.unaryInterceptors...),
		//	grpc.WithDefaultCallOptions(grpc.CallContentSubtype(codec.Name)),
		//}
		//ctx := context.Background()
		//if c.timeoutMs > 0 {
		//	deadline := time.Now().Add(time.Duration(c.timeoutMs) * time.Millisecond)
		//	ctx, _ = context.WithDeadline(ctx, deadline)
		//}
		//target := c.target
		//switch {
		//case target[0] == '@':
		//	target = strings.Replace(target, "@", "unix-abstract:", 1)
		//case c.port > 0:
		//	target = target + ":" + strconv.Itoa(int(c.port))
		//}
		//conn, err := grpc.Dial(ctx, target, opts...)
		//task.Assert(err)
		//c.conn = conn
		//c.connProvider(conn)
		task.Done()
	}
	return nil, run, nil
}
