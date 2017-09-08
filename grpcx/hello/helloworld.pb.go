// Code generated by protoc-gen-go. DO NOT EDIT.
// source: helloworld.proto

/*
Package helloworld is a generated protocol buffer package.

It is generated from these files:
	helloworld.proto

It has these top-level messages:
	HelloRequest
	HelloReply
*/
package helloworld

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

type HelloRequest struct {
	ParaFoo string `protobuf:"bytes,1,opt,name=paraFoo" json:"paraFoo,omitempty"`
	ParaBar string `protobuf:"bytes,2,opt,name=paraBar" json:"paraBar,omitempty"`
	ParaHoo int32  `protobuf:"varint,3,opt,name=paraHoo" json:"paraHoo,omitempty"`
}

func (m *HelloRequest) Reset()                    { *m = HelloRequest{} }
func (m *HelloRequest) String() string            { return proto.CompactTextString(m) }
func (*HelloRequest) ProtoMessage()               {}
func (*HelloRequest) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{0} }

func (m *HelloRequest) GetParaFoo() string {
	if m != nil {
		return m.ParaFoo
	}
	return ""
}

func (m *HelloRequest) GetParaBar() string {
	if m != nil {
		return m.ParaBar
	}
	return ""
}

func (m *HelloRequest) GetParaHoo() int32 {
	if m != nil {
		return m.ParaHoo
	}
	return 0
}

type HelloReply struct {
	Message string `protobuf:"bytes,1,opt,name=message" json:"message,omitempty"`
}

func (m *HelloReply) Reset()                    { *m = HelloReply{} }
func (m *HelloReply) String() string            { return proto.CompactTextString(m) }
func (*HelloReply) ProtoMessage()               {}
func (*HelloReply) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{1} }

func (m *HelloReply) GetMessage() string {
	if m != nil {
		return m.Message
	}
	return ""
}

func init() {
	proto.RegisterType((*HelloRequest)(nil), "HelloRequest")
	proto.RegisterType((*HelloReply)(nil), "HelloReply")
}

// Reference imports to suppress errors if they are not otherwise used.
var _ context.Context
var _ grpc.ClientConn

// This is a compile-time assertion to ensure that this generated file
// is compatible with the grpc package it is being compiled against.
const _ = grpc.SupportPackageIsVersion4

// Client API for Greeter service

type GreeterClient interface {
	SayHello(ctx context.Context, in *HelloRequest, opts ...grpc.CallOption) (*HelloReply, error)
}

type greeterClient struct {
	cc *grpc.ClientConn
}

func NewGreeterClient(cc *grpc.ClientConn) GreeterClient {
	return &greeterClient{cc}
}

func (c *greeterClient) SayHello(ctx context.Context, in *HelloRequest, opts ...grpc.CallOption) (*HelloReply, error) {
	out := new(HelloReply)
	err := grpc.Invoke(ctx, "/Greeter/SayHello", in, out, c.cc, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// Server API for Greeter service

type GreeterServer interface {
	SayHello(context.Context, *HelloRequest) (*HelloReply, error)
}

func RegisterGreeterServer(s *grpc.Server, srv GreeterServer) {
	s.RegisterService(&_Greeter_serviceDesc, srv)
}

func _Greeter_SayHello_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(HelloRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(GreeterServer).SayHello(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/Greeter/SayHello",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(GreeterServer).SayHello(ctx, req.(*HelloRequest))
	}
	return interceptor(ctx, in, info, handler)
}

var _Greeter_serviceDesc = grpc.ServiceDesc{
	ServiceName: "Greeter",
	HandlerType: (*GreeterServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "SayHello",
			Handler:    _Greeter_SayHello_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "helloworld.proto",
}

func init() { proto.RegisterFile("helloworld.proto", fileDescriptor0) }

var fileDescriptor0 = []byte{
	// 161 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0xe2, 0x12, 0xc8, 0x48, 0xcd, 0xc9,
	0xc9, 0x2f, 0xcf, 0x2f, 0xca, 0x49, 0xd1, 0x2b, 0x28, 0xca, 0x2f, 0xc9, 0x57, 0x8a, 0xe1, 0xe2,
	0xf1, 0x00, 0x89, 0x05, 0xa5, 0x16, 0x96, 0xa6, 0x16, 0x97, 0x08, 0x49, 0x70, 0xb1, 0x17, 0x24,
	0x16, 0x25, 0xba, 0xe5, 0xe7, 0x4b, 0x30, 0x2a, 0x30, 0x6a, 0x70, 0x06, 0xc1, 0xb8, 0x30, 0x19,
	0xa7, 0xc4, 0x22, 0x09, 0x26, 0x84, 0x8c, 0x53, 0x62, 0x11, 0x4c, 0xc6, 0x23, 0x3f, 0x5f, 0x82,
	0x59, 0x81, 0x51, 0x83, 0x35, 0x08, 0xc6, 0x55, 0x52, 0xe3, 0xe2, 0x82, 0x9a, 0x5e, 0x90, 0x53,
	0x09, 0x52, 0x97, 0x9b, 0x5a, 0x5c, 0x9c, 0x98, 0x9e, 0x0a, 0x33, 0x1b, 0xca, 0x35, 0x32, 0xe6,
	0x62, 0x77, 0x2f, 0x4a, 0x4d, 0x2d, 0x49, 0x2d, 0x12, 0xd2, 0xe0, 0xe2, 0x08, 0x4e, 0xac, 0x04,
	0xeb, 0x12, 0xe2, 0xd5, 0x43, 0x76, 0x9b, 0x14, 0xb7, 0x1e, 0xc2, 0x30, 0x25, 0x86, 0x24, 0x36,
	0xb0, 0x0f, 0x8c, 0x01, 0x01, 0x00, 0x00, 0xff, 0xff, 0x16, 0x53, 0xdd, 0x58, 0xd5, 0x00, 0x00,
	0x00,
}
