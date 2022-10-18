package middleware

import (
	"context"
	net_call "github.com/go-serv/foundation/internal/grpc/callctx/net"
	"github.com/go-serv/foundation/internal/grpc/msg/request"
	"github.com/go-serv/foundation/internal/service"
	"github.com/go-serv/foundation/pkg/z"
	"google.golang.org/grpc"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/metadata"
	"google.golang.org/grpc/status"
	"google.golang.org/protobuf/proto"
)

type serverMw struct {
	middleware
}

func (m *serverMw) UnaryServerInterceptor() grpc.UnaryServerInterceptor {
	return func(ctx context.Context, in interface{}, info *grpc.UnaryServerInfo, handler grpc.UnaryHandler) (out interface{}, err error) {
		var (
			md   metadata.MD
			ok   bool
			req  z.RequestInterface
			sref z.ServiceReflectInterface
		)
		if md, ok = metadata.FromIncomingContext(ctx); !ok {
			return nil, status.Error(codes.Internal, "failed to retrieve request metadata")
		}

		//
		clientInfo := request.NewClientInfo(ctx)
		if req, err = request.NewServerRequest(in, &md, clientInfo); err != nil {
			return
		}

		msg := in.(proto.Message)
		if sref, err = service.Reflection().ServiceReflectionFromMessage(msg); err != nil {
			return
		}

		// Pass server context through the request/response handler chains.
		srvCxt := net_call.NewServerContext(ctx, req, handler)
		if err = m.requestPassThrough(srvCxt, sref.FullName()); err != nil {
			return
		}
		if err = m.responsePassThrough(srvCxt, sref.FullName()); err != nil {
			return
		}

		// Send response headers.
		if md, err = srvCxt.Response().Meta().Dehydrate(); err != nil {
			return
		}
		if err = grpc.SendHeader(ctx, md); err != nil {
			return
		}
		out = srvCxt.Response().Data()
		return
	}
}

func (m *serverMw) StreamServerInterceptor() grpc.StreamServerInterceptor {
	return func(srv interface{}, ss grpc.ServerStream, info *grpc.StreamServerInfo, handler grpc.StreamHandler) (err error) {
		return
	}
}
