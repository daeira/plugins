// Code generated by protoc-gen-go.
// source: greeter.proto
// DO NOT EDIT!

/*
Package greeter is a generated protocol buffer package.

It is generated from these files:
	greeter.proto

It has these top-level messages:
	GreeterRequest
	GreeterReply
*/
package greeter

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

type GreeterRequest struct {
	Data string `protobuf:"bytes,1,opt,name=data" json:"data,omitempty"`
}

func (m *GreeterRequest) Reset()         { *m = GreeterRequest{} }
func (m *GreeterRequest) String() string { return proto.CompactTextString(m) }
func (*GreeterRequest) ProtoMessage()    {}

type GreeterReply struct {
	Reply string `protobuf:"bytes,1,opt,name=reply" json:"reply,omitempty"`
}

func (m *GreeterReply) Reset()         { *m = GreeterReply{} }
func (m *GreeterReply) String() string { return proto.CompactTextString(m) }
func (*GreeterReply) ProtoMessage()    {}

// Reference imports to suppress errors if they are not otherwise used.
var _ context.Context
var _ grpc.ClientConn

// Client API for Greeter service

type GreeterClient interface {
	Process(ctx context.Context, in *GreeterRequest, opts ...grpc.CallOption) (*GreeterReply, error)
}

type greeterClient struct {
	cc *grpc.ClientConn
}

func NewGreeterClient(cc *grpc.ClientConn) GreeterClient {
	return &greeterClient{cc}
}

func (c *greeterClient) Process(ctx context.Context, in *GreeterRequest, opts ...grpc.CallOption) (*GreeterReply, error) {
	out := new(GreeterReply)
	err := grpc.Invoke(ctx, "/greeter.Greeter/Process", in, out, c.cc, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// Server API for Greeter service

type GreeterServer interface {
	Process(context.Context, *GreeterRequest) (*GreeterReply, error)
}

func RegisterGreeterServer(s *grpc.Server, srv GreeterServer) {
	s.RegisterService(&_Greeter_serviceDesc, srv)
}

func _Greeter_Process_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error) (interface{}, error) {
	in := new(GreeterRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	out, err := srv.(GreeterServer).Process(ctx, in)
	if err != nil {
		return nil, err
	}
	return out, nil
}

var _Greeter_serviceDesc = grpc.ServiceDesc{
	ServiceName: "greeter.Greeter",
	HandlerType: (*GreeterServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "Process",
			Handler:    _Greeter_Process_Handler,
		},
	},
	Streams: []grpc.StreamDesc{},
}
