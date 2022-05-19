package service

import (
	"github.com/go-serv/service/internal/ancillary"
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
	name             string
	state            State
	cfg              ConfigInterface
	methodProtoExts  []*methodProtoExt
	serviceProtoExts []*serviceProtoExt
	ancillary.MethodMustBeImplemented
}

type localService struct {
	baseService
}

type networkService struct {
	baseService
}

//
// Protobuf extensions
//
type protoExt struct {
	info       *protoimpl.ExtensionInfo
	defaultVal interface{}
}

type methodProtoExt struct {
	protoExt
}

type serviceProtoExt struct {
	protoExt
}

func (s *baseService) Service_AddMethodProtoExtension(info *protoimpl.ExtensionInfo, defaultVal interface{}) {
	s.methodProtoExts = append(s.methodProtoExts, &methodProtoExt{protoExt{info, defaultVal}})
}

func (s *baseService) Service_AddServiceProtoExtension(info *protoimpl.ExtensionInfo, defaultVal interface{}) {
	s.serviceProtoExts = append(s.serviceProtoExts, &serviceProtoExt{protoExt{info, defaultVal}})
}

func (s *baseService) Service_ServiceProtoExtensions() []*serviceProtoExt {
	return s.serviceProtoExts
}

func (s *baseService) Service_MethodProtoExtensions() []*methodProtoExt {
	return s.methodProtoExts
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

func (b *baseService) Service_OnNewSession(req request.RequestInterface) error {
	b.MethodMustBeImplemented.Panic()
	return nil
}
