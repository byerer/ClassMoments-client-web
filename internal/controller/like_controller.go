package controller

import (
	"ClassMoments-client-web/internal/schema"
	"ClassMoments-client-web/internal/service/like"
	"github.com/gin-gonic/gin"
)

type LikeController struct {
	likeService like.LikeService
}

func NewLikeController(likeService like.LikeService) *LikeController {
	return &LikeController{
		likeService: likeService,
	}
}
func (controller *LikeController) Like(ctx *gin.Context) {
	req := &schema.LikeReq{}
	err := ctx.ShouldBindJSON(req)
	if err != nil {
		ctx.JSON(400, gin.H{"error": err.Error()})
		return
	}
	resp, err := controller.likeService.Like(req)
	if err != nil {
		ctx.JSON(500, gin.H{"error": err.Error()})
		return
	}
	ctx.JSON(200, gin.H{
		"msg":  "success",
		"data": resp,
	})
}

func (controller *LikeController) UnLike(ctx *gin.Context) {
	unlikeInfo := &struct {
		LikerID  uint `json:"liker_id"`
		MomentID uint `json:"moment_id"`
	}{}
	err := ctx.ShouldBind(&unlikeInfo)
	if err != nil {
		ctx.JSON(400, gin.H{"msg": "bind failed"})
		return
	}
	err = controller.likeService.UnLike(unlikeInfo.LikerID, unlikeInfo.MomentID)
	if err != nil {
		ctx.JSON(500, gin.H{"error": err.Error()})
		return
	}
	ctx.JSON(200, gin.H{"msg": "success"})
}
