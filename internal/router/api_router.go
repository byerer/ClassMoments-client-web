package router

import (
	"ClassMoments-client-web/internal/controller"
	"github.com/gin-gonic/gin"
)

type ClassMomentsAPIRouter struct {
	userController   *controller.UserController
	uploadController *controller.UploadController
	likeController   *controller.LikeController
}

func NewClassMomentsAPIRouter(
	userController *controller.UserController,
	uploadController *controller.UploadController,
	likeController *controller.LikeController,
) *ClassMomentsAPIRouter {
	return &ClassMomentsAPIRouter{
		userController:   userController,
		uploadController: uploadController,
		likeController:   likeController,
	}
}

func (c ClassMomentsAPIRouter) RegisterUnAuthAPIRouter(r *gin.RouterGroup) {
	//user
	r.POST("/user/login", c.userController.UserLogin)
	r.POST("/user/register", c.userController.UserRegister)

	//upload
	r.POST("/upload", c.uploadController.UploadFile)

	//like
	r.POST("/like", c.likeController.Like)
	r.DELETE("/like", c.likeController.UnLike)
}
