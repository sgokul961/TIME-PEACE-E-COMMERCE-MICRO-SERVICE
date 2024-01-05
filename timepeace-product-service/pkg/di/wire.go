//go:build wireinject
// +build wireinject

package di

import (
	"github.com/google/wire"
	"github.com/sgokul961/timepeace-product-service/pkg/api"
	"github.com/sgokul961/timepeace-product-service/pkg/api/service"
	"github.com/sgokul961/timepeace-product-service/pkg/config"
	"github.com/sgokul961/timepeace-product-service/pkg/db"
	"github.com/sgokul961/timepeace-product-service/pkg/repository"
)

func InitializeAPI(c config.Config) (*api.ServerHTTP, error) {
	wire.Build(db.ConnectToDataBase,
		repository.NewProductRepository,
		service.NewProductService,
		api.NewServerHTTP)

	return &api.ServerHTTP{}, nil
}
