package main

import (
	"context"
	"fmt"
	"net"
	"os"

	"github.com/AleksanderBobinski/learntogo/v2/proto"
	"google.golang.org/grpc"
	"google.golang.org/protobuf/types/known/emptypb"
)

type senderServer struct {
	proto.UnimplementedSenderServer
	Stdout *os.File
}

func newServer(stdout *os.File) senderServer {
	return senderServer{Stdout: stdout}
}

func (s senderServer) Send(_ context.Context, msg *proto.Msg) (*emptypb.Empty, error) {
	fmt.Fprintf(s.Stdout, "%s", msg.Msg)
	return &emptypb.Empty{}, nil
}

func serve(stdout *os.File, port int) int {
	lis, _ := net.Listen("tcp", fmt.Sprintf("localhost:%d", port))
	grpcServer := grpc.NewServer()
	proto.RegisterSenderServer(grpcServer, newServer(stdout))
	err := grpcServer.Serve(lis)
	fmt.Fprintf(os.Stderr, "gRPC server failed with: %s", err.Error())
	return -1
}

func main() {
	os.Exit(serve(os.Stdout, 8080))
}
