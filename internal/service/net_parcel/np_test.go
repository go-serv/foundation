package net_parcel

import (
	job "github.com/AgentCoop/go-work"
	"github.com/go-serv/service/internal/ancillary/memoize"
	"github.com/go-serv/service/internal/autogen/proto/net"
	"github.com/go-serv/service/internal/runtime"
	srv "github.com/go-serv/service/internal/server"
	net_srv "github.com/go-serv/service/internal/server/net"
	np_client "github.com/go-serv/service/internal/service/net_parcel/client"
	"github.com/go-serv/service/internal/service/net_parcel/server/ftp"
	"github.com/go-serv/service/pkg/y/netparcel"
	"github.com/go-serv/service/pkg/z"
	"github.com/go-serv/service/pkg/z/platform"
	"os"
	"testing"
	"time"
)

func init() {
	ftpResolver := memoize.NewMemoizer(func(...any) (v any, err error) {
		v = []z.FtpUploadProfileInterface{ftp.NewUploadProfile("./", 10_000, 0755)}
		return
	})
	runtime.Runtime().AddResolver(z.FtpUploadProfilerResolver, ftpResolver)
}

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
		secSessRes *net.Session_Response
		err        error
	)
	tf := setup(t)
	getNonceTask := func(j job.JobInterface) (job.Init, job.Run, job.Finalize) {
		const nonceLen = 32
		run := func(task job.TaskInterface) {
			cc := j.GetValue().(netparcel.NetParcelClientInterface)
			req := &net.Session_Request{}
			req.NonceLength = nonceLen
			req.Lifetime = 3600
			secSessRes, err = cc.SecureSession(req)
			task.Assert(err)
			if len(secSessRes.GetNonce()) != nonceLen {
				t.Fatalf("expected nonce length %d, got %d", nonceLen, len(secSessRes.GetNonce()))
			}
			cwd, _ := os.Getwd()
			tarball := platform.Pathname("").ComposePath(cwd, "..", "..", "..", "testsuite", "fixtures", "tarball.tar.gz")
			err = cc.FtpTransferFile(tarball, false, true)
			if err != nil {
				t.Fatalf("file transfer failed with %v", err)
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
