// Code generated by protoc-gen-go-grpc. DO NOT EDIT.
// versions:
// - protoc-gen-go-grpc v1.2.0
// - protoc             v3.19.4
// source: local/local.proto

package local

import (
	context "context"
	grpc "google.golang.org/grpc"
	codes "google.golang.org/grpc/codes"
	encoding "google.golang.org/grpc/encoding"
	status "google.golang.org/grpc/status"
)

// This is a compile-time assertion to ensure that this generated file
// is compatible with the grpc package it is being compiled against.
// Requires gRPC-Go v1.32.0 or later.
const _ = grpc.SupportPackageIsVersion7

var msgWrapperHandler encoding.MessageWrapperHandler

func RegisterMessageWrapper(handler encoding.MessageWrapperHandler) {
	msgWrapperHandler = handler
}

// SampleClient is the client API for Sample service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://pkg.go.dev/google.golang.org/grpc/?tab=doc#ClientConn.NewStream.
type SampleClient interface {
	DoLargeRequest(ctx context.Context, in *LargeRequest_Request, opts ...grpc.CallOption) (*LargeRequest_Response, error)
	//
	DoLargeRequestIpc(ctx context.Context, in *LargeRequestIpc_Request, opts ...grpc.CallOption) (*LargeRequestIpc_Response, error)
}

type sampleClient struct {
	cc grpc.ClientConnInterface
}

func NewSampleClient(cc grpc.ClientConnInterface) SampleClient {
	return &sampleClient{cc}
}

func (c *sampleClient) DoLargeRequest(ctx context.Context, in *LargeRequest_Request, opts ...grpc.CallOption) (*LargeRequest_Response, error) {
	var inw, outw interface{}
	out := new(LargeRequest_Response)
	if msgWrapperHandler != nil {
		inw, outw = msgWrapperHandler(in), msgWrapperHandler(out)
	} else {
		inw, outw = in, out
	}
	err := c.cc.Invoke(ctx, "/go_serv.local.Sample/DoLargeRequest", inw, outw, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *sampleClient) DoLargeRequestIpc(ctx context.Context, in *LargeRequestIpc_Request, opts ...grpc.CallOption) (*LargeRequestIpc_Response, error) {
	var inw, outw interface{}
	out := new(LargeRequestIpc_Response)
	if msgWrapperHandler != nil {
		inw, outw = msgWrapperHandler(in), msgWrapperHandler(out)
	} else {
		inw, outw = in, out
	}
	err := c.cc.Invoke(ctx, "/go_serv.local.Sample/DoLargeRequestIpc", inw, outw, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// SampleServer is the server API for Sample service.
// All implementations must embed UnimplementedSampleServer
// for forward compatibility
type SampleServer interface {
	DoLargeRequest(context.Context, *LargeRequest_Request) (*LargeRequest_Response, error)
	//
	DoLargeRequestIpc(context.Context, *LargeRequestIpc_Request) (*LargeRequestIpc_Response, error)
	mustEmbedUnimplementedSampleServer()
}

// UnimplementedSampleServer must be embedded to have forward compatible implementations.
type UnimplementedSampleServer struct {
}

func (UnimplementedSampleServer) DoLargeRequest(context.Context, *LargeRequest_Request) (*LargeRequest_Response, error) {
	return nil, status.Errorf(codes.Unimplemented, "method DoLargeRequest not implemented")
}
func (UnimplementedSampleServer) DoLargeRequestIpc(context.Context, *LargeRequestIpc_Request) (*LargeRequestIpc_Response, error) {
	return nil, status.Errorf(codes.Unimplemented, "method DoLargeRequestIpc not implemented")
}
func (UnimplementedSampleServer) mustEmbedUnimplementedSampleServer() {}

// UnsafeSampleServer may be embedded to opt out of forward compatibility for this service.
// Use of this interface is not recommended, as added methods to SampleServer will
// result in compilation errors.
type UnsafeSampleServer interface {
	mustEmbedUnimplementedSampleServer()
}

func RegisterSampleServer(s grpc.ServiceRegistrar, srv SampleServer) {
	s.RegisterService(&Sample_ServiceDesc, srv)
}

func _Sample_DoLargeRequest_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	var in interface{}
	var unwrap func() interface{}
	if msgWrapperHandler != nil {
		in = msgWrapperHandler(new(LargeRequest_Request))
		unwrap = func() interface{} {
			return in.(encoding.MessageWrapper).Interface().(*LargeRequest_Request)
		}
	} else {
		in = new(LargeRequest_Request)
		unwrap = func() interface{} {
			return in
		}
	}
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(SampleServer).DoLargeRequest(ctx, unwrap().(*LargeRequest_Request))
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/go_serv.local.Sample/DoLargeRequest",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(SampleServer).DoLargeRequest(ctx, unwrap().(*LargeRequest_Request))
	}
	return interceptor(ctx, in, info, handler)
}

func _Sample_DoLargeRequestIpc_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	var in interface{}
	var unwrap func() interface{}
	if msgWrapperHandler != nil {
		in = msgWrapperHandler(new(LargeRequestIpc_Request))
		unwrap = func() interface{} {
			return in.(encoding.MessageWrapper).Interface().(*LargeRequestIpc_Request)
		}
	} else {
		in = new(LargeRequestIpc_Request)
		unwrap = func() interface{} {
			return in
		}
	}
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(SampleServer).DoLargeRequestIpc(ctx, unwrap().(*LargeRequestIpc_Request))
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/go_serv.local.Sample/DoLargeRequestIpc",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(SampleServer).DoLargeRequestIpc(ctx, unwrap().(*LargeRequestIpc_Request))
	}
	return interceptor(ctx, in, info, handler)
}

// Sample_ServiceDesc is the grpc.ServiceDesc for Sample service.
// It's only intended for direct use with grpc.RegisterService,
// and not to be introspected or modified (even as a copy)
var Sample_ServiceDesc = grpc.ServiceDesc{
	ServiceName: "go_serv.local.Sample",
	HandlerType: (*SampleServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "DoLargeRequest",
			Handler:    _Sample_DoLargeRequest_Handler,
		},
		{
			MethodName: "DoLargeRequestIpc",
			Handler:    _Sample_DoLargeRequestIpc_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "local/local.proto",
}