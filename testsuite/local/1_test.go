package local_test

import (
	"bytes"
	"context"
	"crypto/md5"
	"crypto/rand"
	job "github.com/AgentCoop/go-work"
	i "github.com/go-serv/service/internal"
	proto "github.com/go-serv/service/internal/autogen/proto/local"
	srv "github.com/go-serv/service/internal/server"
	"github.com/go-serv/service/internal/service/local"
	"github.com/go-serv/service/testsuite/local/client"
	"google.golang.org/protobuf/reflect/protoreflect"
	"strings"
	"testing"
	"time"
)

var svcName = protoreflect.FullName(proto.Sample_ServiceDesc.ServiceName)

type testFixtures struct {
	t         *testing.T
	endpoint  i.EndpointInterface
	svc       i.LocalServiceInterface
	srv       i.NetworkServerInterface
	clientJob job.JobInterface
}

func setup(t *testing.T) *testFixtures {
	tf := new(testFixtures)
	tf.t = t
	tf.svc = local.NewService(svcName)
	tf.endpoint = srv.NewLocalEndpoint(tf.svc)
	tf.clientJob = client.NewClientJob(tf.endpoint)
	return tf
}

func (tf *testFixtures) finalize() {
	_, err := tf.clientJob.GetInterruptedBy()
	if err != nil {
		tf.t.Fatalf("client job failed with %v\n", err)
	}
}
func TestSharedMemIpc(t *testing.T) {
	tf := setup(t)
	doLargeReq := func(j job.JobInterface) (job.Init, job.Run, job.Finalize) {
		chunkSize := 4 * 1024 * 1024
		run := func(task job.TaskInterface) {
			cc := j.GetValue().(proto.SampleClient)
			dataChunk := make([]byte, chunkSize)
			_, err := rand.Read(dataChunk)
			task.Assert(err)
			// Do call
			req := &proto.LargeRequestIpc_Request{}
			req.Data = dataChunk
			senderMd5Hash := md5.Sum(dataChunk)
			req.Ping = "Hello, World!"
			res, err := cc.DoLargeRequestIpc(context.Background(), req)
			task.Assert(err)
			// Check the call response
			if strings.Compare(res.GetPong(), req.GetPing()) != 0 {
				t.Fatalf("expected %s, got %s", req.GetPing(), res.GetPong())
			}
			if bytes.Compare(senderMd5Hash[:], res.Md5Hash) != 0 {
				t.Fatal("wrong data hash")
			}
			task.Done()
		}
		return nil, run, nil
	}
	defer func() {
		tf.finalize()
	}()
	time.Sleep(time.Millisecond * 10)
	tf.clientJob.AddTask(doLargeReq)
	<-tf.clientJob.Run()
}
