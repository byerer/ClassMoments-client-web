package cmd

import (
	"github.com/gin-gonic/gin"
	"go.uber.org/zap"
)

type Application struct {
	Logger *zap.Logger
	Engine *gin.Engine
}

func NewApplication(
	engine *gin.Engine,
	logger *zap.Logger,
) *Application {
	return &Application{
		Engine: engine,
		Logger: logger,
	}
}

func (a *Application) Run() error {
	return a.Engine.Run()
}
