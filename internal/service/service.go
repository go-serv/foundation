package service

import (
	"github.com/go-serv/service/internal/ancillary"
	mw_codec "github.com/go-serv/service/internal/grpc/middleware/codec"
	_ "github.com/go-serv/service/internal/logger"
	"github.com/go-serv/service/pkg/z"
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
	codec        z.CodecInterface
	codecMwGroup z.CodecMiddlewareGroupInterface
	State        State
	ancillary.MethodMustBeImplemented
}

func (s *service) Name() protoreflect.FullName {
	return s.name
}

func (s *service) Codec() z.CodecInterface {
	return s.codec
}

func (s *service) WithCodec(cc z.CodecInterface) {
	s.codec = cc
	s.codecMwGroup = mw_codec.NewCodecMiddlewareGroup(cc)
}

func (s *service) CodecMiddlewareGroup() z.CodecMiddlewareGroupInterface {
	return s.codecMwGroup
}

func (s service) Service_State() State {
	return s.State
}

func (b service) Register(srv *grpc.Server) {
	b.MethodMustBeImplemented.Panic()
}
