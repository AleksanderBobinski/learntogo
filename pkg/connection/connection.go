package connection

import (
	"context"

	"github.com/AleksanderBobinski/learntogo/v2/proto"
	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials/insecure"
)

func Send(address string, msg string) error {
	conn, err := grpc.Dial(address, grpc.WithTransportCredentials(insecure.NewCredentials()))
	if err != nil {
		return err
	}
	defer conn.Close()
	c := proto.NewSenderClient(conn)
	_, err = c.Send(context.Background(), &proto.Msg{Msg: msg})
	return err
}
