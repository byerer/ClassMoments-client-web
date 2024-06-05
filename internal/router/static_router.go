package router

import "github.com/gin-gonic/gin"

type StaticRouter struct {
}

func NewStaticRouter() *StaticRouter {
	return &StaticRouter{}
}

func (st *StaticRouter) RegisterStaticRouter(r *gin.RouterGroup) {
	r.Static("/uploads", "/uploads")
}
