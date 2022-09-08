// Code generated by protoc-gen-go-grpc. DO NOT EDIT.
// versions:
// - protoc-gen-go-grpc v1.2.0
// - protoc             v3.19.4
// source: net/net.proto

package net

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

// NetParcelClient is the client API for NetParcel service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://pkg.go.dev/google.golang.org/grpc/?tab=doc#ClientConn.NewStream.
type NetParcelClient interface {
	// A dummy call for testing.
	Ping(ctx context.Context, in *Ping_Request, opts ...grpc.CallOption) (*Ping_Response, error)
	// Creates a new secure session. Can be used when TLS layer is not available.
	SecureSession(ctx context.Context, in *Session_Request, opts ...grpc.CallOption) (*Session_Response, error)
	//
	FtpNewSession(ctx context.Context, in *Ftp_NewSession_Request, opts ...grpc.CallOption) (*Ftp_NewSession_Response, error)
	//
	FtpTransfer(ctx context.Context, in *Ftp_FileChunk_Request, opts ...grpc.CallOption) (*Ftp_FileChunk_Response, error)
	//
	FtpInquiry(ctx context.Context, in *Ftp_Inquiry_Request, opts ...grpc.CallOption) (*Ftp_Inquiry_Response, error)
}

type netParcelClient struct {
	cc grpc.ClientConnInterface
}

func NewNetParcelClient(cc grpc.ClientConnInterface) NetParcelClient {
	return &netParcelClient{cc}
}

func (c *netParcelClient) Ping(ctx context.Context, in *Ping_Request, opts ...grpc.CallOption) (*Ping_Response, error) {
	var inw, outw interface{}
	out := new(Ping_Response)
	if msgWrapperHandler != nil {
		inw, outw = msgWrapperHandler(in), msgWrapperHandler(out)
	} else {
		inw, outw = in, out
	}
	err := c.cc.Invoke(ctx, "/go_serv.net.NetParcel/Ping", inw, outw, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *netParcelClient) SecureSession(ctx context.Context, in *Session_Request, opts ...grpc.CallOption) (*Session_Response, error) {
	var inw, outw interface{}
	out := new(Session_Response)
	if msgWrapperHandler != nil {
		inw, outw = msgWrapperHandler(in), msgWrapperHandler(out)
	} else {
		inw, outw = in, out
	}
	err := c.cc.Invoke(ctx, "/go_serv.net.NetParcel/SecureSession", inw, outw, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *netParcelClient) FtpNewSession(ctx context.Context, in *Ftp_NewSession_Request, opts ...grpc.CallOption) (*Ftp_NewSession_Response, error) {
	var inw, outw interface{}
	out := new(Ftp_NewSession_Response)
	if msgWrapperHandler != nil {
		inw, outw = msgWrapperHandler(in), msgWrapperHandler(out)
	} else {
		inw, outw = in, out
	}
	err := c.cc.Invoke(ctx, "/go_serv.net.NetParcel/FtpNewSession", inw, outw, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *netParcelClient) FtpTransfer(ctx context.Context, in *Ftp_FileChunk_Request, opts ...grpc.CallOption) (*Ftp_FileChunk_Response, error) {
	var inw, outw interface{}
	out := new(Ftp_FileChunk_Response)
	if msgWrapperHandler != nil {
		inw, outw = msgWrapperHandler(in), msgWrapperHandler(out)
	} else {
		inw, outw = in, out
	}
	err := c.cc.Invoke(ctx, "/go_serv.net.NetParcel/FtpTransfer", inw, outw, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *netParcelClient) FtpInquiry(ctx context.Context, in *Ftp_Inquiry_Request, opts ...grpc.CallOption) (*Ftp_Inquiry_Response, error) {
	var inw, outw interface{}
	out := new(Ftp_Inquiry_Response)
	if msgWrapperHandler != nil {
		inw, outw = msgWrapperHandler(in), msgWrapperHandler(out)
	} else {
		inw, outw = in, out
	}
	err := c.cc.Invoke(ctx, "/go_serv.net.NetParcel/FtpInquiry", inw, outw, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// NetParcelServer is the server API for NetParcel service.
// All implementations must embed UnimplementedNetParcelServer
// for forward compatibility
type NetParcelServer interface {
	// A dummy call for testing.
	Ping(context.Context, *Ping_Request) (*Ping_Response, error)
	// Creates a new secure session. Can be used when TLS layer is not available.
	SecureSession(context.Context, *Session_Request) (*Session_Response, error)
	//
	FtpNewSession(context.Context, *Ftp_NewSession_Request) (*Ftp_NewSession_Response, error)
	//
	FtpTransfer(context.Context, *Ftp_FileChunk_Request) (*Ftp_FileChunk_Response, error)
	//
	FtpInquiry(context.Context, *Ftp_Inquiry_Request) (*Ftp_Inquiry_Response, error)
	mustEmbedUnimplementedNetParcelServer()
}

// UnimplementedNetParcelServer must be embedded to have forward compatible implementations.
type UnimplementedNetParcelServer struct {
}

func (UnimplementedNetParcelServer) Ping(context.Context, *Ping_Request) (*Ping_Response, error) {
	return nil, status.Errorf(codes.Unimplemented, "method Ping not implemented")
}
func (UnimplementedNetParcelServer) SecureSession(context.Context, *Session_Request) (*Session_Response, error) {
	return nil, status.Errorf(codes.Unimplemented, "method SecureSession not implemented")
}
func (UnimplementedNetParcelServer) FtpNewSession(context.Context, *Ftp_NewSession_Request) (*Ftp_NewSession_Response, error) {
	return nil, status.Errorf(codes.Unimplemented, "method FtpNewSession not implemented")
}
func (UnimplementedNetParcelServer) FtpTransfer(context.Context, *Ftp_FileChunk_Request) (*Ftp_FileChunk_Response, error) {
	return nil, status.Errorf(codes.Unimplemented, "method FtpTransfer not implemented")
}
func (UnimplementedNetParcelServer) FtpInquiry(context.Context, *Ftp_Inquiry_Request) (*Ftp_Inquiry_Response, error) {
	return nil, status.Errorf(codes.Unimplemented, "method FtpInquiry not implemented")
}
func (UnimplementedNetParcelServer) mustEmbedUnimplementedNetParcelServer() {}

// UnsafeNetParcelServer may be embedded to opt out of forward compatibility for this service.
// Use of this interface is not recommended, as added methods to NetParcelServer will
// result in compilation errors.
type UnsafeNetParcelServer interface {
	mustEmbedUnimplementedNetParcelServer()
}

func RegisterNetParcelServer(s grpc.ServiceRegistrar, srv NetParcelServer) {
	s.RegisterService(&NetParcel_ServiceDesc, srv)
}

func _NetParcel_Ping_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	var in interface{}
	var unwrap func() interface{}
	if msgWrapperHandler != nil {
		in = msgWrapperHandler(new(Ping_Request))
		unwrap = func() interface{} {
			return in.(encoding.MessageWrapper).Interface().(*Ping_Request)
		}
	} else {
		in = new(Ping_Request)
		unwrap = func() interface{} {
			return in
		}
	}
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(NetParcelServer).Ping(ctx, unwrap().(*Ping_Request))
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/go_serv.net.NetParcel/Ping",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(NetParcelServer).Ping(ctx, unwrap().(*Ping_Request))
	}
	return interceptor(ctx, in, info, handler)
}

func _NetParcel_SecureSession_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	var in interface{}
	var unwrap func() interface{}
	if msgWrapperHandler != nil {
		in = msgWrapperHandler(new(Session_Request))
		unwrap = func() interface{} {
			return in.(encoding.MessageWrapper).Interface().(*Session_Request)
		}
	} else {
		in = new(Session_Request)
		unwrap = func() interface{} {
			return in
		}
	}
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(NetParcelServer).SecureSession(ctx, unwrap().(*Session_Request))
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/go_serv.net.NetParcel/SecureSession",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(NetParcelServer).SecureSession(ctx, unwrap().(*Session_Request))
	}
	return interceptor(ctx, in, info, handler)
}

func _NetParcel_FtpNewSession_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	var in interface{}
	var unwrap func() interface{}
	if msgWrapperHandler != nil {
		in = msgWrapperHandler(new(Ftp_NewSession_Request))
		unwrap = func() interface{} {
			return in.(encoding.MessageWrapper).Interface().(*Ftp_NewSession_Request)
		}
	} else {
		in = new(Ftp_NewSession_Request)
		unwrap = func() interface{} {
			return in
		}
	}
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(NetParcelServer).FtpNewSession(ctx, unwrap().(*Ftp_NewSession_Request))
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/go_serv.net.NetParcel/FtpNewSession",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(NetParcelServer).FtpNewSession(ctx, unwrap().(*Ftp_NewSession_Request))
	}
	return interceptor(ctx, in, info, handler)
}

