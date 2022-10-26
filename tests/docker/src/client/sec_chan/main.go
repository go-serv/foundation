package main

import (
	"errors"
	"fmt"
	job "github.com/AgentCoop/go-work"
	"github.com/go-serv/foundation/addon/sec_chan"
	"github.com/go-serv/foundation/addon/sec_chan/x"
	"github.com/go-serv/foundation/pkg/ancillary/memoize"
	"github.com/go-serv/foundation/pkg/z"
	"github.com/go-serv/foundation/runtime"
	"github.com/go-serv/foundation/service/net"
	src "go-server-tests-endpoints"
	"log"
)

const nonceLen = 12

var createSessionTask = func(j job.JobInterface) (job.Init, job.Run, job.Finalize) {
	run := func(task job.TaskInterface) {
		cc := j.GetValue().(sec_chan.ClientInterface)
		keyExchAlgs := []x.KeyExchangeAlgoType{
			//x.KeyExchangePSK,
			x.KeyExchangeDH,
		}
		for _, algo := range keyExchAlgs {
			_, nonce, err := cc.Create(3600, nonceLen, algo)
			task.Assert(err)
			if len(nonce) != nonceLen {
				j.Cancel(errors.New(fmt.Sprintf("expected %d got %v", nonceLen, len(nonce))))
			}
			if cc.SessionId() == 0 {
				j.Cancel(errors.New("zero session ID"))
			}
			// Close the created session.
			err = cc.Close()
			cc.Reset()
			task.Assert(err)
		}
		task.Done()
	}
	return nil, run, nil
}

func main() {
	var (
		cj job.JobInterface
	)

	// Register PSK key resolver.
	runtime.RegisterResolver(x.PskResolverKey, memoize.NewMemoizer(func(a ...any) (any, error) {
		fmt.Printf("enc key %v", src.PskKey)
		return src.PskKey, nil
	}))

	runtime.RegisterResolver(z.ApiKeyResolver, memoize.NewMemoizer(func(a ...any) (any, error) {
		return src.ApiKey, nil
	}))

	unsafeEp := net.NewTcp4Endpoint(src.ServerAddr, src.UnsafePort)
	_, cj = sec_chan.NewSecChanClient(unsafeEp, nil)
	cj.AddTask(createSessionTask)
	<-cj.Run()

	if _, err := cj.GetInterruptedBy(); err != nil {
		log.Fatalf("plain TCP call failed with %v", err)
	}
}
