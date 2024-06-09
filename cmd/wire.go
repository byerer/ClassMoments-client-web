//go:build wireinject
// +build wireinject

package cmd

import (
	"ClassMoments-client-web/internal/base/log"
	"ClassMoments-client-web/internal/base/server"
	"ClassMoments-client-web/internal/controller"
	"ClassMoments-client-web/internal/repo"
	"ClassMoments-client-web/internal/router"
	"ClassMoments-client-web/internal/service"
	"github.com/google/wire"
)

func InitApplication() (*Application, error) {
	panic(wire.Build(
		log.NewLogger,
		repo.ProviderSetRepo,
		server.ProviderSetServer,
		router.ProviderSetRouter,
		controller.ProviderSetController,
		service.ProviderSetService,
		NewApplication,
	))
}
