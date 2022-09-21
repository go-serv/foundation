package tests

import (
	job "github.com/AgentCoop/go-work"
	"github.com/go-serv/foundation/app"
	np_client "github.com/go-serv/foundation/app/net_parcel/client"
	np_svc "github.com/go-serv/foundation/app/net_parcel/server"
	"github.com/go-serv/foundation/app/net_parcel/server/ftp"
	net "github.com/go-serv/foundation/internal/autogen/foundation"
	"github.com/go-serv/foundation/internal/runtime"
	"github.com/go-serv/foundation/pkg/ancillary/memoize"
	"github.com/go-serv/foundation/pkg/y/netparcel"
	"github.com/go-serv/foundation/pkg/z"
	"github.com/go-serv/foundation/pkg/z/platform"
	net_svc "github.com/go-serv/foundation/service/net"
	"os"
	"testing"
	"time"
)

func init() {
	ftpResolver := memoize.NewMemoizer(func(...any) (v any, err error) {
		v = []z.FtpUploadProfileInterface{ftp.NewUploadProfile("./", 10_000, 0755)}
		return
	})
	runtime.Runtime().RegisterResolver(z.FtpUploadProfilerResolver, ftpResolver)
}

const (
	testHost = "localhost"
	testPort = 4411
)

type testFixtures struct {
	t         *testing.T
	app       z.AppInterface
	srv       z.NetworkServiceInterface
	clientJob job.JobInterface
}

func setup(t *testing.T) *testFixtures {
	tf := new(testFixtures)
	tf.t = t
	tf.app = app.NewApp(nil)
	ep := net_svc.NewTcp4Endpoint(testHost, testPort)
	netParcelSvc := np_svc.NewNetParcel([]z.EndpointInterface{ep}, nil)
	tf.app.AddService(netParcelSvc)
	//tf.srv = net_svc.NewNetworkService(net_parcel.Name, nil, []z.EndpointInterface{ep}, tf.app)
	tf.clientJob = np_client.NewClientJob(ep)
	return tf
}

func (tf *testFixtures) serverStatus() {
	_, err := tf.app.Job().GetInterruptedBy()
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
			tarball := platform.Pathname("").ComposePath(cwd, "fixtures", "tarball.tar.gz")
			err = cc.FtpTransferFile(tarball, false, true)
			if err != nil {
				t.Fatalf("file transfer failed with %v", err)
			}
			task.Done()
		}
		return nil, run, nil
	}
	go func() {
		tf.app.Start()
	}()
	defer func() {
		tf.finalize()
	}()
	time.Sleep(time.Millisecond * 10)
	tf.serverStatus()
	tf.clientJob.AddTask(getNonceTask)
	<-tf.clientJob.Run()
	tf.app.Stop(nil)
}
