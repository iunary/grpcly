// Code generated by protoc-gen-go-grpc. DO NOT EDIT.
// versions:
// - protoc-gen-go-grpc v1.2.0
// - protoc             v3.19.1
// source: proto/reverser.proto

package reverserpb

import (
	context "context"
	grpc "google.golang.org/grpc"
	codes "google.golang.org/grpc/codes"
	status "google.golang.org/grpc/status"
)

// This is a compile-time assertion to ensure that this generated file
// is compatible with the grpc package it is being compiled against.
// Requires gRPC-Go v1.32.0 or later.
const _ = grpc.SupportPackageIsVersion7

// ReverserClient is the client API for Reverser service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://pkg.go.dev/google.golang.org/grpc/?tab=doc#ClientConn.NewStream.
type ReverserClient interface {
	ReverseString(ctx context.Context, opts ...grpc.CallOption) (Reverser_ReverseStringClient, error)
}

type reverserClient struct {
	cc grpc.ClientConnInterface
}

func NewReverserClient(cc grpc.ClientConnInterface) ReverserClient {
	return &reverserClient{cc}
}

func (c *reverserClient) ReverseString(ctx context.Context, opts ...grpc.CallOption) (Reverser_ReverseStringClient, error) {
	stream, err := c.cc.NewStream(ctx, &Reverser_ServiceDesc.Streams[0], "/reverser.Reverser/ReverseString", opts...)
	if err != nil {
		return nil, err
	}
	x := &reverserReverseStringClient{stream}
	return x, nil
}

type Reverser_ReverseStringClient interface {
	Send(*ReverserRequest) error
	CloseAndRecv() (*ReverserResponse, error)
	grpc.ClientStream
}

type reverserReverseStringClient struct {
	grpc.ClientStream
}

func (x *reverserReverseStringClient) Send(m *ReverserRequest) error {
	return x.ClientStream.SendMsg(m)
}

func (x *reverserReverseStringClient) CloseAndRecv() (*ReverserResponse, error) {
	if err := x.ClientStream.CloseSend(); err != nil {
		return nil, err
	}
	m := new(ReverserResponse)
	if err := x.ClientStream.RecvMsg(m); err != nil {
		return nil, err
	}
	return m, nil
}

// ReverserServer is the server API for Reverser service.
// All implementations must embed UnimplementedReverserServer
// for forward compatibility
type ReverserServer interface {
	ReverseString(Reverser_ReverseStringServer) error
	mustEmbedUnimplementedReverserServer()
}

// UnimplementedReverserServer must be embedded to have forward compatible implementations.
type UnimplementedReverserServer struct {
}

func (UnimplementedReverserServer) ReverseString(Reverser_ReverseStringServer) error {
	return status.Errorf(codes.Unimplemented, "method ReverseString not implemented")
}
func (UnimplementedReverserServer) mustEmbedUnimplementedReverserServer() {}

// UnsafeReverserServer may be embedded to opt out of forward compatibility for this service.
// Use of this interface is not recommended, as added methods to ReverserServer will
// result in compilation errors.
type UnsafeReverserServer interface {
	mustEmbedUnimplementedReverserServer()
}

func RegisterReverserServer(s grpc.ServiceRegistrar, srv ReverserServer) {
	s.RegisterService(&Reverser_ServiceDesc, srv)
}

func _Reverser_ReverseString_Handler(srv interface{}, stream grpc.ServerStream) error {
	return srv.(ReverserServer).ReverseString(&reverserReverseStringServer{stream})
}

type Reverser_ReverseStringServer interface {
	SendAndClose(*ReverserResponse) error
	Recv() (*ReverserRequest, error)
	grpc.ServerStream
}

type reverserReverseStringServer struct {
	grpc.ServerStream
}

func (x *reverserReverseStringServer) SendAndClose(m *ReverserResponse) error {
	return x.ServerStream.SendMsg(m)
}

func (x *reverserReverseStringServer) Recv() (*ReverserRequest, error) {
	m := new(ReverserRequest)
	if err := x.ServerStream.RecvMsg(m); err != nil {
		return nil, err
	}
	return m, nil
}

// Reverser_ServiceDesc is the grpc.ServiceDesc for Reverser service.
// It's only intended for direct use with grpc.RegisterService,
// and not to be introspected or modified (even as a copy)
var Reverser_ServiceDesc = grpc.ServiceDesc{
	ServiceName: "reverser.Reverser",
	HandlerType: (*ReverserServer)(nil),
	Methods:     []grpc.MethodDesc{},
	Streams: []grpc.StreamDesc{
		{
			StreamName:    "ReverseString",
			Handler:       _Reverser_ReverseString_Handler,
			ClientStreams: true,
		},
	},
	Metadata: "proto/reverser.proto",
}
