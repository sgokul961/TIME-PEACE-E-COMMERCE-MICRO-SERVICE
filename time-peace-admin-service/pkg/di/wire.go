//go:build wireinject
// +build wireinject

package di

import (
	"github.com/google/wire"
	"github.com/sgokul961/time-peace-admin-service/pkg/api"
	"github.com/sgokul961/time-peace-admin-service/pkg/api/service"
	"github.com/sgokul961/time-peace-admin-service/pkg/config"
	"github.com/sgokul961/time-peace-admin-service/pkg/db"
	"github.com/sgokul961/time-peace-admin-service/pkg/repository"
)

func InitializeAPI(c config.Config) (*api.ServerHTTP, error) {
	wire.Build(db.ConnectDataBase,
		repository.NewAdminRepository,
		service.NewAdminService,
		api.NewServerHTTP)
	return &api.ServerHTTP{}, nil
}
