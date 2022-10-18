package main

import (
	"errors"
	"fmt"
	job "github.com/AgentCoop/go-work"
	sec_chan "github.com/go-serv/foundation/addon/sec_chan"
	"github.com/go-serv/foundation/service/net"
	src "go-server-tests-endpoints"
	"log"
)

const nonceLen = 12

var newSessTask = func(j job.JobInterface) (job.Init, job.Run, job.Finalize) {
	run := func(task job.TaskInterface) {
		cc := j.GetValue().(sec_chan.ClientInterface)
		_, nonce, err := cc.NewSession(3600, nonceLen)
		task.Assert(err)
		if len(nonce) != nonceLen {
			j.Cancel(errors.New(fmt.Sprintf("expected %d got %v", nonceLen, len(nonce))))
		}
		if cc.SessionId() == 0 {
			j.Cancel(errors.New("zero session ID"))
		}
		task.Done()
	}
	return nil, run, nil
}

func main() {
	var (
		cj job.JobInterface
	)
	unsafeEp := net.NewTcp4Endpoint(src.ServerAddr, src.UnsafePort)
	_, cj = sec_chan.NewSecChanClient(unsafeEp, nil)
	cj.AddTask(newSessTask)
	<-cj.Run()
	if _, err := cj.GetInterruptedBy(); err != nil {
		log.Fatalf("plain TCP call failed with %v", err)
	}
}
