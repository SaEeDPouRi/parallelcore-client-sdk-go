// Code generated by protoc-gen-go. DO NOT EDIT.
// source: interface.proto

/*
Package engine_client_proto is a generated protocol buffer package.

It is generated from these files:
	interface.proto

It has these top-level messages:
	AuthRequest
	Request
	Response
*/
package engine_client_proto

import proto "github.com/golang/protobuf/proto"
import fmt "fmt"
import math "math"

import (
	context "golang.org/x/net/context"
	grpc "google.golang.org/grpc"
)

// Reference imports to suppress errors if they are not otherwise used.
var _ = proto.Marshal
var _ = fmt.Errorf
var _ = math.Inf

// This is a compile-time assertion to ensure that this generated file
// is compatible with the proto package it is being compiled against.
// A compilation error at this line likely means your copy of the
// proto package needs to be updated.
const _ = proto.ProtoPackageIsVersion2 // please upgrade the proto package

type AuthRequest struct {
	ClientId   []byte `protobuf:"bytes,1,opt,name=clientId,proto3" json:"clientId,omitempty"`
	Credential []byte `protobuf:"bytes,2,opt,name=credential,proto3" json:"credential,omitempty"`
}

func (m *AuthRequest) Reset()                    { *m = AuthRequest{} }
func (m *AuthRequest) String() string            { return proto.CompactTextString(m) }
func (*AuthRequest) ProtoMessage()               {}
func (*AuthRequest) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{0} }

func (m *AuthRequest) GetClientId() []byte {
	if m != nil {
		return m.ClientId
	}
	return nil
}

func (m *AuthRequest) GetCredential() []byte {
	if m != nil {
		return m.Credential
	}
	return nil
}

type Request struct {
	Payload []byte `protobuf:"bytes,1,opt,name=payload,proto3" json:"payload,omitempty"`
}

func (m *Request) Reset()                    { *m = Request{} }
func (m *Request) String() string            { return proto.CompactTextString(m) }
func (*Request) ProtoMessage()               {}
func (*Request) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{1} }

func (m *Request) GetPayload() []byte {
	if m != nil {
		return m.Payload
	}
	return nil
}

type Response struct {
	Payload []byte `protobuf:"bytes,1,opt,name=payload,proto3" json:"payload,omitempty"`
	Error   []byte `protobuf:"bytes,2,opt,name=error,proto3" json:"error,omitempty"`
}

func (m *Response) Reset()                    { *m = Response{} }
func (m *Response) String() string            { return proto.CompactTextString(m) }
func (*Response) ProtoMessage()               {}
func (*Response) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{2} }

func (m *Response) GetPayload() []byte {
	if m != nil {
		return m.Payload
	}
	return nil
}

func (m *Response) GetError() []byte {
	if m != nil {
		return m.Error
	}
	return nil
}

func init() {
	proto.RegisterType((*AuthRequest)(nil), "engine_client_proto.AuthRequest")
	proto.RegisterType((*Request)(nil), "engine_client_proto.Request")
	proto.RegisterType((*Response)(nil), "engine_client_proto.Response")
}

// Reference imports to suppress errors if they are not otherwise used.
var _ context.Context
var _ grpc.ClientConn

// This is a compile-time assertion to ensure that this generated file
// is compatible with the grpc package it is being compiled against.
const _ = grpc.SupportPackageIsVersion4

// Client API for RequestHandler service

type RequestHandlerClient interface {
	Invoke(ctx context.Context, in *Request, opts ...grpc.CallOption) (*Response, error)
	SysMan(ctx context.Context, in *Request, opts ...grpc.CallOption) (*Response, error)
	Auth(ctx context.Context, in *AuthRequest, opts ...grpc.CallOption) (*Response, error)
	Ping(ctx context.Context, in *Request, opts ...grpc.CallOption) (*Response, error)
}

type requestHandlerClient struct {
	cc *grpc.ClientConn
}

func NewRequestHandlerClient(cc *grpc.ClientConn) RequestHandlerClient {
	return &requestHandlerClient{cc}
}

