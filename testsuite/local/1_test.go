package local

import (
	"bytes"
	"context"
	"crypto/md5"
	"crypto/rand"
	"fmt"
	job "github.com/AgentCoop/go-work"
	i "github.com/go-serv/service/internal"
	proto "github.com/go-serv/service/internal/autogen/proto/local"
	srv "github.com/go-serv/service/internal/server"
	"github.com/go-serv/service/internal/service/local"
	"github.com/go-serv/service/testsuite/local/client"
	"google.golang.org/protobuf/reflect/protoreflect"
	rr "math/rand"
	"strings"
	"testing"
	"time"
)

var svcName = protoreflect.FullName(proto.Sample_ServiceDesc.ServiceName)

type taskData struct {
	endpoint  i.EndpointInterface
	svc       i.LocalServiceInterface
	srv       i.NetworkServerInterface
	clientJob job.JobInterface
}

type testFixtures struct {
	taskData
	t *testing.T
}

type benchFixtures struct {
	taskData
	b *testing.B
}

const pingMsg = "Hello, World!"
const (
	minDataSize = 30000
	maxDataSize = 40000
)

func largeDataIpcCall(cc proto.SampleClient) (req *proto.LargeRequestIpc_Request, res *proto.LargeRequestIpc_Response, err error) {
	size := minDataSize + rr.Intn(maxDataSize-minDataSize)
	dataChunk := make([]byte, size)
	_, err = rand.Read(dataChunk)
	if err != nil {
		return
	}
	// Do call
	req = &proto.LargeRequestIpc_Request{}
	req.Data = dataChunk
	req.Ping = pingMsg
	res, err = cc.DoLargeRequestIpc(context.Background(), req)
	if err != nil {
		return
	}
	return
}

func largeDataCall(cc proto.SampleClient) (req *proto.LargeRequest_Request, res *proto.LargeRequest_Response, err error) {
	size := minDataSize + rr.Intn(maxDataSize-minDataSize)
	dataChunk := make([]byte, size)
	_, err = rand.Read(dataChunk)
	if err != nil {
		return
	}
	// Do call
	req = &proto.LargeRequest_Request{}
	req.Data = dataChunk
	req.Ping = pingMsg
	res, err = cc.DoLargeRequest(context.Background(), req)
	if err != nil {
		return
	}
	return
}

const N = 100

func (f *testFixtures) doLargeReqTask(j job.JobInterface) (job.Init, job.Run, job.Finalize) {
	run := func(task job.TaskInterface) {
		cc := f.clientJob.GetValue().(proto.SampleClient)
		for kk := 0; kk < N; kk++ {
			req, res, err := largeDataCall(cc)
			task.Assert(err)
			reqHash := md5.Sum(req.Data)
			if strings.Compare(res.GetPong(), req.GetPing()) != 0 {
				f.t.Fatalf("expected %s, got %s", req.GetPing(), res.GetPong())
			}
			if bytes.Compare(reqHash[:], res.Hash) != 0 {
				f.t.Fatal("wrong data hash")
			}
		}
		task.Done()
	}
	return nil, run, nil
}

func (f *testFixtures) doLargeReqIpcTask(j job.JobInterface) (job.Init, job.Run, job.Finalize) {
	run := func(task job.TaskInterface) {
		cc := f.clientJob.GetValue().(proto.SampleClient)
		for kk := 0; kk < N; kk++ {
			//time.Sleep(time.Millisecond * 50)
			req, res, err := largeDataIpcCall(cc)
			task.Assert(err)
			reqHash := md5.Sum(req.Data)
			if strings.Compare(res.GetPong(), req.GetPing()) != 0 {
				f.t.Fatalf("expected %s, got %s", req.GetPing(), res.GetPong())
			}
			if bytes.Compare(reqHash[:], res.Md5Hash) != 0 {
				f.t.Fatal("wrong data hash")
			}
		}
		task.Done()
	}
	return nil, run, nil
}

func (b *benchFixtures) largeDataIpcCallTask(j job.JobInterface) (job.Init, job.Run, job.Finalize) {
	run := func(task job.TaskInterface) {
		cc := b.clientJob.GetValue().(proto.SampleClient)
		for k := 0; k < b.b.N; k++ {
			_, _, err := largeDataIpcCall(cc)
			task.Assert(err)
		}
		task.Done()
	}
	return nil, run, nil
}

func (b *benchFixtures) largeDataCallTask(j job.JobInterface) (job.Init, job.Run, job.Finalize) {
	run := func(task job.TaskInterface) {
		cc := b.clientJob.GetValue().(proto.SampleClient)
		for k := 0; k < b.b.N; k++ {
			_, _, err := largeDataCall(cc)
			task.Assert(err)
		}
		task.Done()
	}
	return nil, run, nil
}

func setup(t *testing.T) *testFixtures {
	tf := new(testFixtures)
	tf.t = t
	tf.svc = local.NewService(svcName)
	tf.endpoint = srv.NewLocalEndpoint(tf.svc)
	tf.clientJob = client.NewClientJob(tf.endpoint)
	return tf
}

func setupBench(b *testing.B) *benchFixtures {
	bf := new(benchFixtures)
	bf.b = b
	bf.svc = local.NewService(svcName)
	bf.endpoint = srv.NewLocalEndpoint(bf.svc)
	bf.clientJob = client.NewClientJob(bf.endpoint)
	return bf
}

func (bf *benchFixtures) finalize() {
	_, err := bf.clientJob.GetInterruptedBy()
	if err != nil {
		panic(fmt.Errorf("client job failed with %v\n", err))
	}
}

func (tf *testFixtures) finalize() {
	_, err := tf.clientJob.GetInterruptedBy()
	if err != nil {
		tf.t.Fatalf("client job failed with %v\n", err)
	}
}

func TestSharedMemIpc(t *testing.T) {
	tf := setup(t)
	defer func() {
		tf.finalize()
	}()
	//time.Sleep(time.Millisecond * 10)
	now := time.Now().UnixMilli()
	for k := 0; k < 1; k++ {
		//tf.clientJob.AddTask(tf.doLargeReqTask)
		tf.clientJob.AddTask(tf.doLargeReqIpcTask)
	}
	<-tf.clientJob.Run()
	end := time.Now().UnixMilli()
	fmt.Printf("time elapsed: %d milliseconds\n", end-now)
}

func BenchmarkIpc1M(b *testing.B) {
	bf := setupBench(b)
	bf.clientJob.AddTask(bf.largeDataIpcCallTask)
	<-bf.clientJob.Run()
}

func BenchmarkNoIpc(b *testing.B) {
	bf := setupBench(b)
	bf.clientJob.AddTask(bf.largeDataCallTask)
	<-bf.clientJob.Run()
}
