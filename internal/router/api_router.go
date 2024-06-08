package router

import (
	"ClassMoments-client-web/internal/controller"
	"github.com/gin-gonic/gin"
)

type ClassMomentsAPIRouter struct {
	userController   *controller.UserController
	uploadController *controller.UploadController
	likeController   *controller.LikeController
	momentController *controller.MomentController
}

func NewClassMomentsAPIRouter(
	userController *controller.UserController,
	uploadController *controller.UploadController,
	likeController *controller.LikeController,
	momentController *controller.MomentController,
) *ClassMomentsAPIRouter {
	return &ClassMomentsAPIRouter{
		userController:   userController,
		uploadController: uploadController,
		likeController:   likeController,
		momentController: momentController,
	}
}

func (c ClassMomentsAPIRouter) RegisterUnAuthAPIRouter(r *gin.RouterGroup) {
	//user
	r.POST("/user/login", c.userController.UserLogin)
	r.POST("/user/register", c.userController.UserRegister)

	//upload
	r.POST("/upload", c.uploadController.UploadFile)

	//moment
	r.POST("/moment", c.momentController.AddMoment)

	//like
	r.POST("/like", c.likeController.Like)
	r.DELETE("/like", c.likeController.UnLike)
}
