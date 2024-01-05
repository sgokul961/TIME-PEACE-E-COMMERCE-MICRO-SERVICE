package client

import (
	"log"

	"github.com/sgokul961/timepeace-api-gateway/pkg/config"
	"github.com/sgokul961/timepeace-api-gateway/pkg/pb"
	"google.golang.org/grpc"
)

type AdminClient struct {
	Client pb.AdminServiceClient
}

func NewAdminClient(c config.Config) *AdminClient {
	cc, err := grpc.Dial(c.AdminSvcUrl, grpc.WithInsecure())

	if err != nil {
		log.Fatalln("client connection failed for admin service ", err)

	}
	return &AdminClient{
		Client: pb.NewAdminServiceClient(cc),
	}
}
