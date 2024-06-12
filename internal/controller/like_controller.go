package controller

import (
	"ClassMoments-client-web/internal/schema"
	"ClassMoments-client-web/internal/service/like"
	"github.com/gin-gonic/gin"
	"go.uber.org/zap"
	"net/http"
)

type LikeController struct {
	likeService like.LikeService
	logger      *zap.Logger
}

func NewLikeController(
	likeService like.LikeService,
	logger *zap.Logger,
) *LikeController {
	return &LikeController{
		likeService: likeService,
		logger:      logger,
	}
}

func (lc *LikeController) Like(ctx *gin.Context) {
	req := &schema.LikeReq{}
	err := ctx.ShouldBindJSON(req)
	lc.logger.Info("like request", zap.Any("req", req))
	if err != nil {
		lc.logger.Error("bind json failed", zap.Error(err))
		ctx.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	resp, err := lc.likeService.Like(req)
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	ctx.JSON(http.StatusOK, gin.H{
		"code": 200,
		"data": resp,
	})
}

func (lc *LikeController) UnLike(ctx *gin.Context) {
	unlikeInfo := &struct {
		LikerID  uint `json:"liker_id"`
		MomentID uint `json:"moment_id"`
	}{}
	err := ctx.ShouldBind(&unlikeInfo)
	if err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"msg": "bind failed"})
		return
	}
	err = lc.likeService.UnLike(unlikeInfo.LikerID, unlikeInfo.MomentID)
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	ctx.JSON(http.StatusOK, gin.H{"msg": "success"})
}
