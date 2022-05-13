package service_test

import (
	sample "github.com/go-serv/service/internal/autogen/proto"
	"github.com/go-serv/service/pkg"
	"google.golang.org/grpc"
	"testing"
	"time"
)

var serviceName = sample.SampleService_ServiceDesc.ServiceName

type sampleService struct {
	sample.SampleServiceServer
	pkg.NetworkServiceInterface
}

func (s *sampleService) Service_Register(srv *grpc.Server) {
	sample.RegisterSampleServiceServer(srv, s)
}

func TestNetworkServiceStart(t *testing.T) {
	svc := &sampleService{}
	svc.NetworkServiceInterface = pkg.NewNetworkService(serviceName, nil)
	endpoint := pkg.NewTcp4Endpoint(svc, "localhost", 9090)
	svc.Service_AddEndpoint(endpoint)
	time.AfterFunc(time.Millisecond*10, func() {
		svc.Service_Stop()
	})
	svc.Service_Start()
	state := pkg.ServiceState(svc.Service_State())
	if !state.IsStopped() {
		t.Fatalf("failed to stop service, expected state 'Stopped', got %s\n", state.String())
	}
}
