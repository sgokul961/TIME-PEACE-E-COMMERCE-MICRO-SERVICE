//go:build wireinject
// +build wireinject

package di

import (
	"github.com/google/wire"
	"github.com/sgokul961/timepeace-api-gateway/pkg/api"
	"github.com/sgokul961/timepeace-api-gateway/pkg/api/handler"
	"github.com/sgokul961/timepeace-api-gateway/pkg/client"
	"github.com/sgokul961/timepeace-api-gateway/pkg/config"
)

func InitializeAPI(c config.Config) (*api.ServerHTTP, error) {
	wire.Build(
		api.NewServerHTTP,

		handler.NewAdminHandler,
		client.NewAdminClient,
		handler.NewUserHandler,
		client.NewUserClient,
		handler.NewProductHandler,
		client.NewProductClient,
	)
	return &api.ServerHTTP{}, nil
}
