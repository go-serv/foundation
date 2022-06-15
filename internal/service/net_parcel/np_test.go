package net_parcel

import (
	"context"
	job "github.com/AgentCoop/go-work"
	"github.com/go-serv/service/internal/autogen/proto/net"
	net_ctx "github.com/go-serv/service/internal/grpc/context/net"
	srv "github.com/go-serv/service/internal/server"
	net_srv "github.com/go-serv/service/internal/server/net"
	np_client "github.com/go-serv/service/internal/service/net_parcel/client"
	"github.com/go-serv/service/pkg/z"
	"testing"
	"time"
)

const (
	testHost = "localhost"
	testPort = 4411
)

type testFixtures struct {
	t         *testing.T
	endpoint  z.EndpointInterface
	srv       z.NetworkServerInterface
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

//func TestRun(t *testing.T) {
//	//cwd, _ := os.Getwd()
//	cmd := exec.Command("go", "run", "/home/pihpah/go/src/github.com/go-serv/service/tests/server/main.go")
//	if err := cmd.Run(); err != nil {
//		t.Fatalf("error while running test server: %s", err.Error())
//	}
//	o, _ := cmd.Output()
//	fmt.Printf("got: %v\n", o)
//}

func TestSecureSession(t *testing.T) {
	var (
		secSessRes    *net.Session_Response
		ftpNewSessRes *net.Ftp_NewSession_Response
		err           error
	)
	tf := setup(t)
	getNonceTask := func(j job.JobInterface) (job.Init, job.Run, job.Finalize) {
		const nonceLen = 32
		run := func(task job.TaskInterface) {
			cc := j.GetValue().(net.NetParcelClient)
			req := &net.Session_Request{}
			req.NonceLength = nonceLen
			ctx := net_ctx.NewClientContext(context.Background())
			secSessRes, err = cc.SecureSession(ctx, req)
			task.Assert(err)
			if len(secSessRes.GetNonce()) != nonceLen {
				t.Fatalf("expected nonce length %d, got %d", nonceLen, len(secSessRes.GetNonce()))
			}
			// Start an FTP session
			ftpNewSessReq := &net.Ftp_NewSession_Request{}
			ftpNewSessReq.Files = append(ftpNewSessReq.Files, &net.Ftp_FileInfo{Pathname: "file1.bin", Size: 1200})
			ftpNewSessRes, err = cc.FtpNewSession(ctx, ftpNewSessReq)
			task.Assert(err)
			//
			_ = ftpNewSessRes
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
