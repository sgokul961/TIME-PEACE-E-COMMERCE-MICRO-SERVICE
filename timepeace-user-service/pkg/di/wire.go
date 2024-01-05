//go:build wireinject
// +build wireinject

package di

import (
	"github.com/google/wire"
	"github.com/sgokul961/timepeace-user-service/pkg/api"
	"github.com/sgokul961/timepeace-user-service/pkg/api/service"
	"github.com/sgokul961/timepeace-user-service/pkg/config"
	"github.com/sgokul961/timepeace-user-service/pkg/db"
	"github.com/sgokul961/timepeace-user-service/pkg/repository"
)

func InitilizeApi(c config.Config) (*api.ServerHTTP, error) {
	wire.Build(db.ConnectDataBase,
		repository.NewUserRepository,
		service.NewUserService,
		api.NewServerHTTP)
	return &api.ServerHTTP{}, nil
}
