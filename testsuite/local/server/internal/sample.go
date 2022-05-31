package internal

import (
	"context"
	"crypto/md5"
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

func (s serviceImpl) DoLargeRequest(ctx context.Context, req *proto.LargeRequest_Request) (*proto.LargeRequest_Response, error) {
	res := &proto.LargeRequest_Response{}
	md5Hash := md5.Sum(req.Data)
	res.Hash = md5Hash[:]
	res.Pong = req.Ping
	return res, nil
}

func (s serviceImpl) DoLargeRequestIpc(ctx context.Context, req *proto.LargeRequestIpc_Request) (*proto.LargeRequestIpc_Response, error) {
	res := &proto.LargeRequestIpc_Response{}
	md5Hash := md5.Sum(req.Data)
	res.Md5Hash = md5Hash[:]
	res.Pong = req.Ping
	return res, nil
}
