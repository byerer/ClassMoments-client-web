//go:build wireinject
// +build wireinject

package cmd

import (
	"ClassMoments-client-web/internal/base/server"
	"ClassMoments-client-web/internal/controller"
	"ClassMoments-client-web/internal/repo"
	"ClassMoments-client-web/internal/router"
	"ClassMoments-client-web/internal/service"
	"github.com/google/wire"
)

func InitApplication() (*Application, error) {
	panic(wire.Build(
		server.ProviderSetServer,
		router.ProviderSetRouter,
		repo.ProviderSetRepo,
		controller.ProviderSetController,
		service.ProviderSetService,
		NewApplication,
	))
}
