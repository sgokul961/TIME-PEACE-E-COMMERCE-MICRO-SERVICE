package api

import (
	"log"
	"net"

	"github.com/sgokul961/time-peace-admin-service/pkg/api/service"
	"github.com/sgokul961/time-peace-admin-service/pkg/config"
	"github.com/sgokul961/time-peace-admin-service/pkg/pb"
	"google.golang.org/grpc"
)

type ServerHTTP struct {
	engine *grpc.Server
}

func NewServerHTTP(adminService *service.AdminService) *ServerHTTP {
	engine := grpc.NewServer()

	pb.RegisterAdminServiceServer(engine, adminService)
	return &ServerHTTP{engine: engine}
}

func (s *ServerHTTP) Start(c config.Config) {
	lis, err := net.Listen("tcp", c.Port)

	if err != nil {
		log.Fatalln("problem while listening ", err)

	}
	if err := s.engine.Serve(lis); err != nil {
		log.Fatalln("failed to serve ", err)
	}
}
