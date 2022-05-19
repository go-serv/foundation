package client

import (
	job "github.com/AgentCoop/go-work"
	"github.com/go-serv/service/internal/server"
	"google.golang.org/grpc"
)

type clientInterface interface {
	Client_Endpoint() server.EndpointInterface
	Client_NewClient(cc grpc.ClientConnInterface)
	Client_ConnectTask(j job.JobInterface) (job.Init, job.Run, job.Finalize)
}

type NetworkClientInterface interface {
	Client_Endpoint() server.EndpointInterface
	Client_NewClient(cc grpc.ClientConnInterface)
	Client_ConnectTask(j job.JobInterface) (job.Init, job.Run, job.Finalize)
}

type LocalClientInterface interface {
	clientInterface
}
