package service

import (
	"github.com/go-serv/foundation/internal/ancillary"
	"github.com/go-serv/foundation/pkg/z"
	"google.golang.org/grpc"
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
	name string
	//codec     z.CodecInterface
	app       z.AppServerInterface
	cfg       z.ServiceCfgInterface
	endpoints []z.EndpointInterface
	state     State
	mwGroup   z.MiddlewareInterface
	ancillary.MethodMustBeImplemented
}

func Reflection() z.ReflectInterface {
	return ref
}

func (s *service) Name() string {
	return s.name
}

func (s *service) State() State {
	return s.state
}

func (s *service) Endpoints() []z.EndpointInterface {
	return s.endpoints
}

func (s *service) App() z.AppServerInterface {
	return s.app
}

func (s *service) BindApp(app z.AppServerInterface) {
	s.app = app
}

//func (s *service) Codec() z.CodecInterface {
//	return s.codec
//}
//
//func (s *service) WithCodec(cc z.CodecInterface) {
//	s.codec = cc
//}

func (s *service) Register(srv *grpc.Server) {
	s.MethodMustBeImplemented.Panic()
}
