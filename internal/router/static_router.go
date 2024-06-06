package router

import (
	"github.com/gin-gonic/gin"
	"net/http"
)

type StaticRouter struct {
}

func NewStaticRouter() *StaticRouter {
	return &StaticRouter{}
}

func (st *StaticRouter) RegisterStaticRouter(r *gin.RouterGroup) {

	r.StaticFS("/uploads", http.Dir("uploads"))
}
