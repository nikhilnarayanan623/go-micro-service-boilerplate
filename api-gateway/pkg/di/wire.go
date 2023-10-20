//go:build wireinject
// +build wireinject

package di

import (
	"api-gateway/pkg/api"
	"api-gateway/pkg/api/handler"
	"api-gateway/pkg/api/routes"
	"api-gateway/pkg/client"
	"api-gateway/pkg/config"

	"github.com/google/wire"
)

func InitializeAPI(cfg config.Config) (*api.Server, error) {

	wire.Build(
		client.NewAuthServiceClient,
		handler.NewAuthHandler,

		handler.NewStudentHandler,

		routes.NewGinRouter,

		api.NewServerHTTP,
	)

	return &api.Server{}, nil
}