func (c *requestHandlerClient) Invoke(ctx context.Context, in *Request, opts ...grpc.CallOption) (*Response, error) {
	out := new(Response)
	err := grpc.Invoke(ctx, "/engine_client_proto.RequestHandler/Invoke", in, out, c.cc, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *requestHandlerClient) SysMan(ctx context.Context, in *Request, opts ...grpc.CallOption) (*Response, error) {
	out := new(Response)
	err := grpc.Invoke(ctx, "/engine_client_proto.RequestHandler/SysMan", in, out, c.cc, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *requestHandlerClient) Auth(ctx context.Context, in *AuthRequest, opts ...grpc.CallOption) (*Response, error) {
	out := new(Response)
	err := grpc.Invoke(ctx, "/engine_client_proto.RequestHandler/Auth", in, out, c.cc, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *requestHandlerClient) Ping(ctx context.Context, in *Request, opts ...grpc.CallOption) (*Response, error) {
	out := new(Response)
	err := grpc.Invoke(ctx, "/engine_client_proto.RequestHandler/Ping", in, out, c.cc, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// Server API for RequestHandler service

type RequestHandlerServer interface {
	Invoke(context.Context, *Request) (*Response, error)
	SysMan(context.Context, *Request) (*Response, error)
	Auth(context.Context, *AuthRequest) (*Response, error)
	Ping(context.Context, *Request) (*Response, error)
}

func RegisterRequestHandlerServer(s *grpc.Server, srv RequestHandlerServer) {
	s.RegisterService(&_RequestHandler_serviceDesc, srv)
}

func _RequestHandler_Invoke_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(Request)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(RequestHandlerServer).Invoke(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/engine_client_proto.RequestHandler/Invoke",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(RequestHandlerServer).Invoke(ctx, req.(*Request))
	}
	return interceptor(ctx, in, info, handler)
}

func _RequestHandler_SysMan_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(Request)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(RequestHandlerServer).SysMan(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/engine_client_proto.RequestHandler/SysMan",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(RequestHandlerServer).SysMan(ctx, req.(*Request))
	}
	return interceptor(ctx, in, info, handler)
}

func _RequestHandler_Auth_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(AuthRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(RequestHandlerServer).Auth(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/engine_client_proto.RequestHandler/Auth",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(RequestHandlerServer).Auth(ctx, req.(*AuthRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _RequestHandler_Ping_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(Request)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(RequestHandlerServer).Ping(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/engine_client_proto.RequestHandler/Ping",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(RequestHandlerServer).Ping(ctx, req.(*Request))
	}
	return interceptor(ctx, in, info, handler)
}

var _RequestHandler_serviceDesc = grpc.ServiceDesc{
	ServiceName: "engine_client_proto.RequestHandler",
	HandlerType: (*RequestHandlerServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "Invoke",
			Handler:    _RequestHandler_Invoke_Handler,
		},
		{
			MethodName: "SysMan",
			Handler:    _RequestHandler_SysMan_Handler,
		},
		{
			MethodName: "Auth",
			Handler:    _RequestHandler_Auth_Handler,
		},
		{
			MethodName: "Ping",
			Handler:    _RequestHandler_Ping_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "interface.proto",
}

func init() { proto.RegisterFile("interface.proto", fileDescriptor0) }

var fileDescriptor0 = []byte{
	// 237 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0xe2, 0xe2, 0xcf, 0xcc, 0x2b, 0x49,
	0x2d, 0x4a, 0x4b, 0x4c, 0x4e, 0xd5, 0x2b, 0x28, 0xca, 0x2f, 0xc9, 0x17, 0x12, 0x4e, 0xcd, 0x4b,
	0xcf, 0xcc, 0x4b, 0x8d, 0x4f, 0xce, 0xc9, 0x4c, 0xcd, 0x2b, 0x89, 0x07, 0x0b, 0x2a, 0x79, 0x72,
	0x71, 0x3b, 0x96, 0x96, 0x64, 0x04, 0xa5, 0x16, 0x96, 0xa6, 0x16, 0x97, 0x08, 0x49, 0x71, 0x71,
	0x40, 0xa4, 0x3d, 0x53, 0x24, 0x18, 0x15, 0x18, 0x35, 0x78, 0x82, 0xe0, 0x7c, 0x21, 0x39, 0x2e,
	0xae, 0xe4, 0xa2, 0xd4, 0x94, 0xd4, 0xbc, 0x92, 0xcc, 0xc4, 0x1c, 0x09, 0x26, 0xb0, 0x2c, 0x92,
	0x88, 0x92, 0x32, 0x17, 0x3b, 0xcc, 0x18, 0x09, 0x2e, 0xf6, 0x82, 0xc4, 0xca, 0x9c, 0xfc, 0x44,
	0x98, 0x29, 0x30, 0xae, 0x92, 0x15, 0x17, 0x47, 0x50, 0x6a, 0x71, 0x41, 0x7e, 0x5e, 0x71, 0x2a,
	0x6e, 0x55, 0x42, 0x22, 0x5c, 0xac, 0xa9, 0x45, 0x45, 0xf9, 0x45, 0x50, 0x5b, 0x20, 0x1c, 0xa3,
	0x2d, 0x4c, 0x5c, 0x7c, 0x50, 0x1b, 0x3c, 0x12, 0xf3, 0x52, 0x72, 0x52, 0x8b, 0x84, 0xdc, 0xb9,
	0xd8, 0x3c, 0xf3, 0xca, 0xf2, 0xb3, 0x53, 0x85, 0x64, 0xf4, 0xb0, 0x78, 0x4f, 0x0f, 0xaa, 0x5c,
	0x4a, 0x16, 0x87, 0x2c, 0xc4, 0x25, 0x4a, 0x0c, 0x20, 0x83, 0x82, 0x2b, 0x8b, 0x7d, 0x13, 0xf3,
	0x28, 0x35, 0xc8, 0x93, 0x8b, 0x05, 0x14, 0xa0, 0x42, 0x0a, 0x58, 0x15, 0x22, 0x85, 0x35, 0x61,
	0xa3, 0x5c, 0xb9, 0x58, 0x02, 0x32, 0xf3, 0xd2, 0x29, 0x74, 0x51, 0x12, 0x1b, 0x58, 0xc4, 0x18,
	0x10, 0x00, 0x00, 0xff, 0xff, 0x14, 0xf3, 0xa9, 0x71, 0x11, 0x02, 0x00, 0x00,
}
