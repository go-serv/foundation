package net

import (
	"context"
	net_call "github.com/go-serv/service/internal/grpc/callctx/net"
	"github.com/go-serv/service/internal/grpc/codec"
	"github.com/go-serv/service/internal/grpc/msg/request"
	"github.com/go-serv/service/pkg/z"
	"google.golang.org/grpc"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/metadata"
	"google.golang.org/grpc/status"
	"google.golang.org/protobuf/proto"
	"reflect"
	"unsafe"
)

func (mw *netMiddleware) UnaryServerInterceptor() grpc.UnaryServerInterceptor {
	return func(ctx context.Context, v interface{}, info *grpc.UnaryServerInfo, handler grpc.UnaryHandler) (out interface{}, err error) {
		var (
			md  metadata.MD
			ok  bool
			req z.RequestInterface
			df  z.DataFrameInterface
		)

		ptr1 := (*reflect.SliceHeader)(unsafe.Pointer(reflect.ValueOf(v).Elem().FieldByName("unknownFields").Addr().Pointer()))
		_ = ptr1
		// Retrieve request metadata.
		if md, ok = metadata.FromIncomingContext(ctx); !ok {
			return nil, status.Error(codes.Internal, "failed to retrieve request metadata")
		}
		//
		if df, err = codec.LoadDataFrameFromPtrPool(v.(proto.Message)); err != nil {
			return
		}
		if err = df.RemoveFromPtrPool(); err != nil {
			return
		}
		//
		clientInfo := request.NewClientInfo(ctx)
		if req, err = request.NewRequest(df, &md, clientInfo); err != nil {
			return
		}
		//
		srvCxt := net_call.NewServerContext(ctx, req, handler)
		srvCxt.WithInput(v.(proto.Message))
		if err = mw.newRequestChain().passThrough(srvCxt); err != nil {
			return
		}
		//srvCxt.WithOutput(srvCxt.Response().DataFrame().ProtoMessage())
		if err = mw.newResponseChain().passThrough(srvCxt); err != nil {
			return
		}
		// Send response headers
		if md, err = srvCxt.Response().Meta().Dehydrate(); err != nil {
			return
		}
		if err = grpc.SendHeader(ctx, md); err != nil {
			return
		}
		//
		out = srvCxt.Response().DataFrame()
		return
	}
}

func (mw *netMiddleware) StreamServerInterceptor() grpc.StreamServerInterceptor {
	return func(srv interface{}, ss grpc.ServerStream, info *grpc.StreamServerInfo, handler grpc.StreamHandler) (err error) {
		return
	}
}
