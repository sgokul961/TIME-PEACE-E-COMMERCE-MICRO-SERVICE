package api

import (
	"log"
	"net"

	"github.com/sgokul961/timepeace-product-service/pkg/api/service"
	"github.com/sgokul961/timepeace-product-service/pkg/config"
	"github.com/sgokul961/timepeace-product-service/pkg/pb"
	"google.golang.org/grpc"
)

type ServerHTTP struct {
	engine *grpc.Server
}

func NewServerHTTP(productService *service.ProductService) *ServerHTTP {
	engine := grpc.NewServer()

	pb.RegisterProductServiceServer(engine, productService)
	return &ServerHTTP{
		engine: engine,
	}
}

func (s *ServerHTTP) Start(c config.Config) {
	lis, err := net.Listen("tcp", c.Port)
	if err != nil {
		log.Fatalln("failed to listen", err)
	}
	if err = s.engine.Serve(lis); err != nil {
		log.Fatalln("failed to serve", err)
	}
}
