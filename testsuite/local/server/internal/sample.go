package internal

import (
	"context"
	i "github.com/go-serv/service/internal"
	proto "github.com/go-serv/service/internal/autogen/proto/local"
	"google.golang.org/grpc"
	"google.golang.org/protobuf/reflect/protoreflect"
)

var Name = protoreflect.FullName(proto.Sample_ServiceDesc.ServiceName)

type serviceImpl struct {
	proto.UnimplementedSampleServer
}

type sample struct {
	i.LocalServiceInterface
	impl serviceImpl
}

func (s *sample) Register(srv *grpc.Server) {
	proto.RegisterSampleServer(srv, s.impl)
}

func (s serviceImpl) DoLargeRequest(context.Context, *proto.LargeRequest_Request) (*proto.LargeRequest_Response, error) {
	res := &proto.LargeRequest_Response{}
	return res, nil
}
