package service

import (
	"github.com/go-serv/service/internal/ancillary"
	"github.com/go-serv/service/internal/autogen/proto/go_serv"
	"github.com/go-serv/service/internal/grpc/descriptor"
	"github.com/go-serv/service/internal/grpc/request"
	_ "github.com/go-serv/service/internal/logger"
	"google.golang.org/grpc"
	"google.golang.org/protobuf/runtime/protoimpl"
)

type State int

const (
	StateInit State = iota
	StateRunning
	StateStopInProgress
	StateStopped
	StateSuspended // service is running but incoming requests are not being processed
)

type baseService struct {
	// Service name in dot notation
	name  string
	state State
	cfg   ConfigInterface
	sd    descriptor.ServiceDescriptorInterface
	ancillary.MethodMustBeImplemented
}

type localService struct {
	baseService
}

type networkService struct {
	baseService
}

func (s *baseService) Service_Descriptor() descriptor.ServiceDescriptorInterface {
	return s.sd
}

func (s *baseService) Service_AddMethodProtoExtension(ext *protoimpl.ExtensionInfo) {
	s.sd.AddMethodProtoExt(ext)
}

func (s *baseService) Service_AddServiceProtoExtension(ext *protoimpl.ExtensionInfo) {
	s.sd.AddServiceProtoExt(ext)
}

func (s *baseService) Service_Name(short bool) string {
	return s.name
}

func (s *baseService) Service_State() State {
	return s.state
}

//
// Methods to implement
//
func (b *baseService) Service_Register(srv *grpc.Server) {
	b.MethodMustBeImplemented.Panic()
}

func (b *networkService) Service_RequestNewSession(req request.RequestInterface) int32 {
	mDesc := b.sd.FindMethodDescriptorByName(req.MethodName())
	if mDesc == nil {
		return 0
	} else {
		v, has := mDesc.Get(go_serv.E_NetNewSession)
		if !has {
			return 0
		} else {
			return v.(int32)
		}
	}
}

func (b *baseService) Service_OnNewSession(req request.RequestInterface) error {
	b.MethodMustBeImplemented.Panic()
	return nil
}
