package client

import (
	"log"

	"github.com/sgokul961/timepeace-api-gateway/pkg/config"
	"github.com/sgokul961/timepeace-api-gateway/pkg/pb"
	"google.golang.org/grpc"
)

type ProductClient struct {
	Client pb.ProductServiceClient
}

func NewProductClient(c config.Config) *ProductClient {
	cc, err := grpc.Dial(c.ProductSvcUrl, grpc.WithInsecure())
	if err != nil {
		log.Fatalln("client connection failiure", err)
	}
	return &ProductClient{
		Client: pb.NewProductServiceClient(cc),
	}
}
