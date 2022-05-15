package service_test

import (
	"github.com/go-serv/service/pkg"
	"google.golang.org/grpc"
	"testing"
	"time"
)

//var serviceName = sample.SampleService_ServiceDesc.ServiceName

type sampleService struct {
	//sample.SampleServiceServer
	pkg.NetworkServiceInterface
}

func (s *sampleService) Service_Register(srv *grpc.Server) {
	//sample.RegisterSampleServiceServer(srv, s)
}

func TestNetworkServiceStart(t *testing.T) {
	svc := new(sampleService)
	svc.NetworkServiceInterface = pkg.NewNetworkService("foo", nil)
	port := 4411
	e4 := pkg.NewTcp4Endpoint(svc, "localhost", port)
	e6 := pkg.NewTcp6Endpoint(svc, "[::1]", port)
	local := pkg.NewLocalEndpoint(svc, svc.Service_Name(true))
	svc.Service_AddEndpoint(e4)
	svc.Service_AddEndpoint(e6)
	svc.Service_AddEndpoint(local)
	time.AfterFunc(time.Millisecond*10, func() {
		svc.Service_Stop()
	})
	svc.Service_Start()
	state := pkg.ServiceState(svc.Service_State())
	if !state.IsStopped() {
		t.Fatalf("failed to stop service, expected state 'Stopped', got %s\n", state.String())
	}
}
