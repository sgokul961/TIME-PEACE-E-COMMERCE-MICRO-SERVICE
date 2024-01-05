package api

import (
	"fmt"
	"log"
	"net"

	"github.com/sgokul961/timepeace-user-service/pkg/api/service"
	"github.com/sgokul961/timepeace-user-service/pkg/config"
	"github.com/sgokul961/timepeace-user-service/pkg/pb"
	"google.golang.org/grpc"
)

type ServerHTTP struct {
	engine *grpc.Server
}

func NewServerHTTP(userservice *service.UserService) *ServerHTTP {
	engine := grpc.NewServer()
	fmt.Println("user server")
	pb.RegisterUserServiceServer(engine, userservice)
	return &ServerHTTP{
		engine: engine,
	}
}

func (s *ServerHTTP) Start(c config.Config) error {
	lis, err := net.Listen("tcp", c.Port)

	if err != nil {
		log.Fatalln("failed to listen", err)
		return err
	}
	if err := s.engine.Serve(lis); err != nil {
		log.Fatalln("failed to serve", err)
		return err
	}
	return nil
}
