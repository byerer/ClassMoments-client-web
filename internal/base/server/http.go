package server

import (
	"ClassMoments-client-web/internal/router"
	"github.com/gin-gonic/gin"
)

func NewHTTPServer(
	apiRouter *router.ClassMomentsAPIRouter,
	staticRouter *router.StaticRouter,
) *gin.Engine {

	r := gin.Default()
	r.GET("/ping", func(c *gin.Context) {
		c.JSON(200, gin.H{
			"message": "pong",
		})
	})

	static := r.Group("")
	staticRouter.RegisterStaticRouter(static)

	api := r.Group("/api/v1")
	apiRouter.RegisterUnAuthAPIRouter(api)

	return r
}
