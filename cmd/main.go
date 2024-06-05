package cmd

import "github.com/gin-gonic/gin"

type Application struct {
	Engine *gin.Engine
}

func NewApplication(engine *gin.Engine) *Application {
	return &Application{
		Engine: engine,
	}
}

func (a *Application) Run() error {
	return a.Engine.Run()
}
