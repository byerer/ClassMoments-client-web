package controller

import (
	"ClassMoments-client-web/internal/base/constant"
	"ClassMoments-client-web/internal/schema"
	"ClassMoments-client-web/internal/service/content"
	"github.com/gin-gonic/gin"
	"net/http"
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
		ctx.JSON(http.StatusBadRequest, gin.H{"error": constant.ErrorBindReqBody})
		return
	}

	resp, err := uc.userService.UserLogin(req)
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	ctx.JSON(http.StatusOK, gin.H{
		"code":    "200",
		"message": "success",
		"data":    resp,
	})
}

func (uc *UserController) UserRegister(ctx *gin.Context) {
	req := &schema.UserRegisterReq{}
	if err := ctx.ShouldBindJSON(req); err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": constant.ErrorBindReqBody})
		return
	}
	err := uc.userService.UserRegister(req)
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	ctx.JSON(http.StatusOK, gin.H{
		"code":    "200",
		"message": "success",
	})
}
