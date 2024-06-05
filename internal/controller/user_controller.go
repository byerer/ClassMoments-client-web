package controller

import (
	"ClassMoments-client-web/internal/schema"
	"ClassMoments-client-web/internal/service/content"
	"github.com/gin-gonic/gin"
)

type UserController struct {
	userService *content.UserService
}

func NewUserController(
	userService *content.UserService) *UserController {
	return &UserController{
		userService: userService,
	}
}

func (uc *UserController) UserLogin(ctx *gin.Context) {
	req := &schema.UserLoginReq{}
	if err := ctx.ShouldBindJSON(req); err != nil {
		ctx.JSON(400, gin.H{"error": err.Error()})
		return
	}
	ctx.JSON(200, gin.H{"message": "success"})
}

func (uc *UserController) UserRegister(ctx *gin.Context) {
	req := &schema.UserRegisterReq{}
	if err := ctx.ShouldBindJSON(req); err != nil {
		ctx.JSON(400, gin.H{"error": err.Error()})
		return
	}
	ctx.JSON(200, gin.H{"message": "success"})
}
