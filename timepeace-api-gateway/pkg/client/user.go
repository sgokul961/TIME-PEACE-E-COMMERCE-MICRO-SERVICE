package client

import (
	"log"

	"github.com/sgokul961/timepeace-api-gateway/pkg/config"
	"github.com/sgokul961/timepeace-api-gateway/pkg/pb"
	"google.golang.org/grpc"
)

type UserClient struct {
	Client pb.UserServiceClient
}

func NewUserClient(c config.Config) *UserClient {
	cc, err := grpc.Dial(c.UserSvcUrl, grpc.WithInsecure())

	if err != nil {
		log.Fatalln("client connection failed for admin service ", err)

	}
	return &UserClient{
		Client: pb.NewUserServiceClient(cc),
	}
}
