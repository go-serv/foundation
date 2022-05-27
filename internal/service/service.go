package service

import (
	i "github.com/go-serv/service/internal"
	"github.com/go-serv/service/internal/ancillary"
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

type BaseService struct {
	// Service Name in the dot notation
	Name  string
	State State
	Cfg   ConfigInterface
	Sd    i.ServiceDescriptorInterface
	cc    i.CodecInterface
	ancillary.MethodMustBeImplemented
}

func (s BaseService) Service_Descriptor() i.ServiceDescriptorInterface {
	return s.Sd
}

func (s BaseService) Service_AddMethodProtoExtension(ext *protoimpl.ExtensionInfo) {
	s.Sd.AddMethodProtoExt(ext)
}

func (s BaseService) Service_AddServiceProtoExtension(ext *protoimpl.ExtensionInfo) {
	s.Sd.AddServiceProtoExt(ext)
}

func (s BaseService) Service_Name(short bool) string {
	return s.Name
}

func (s BaseService) Service_State() State {
	return s.State
}

func (b BaseService) Service_Register(srv *grpc.Server) {
	b.MethodMustBeImplemented.Panic()
}
