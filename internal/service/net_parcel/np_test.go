package server_test

import (
	"context"
	job "github.com/AgentCoop/go-work"
	"github.com/go-serv/service/internal/autogen/proto/net"
	srv "github.com/go-serv/service/internal/server"
	net_srv "github.com/go-serv/service/internal/server/net"
	np_client "github.com/go-serv/service/internal/service/net_parcel/client"
	"testing"
	"time"
)

const (
	testHost = "localhost"
	testPort = 4411
)

type testFixtures struct {
	t         *testing.T
	endpoint  srv.EndpointInterface
	srv       srv.NetworkServerInterface
	clientJob job.JobInterface
}

func setup(t *testing.T) *testFixtures {
	tf := new(testFixtures)
	tf.t = t
	tf.endpoint = srv.NewTcp4Endpoint(testHost, testPort)
	tf.srv = net_srv.NewNetServer()
	tf.srv.AddEndpoint(tf.endpoint)
	tf.clientJob = np_client.NewClientJob(tf.endpoint)
	return tf
}

func (tf *testFixtures) serverStatus() {
	_, err := tf.srv.MainJob().GetInterruptedBy()
	if err != nil {
		tf.t.Fatalf("server failed with %v\n", err)
	}
}

func (tf *testFixtures) finalize() {
	_, err := tf.clientJob.GetInterruptedBy()
	if err != nil {
		tf.t.Fatalf("client job failed with %v\n", err)
	}
}

func TestCryptoNonce(t *testing.T) {
	tf := setup(t)
	getNonceTask := func(j job.JobInterface) (job.Init, job.Run, job.Finalize) {
		const nonceLen = 32
		run := func(task job.TaskInterface) {
			cc := j.GetValue().(net.NetParcelClient)
			req := &net.CryptoNonce_Request{}
			req.Len = nonceLen
			res, err := cc.GetCryptoNonce(context.Background(), req)
			task.Assert(err)
			if len(res.GetNonce()) != nonceLen {
				t.Fatalf("expected nonce length %d, got %d", nonceLen, len(res.GetNonce()))
			}
			task.Done()
		}
		return nil, run, nil
	}
	go func() {
		tf.srv.Start()
	}()
	defer func() {
		tf.finalize()
	}()
	time.Sleep(time.Millisecond * 10)
	tf.serverStatus()
	tf.clientJob.AddTask(getNonceTask)
	<-tf.clientJob.Run()
	tf.srv.Stop()
}
