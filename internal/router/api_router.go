package router

import (
	"ClassMoments-client-web/internal/controller"
	"github.com/gin-gonic/gin"
)

type ClassMomentsAPIRouter struct {
	userController *controller.UserController
}

func NewClassMomentsAPIRouter(userController *controller.UserController) *ClassMomentsAPIRouter {
	return &ClassMomentsAPIRouter{
		userController: userController,
	}
}

func (c ClassMomentsAPIRouter) RegisterUnAuthAPIRouter(r *gin.RouterGroup) {
	//user
	r.POST("/user/login", c.userController.UserLogin)
	r.POST("/user/register", c.userController.UserRegister)
}
