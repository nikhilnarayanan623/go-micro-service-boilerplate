//go:build wireinject
// +build wireinject

package di

import (
	"employee-service/pkg/api"
	"employee-service/pkg/api/service"
	"employee-service/pkg/config"
	"employee-service/pkg/service/random"
	"employee-service/pkg/usecase"

	"github.com/google/wire"
)

func InitializeAPI(cfg config.Config) (*api.Server, error) {

	wire.Build(
		random.NewRandomGenerator,
		usecase.NewEmployeeUseCase,
		service.NewAuthServiceServer,
		api.NewServerGRPC,
	)

	return &api.Server{}, nil
}
