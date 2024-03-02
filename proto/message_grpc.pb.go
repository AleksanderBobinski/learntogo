// Code generated by protoc-gen-go-grpc. DO NOT EDIT.
// versions:
// - protoc-gen-go-grpc v1.2.0
// - protoc             v4.25.2
// source: proto/message.proto

package proto

import (
	context "context"
	grpc "google.golang.org/grpc"
	codes "google.golang.org/grpc/codes"
	status "google.golang.org/grpc/status"
	emptypb "google.golang.org/protobuf/types/known/emptypb"
)

// This is a compile-time assertion to ensure that this generated file
// is compatible with the grpc package it is being compiled against.
// Requires gRPC-Go v1.32.0 or later.
const _ = grpc.SupportPackageIsVersion7

// SenderClient is the client API for Sender service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://pkg.go.dev/google.golang.org/grpc/?tab=doc#ClientConn.NewStream.
type SenderClient interface {
	Send(ctx context.Context, in *Msg, opts ...grpc.CallOption) (*emptypb.Empty, error)
}

type senderClient struct {
	cc grpc.ClientConnInterface
}

func NewSenderClient(cc grpc.ClientConnInterface) SenderClient {
	return &senderClient{cc}
}

func (c *senderClient) Send(ctx context.Context, in *Msg, opts ...grpc.CallOption) (*emptypb.Empty, error) {
	out := new(emptypb.Empty)
	err := c.cc.Invoke(ctx, "/main.Sender/Send", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// SenderServer is the server API for Sender service.
// All implementations must embed UnimplementedSenderServer
// for forward compatibility
type SenderServer interface {
	Send(context.Context, *Msg) (*emptypb.Empty, error)
	mustEmbedUnimplementedSenderServer()
}

// UnimplementedSenderServer must be embedded to have forward compatible implementations.
type UnimplementedSenderServer struct {
}

func (UnimplementedSenderServer) Send(context.Context, *Msg) (*emptypb.Empty, error) {
	return nil, status.Errorf(codes.Unimplemented, "method Send not implemented")
}
func (UnimplementedSenderServer) mustEmbedUnimplementedSenderServer() {}

// UnsafeSenderServer may be embedded to opt out of forward compatibility for this service.
// Use of this interface is not recommended, as added methods to SenderServer will
// result in compilation errors.
type UnsafeSenderServer interface {
	mustEmbedUnimplementedSenderServer()
}

func RegisterSenderServer(s grpc.ServiceRegistrar, srv SenderServer) {
	s.RegisterService(&Sender_ServiceDesc, srv)
}

func _Sender_Send_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(Msg)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(SenderServer).Send(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/main.Sender/Send",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(SenderServer).Send(ctx, req.(*Msg))
	}
	return interceptor(ctx, in, info, handler)
}

// Sender_ServiceDesc is the grpc.ServiceDesc for Sender service.
// It's only intended for direct use with grpc.RegisterService,
// and not to be introspected or modified (even as a copy)
var Sender_ServiceDesc = grpc.ServiceDesc{
	ServiceName: "main.Sender",
	HandlerType: (*SenderServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "Send",
			Handler:    _Sender_Send_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "proto/message.proto",
}
