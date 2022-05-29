package service

import (
	i "github.com/go-serv/service/internal"
	"github.com/go-serv/service/internal/ancillary"
	mw_codec "github.com/go-serv/service/internal/grpc/mw_group/codec"
	_ "github.com/go-serv/service/internal/logger"
	"google.golang.org/grpc"
	"google.golang.org/protobuf/reflect/protoreflect"
)

type State int

const (
	StateInit State = iota
	StateRunning
	StateStopInProgress
	StateStopped
	StateSuspended // service is running but incoming requests are not being processed
)

type service struct {
	name         protoreflect.FullName
	codec        i.CodecInterface
	codecMwGroup i.CodecMiddlewareGroupInterface
	State        State
	sd           i.ServiceDescriptorInterface
	ancillary.MethodMustBeImplemented
}

func (s *service) Name() protoreflect.FullName {
	return s.name
}

func (s *service) Codec() i.CodecInterface {
	return s.codec
}

func (s *service) WithCodec(cc i.CodecInterface) {
	s.codec = cc
	s.codecMwGroup = mw_codec.NewCodecMiddlewareGroup(cc)
}

func (s *service) CodecMiddlewareGroup() i.CodecMiddlewareGroupInterface {
	return s.codecMwGroup
}

func (s service) Service_State() State {
	return s.State
}

func (b service) Register(srv *grpc.Server) {
	b.MethodMustBeImplemented.Panic()
}
