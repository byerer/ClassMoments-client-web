package router

import (
	"ClassMoments-client-web/internal/controller"
	"github.com/gin-gonic/gin"
)

type ClassMomentsAPIRouter struct {
	userController   *controller.UserController
	uploadController *controller.UploadController
}

func NewClassMomentsAPIRouter(
	userController *controller.UserController,
	uploadController *controller.UploadController,
) *ClassMomentsAPIRouter {
	return &ClassMomentsAPIRouter{
		userController:   userController,
		uploadController: uploadController,
	}
}

func (c ClassMomentsAPIRouter) RegisterUnAuthAPIRouter(r *gin.RouterGroup) {
	//user
	r.POST("/user/login", c.userController.UserLogin)
	r.POST("/user/register", c.userController.UserRegister)

	//upload
	r.POST("/upload", c.uploadController.UploadFile)
}