func _NetParcel_FtpTransfer_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	var in interface{}
	var unwrap func() interface{}
	if msgWrapperHandler != nil {
		in = msgWrapperHandler(new(Ftp_FileChunk_Request))
		unwrap = func() interface{} {
			return in.(encoding.MessageWrapper).Interface().(*Ftp_FileChunk_Request)
		}
	} else {
		in = new(Ftp_FileChunk_Request)
		unwrap = func() interface{} {
			return in
		}
	}
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(NetParcelServer).FtpTransfer(ctx, unwrap().(*Ftp_FileChunk_Request))
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/go_serv.net.NetParcel/FtpTransfer",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(NetParcelServer).FtpTransfer(ctx, unwrap().(*Ftp_FileChunk_Request))
	}
	return interceptor(ctx, in, info, handler)
}

func _NetParcel_FtpInquiry_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	var in interface{}
	var unwrap func() interface{}
	if msgWrapperHandler != nil {
		in = msgWrapperHandler(new(Ftp_Inquiry_Request))
		unwrap = func() interface{} {
			return in.(encoding.MessageWrapper).Interface().(*Ftp_Inquiry_Request)
		}
	} else {
		in = new(Ftp_Inquiry_Request)
		unwrap = func() interface{} {
			return in
		}
	}
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(NetParcelServer).FtpInquiry(ctx, unwrap().(*Ftp_Inquiry_Request))
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/go_serv.net.NetParcel/FtpInquiry",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(NetParcelServer).FtpInquiry(ctx, unwrap().(*Ftp_Inquiry_Request))
	}
	return interceptor(ctx, in, info, handler)
}

// NetParcel_ServiceDesc is the grpc.ServiceDesc for NetParcel service.
// It's only intended for direct use with grpc.RegisterService,
// and not to be introspected or modified (even as a copy)
var NetParcel_ServiceDesc = grpc.ServiceDesc{
	ServiceName: "go_serv.net.NetParcel",
	HandlerType: (*NetParcelServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "Ping",
			Handler:    _NetParcel_Ping_Handler,
		},
		{
			MethodName: "SecureSession",
			Handler:    _NetParcel_SecureSession_Handler,
		},
		{
			MethodName: "FtpNewSession",
			Handler:    _NetParcel_FtpNewSession_Handler,
		},
		{
			MethodName: "FtpTransfer",
			Handler:    _NetParcel_FtpTransfer_Handler,
		},
		{
			MethodName: "FtpInquiry",
			Handler:    _NetParcel_FtpInquiry_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "net/net.proto",
}
