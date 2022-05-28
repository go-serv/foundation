package server_test

import (
	"github.com/go-serv/service/pkg"
	"testing"
	"time"
)

func TestStartAndStop(t *testing.T) {
	srv := pkg.NewNetServer()
	port := 4411
	e4 := pkg.NewTcp4Endpoint("localhost", port)
	e6 := pkg.NewTcp6Endpoint("[::1]", port)
	srv.AddEndpoint(e4)
	srv.AddEndpoint(e6)
	time.AfterFunc(time.Millisecond*50, func() {
		srv.Stop()
	})
	srv.Start()
	//state := pkg.ServiceState(svc.Service_State())
	//if !state.IsStopped() {
	//	t.Fatalf("failed to stop service, expected state 'Stopped', got %s\n", state.String())
	//}
}
