package main

import (
	job "github.com/AgentCoop/go-work"
	"github.com/go-serv/foundation/app/net_parcel/client"
	"github.com/go-serv/foundation/pkg/y/netparcel"
	"github.com/go-serv/foundation/service/net"
	src "go-server-tests-endpoints"
	"log"
)

const primeNum = 7

var pingTask = func(j job.JobInterface) (job.Init, job.Run, job.Finalize) {
	run := func(task job.TaskInterface) {
		cc := j.GetValue().(netparcel.NetParcelClientInterface)
		//pingOpts := client.PingOptions{}
		//pingOpts.TimeoutMs = 2000
		//cc.WithOptions(pingOpts)
		res, err := cc.Ping(primeNum)
		task.Assert(err)
		if res != primeNum {
			log.Fatalf("plain TCP call, expected %d got %v", primeNum, res)
		}
		task.Done()
	}
	return nil, run, nil
}

func main() {
	unsafeEp := net.NewTcp4Endpoint(src.ServerAddr, src.UnsafePort)
	j := client.NewClientJob(unsafeEp)
	j.AddTask(pingTask)
	<-j.Run()
	if _, err := j.GetInterruptedBy(); err != nil {
		log.Fatalf("plain TCP call failed with %v", err)
	}
}
