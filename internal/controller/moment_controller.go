package controller

import (
	"ClassMoments-client-web/internal/base/constant"
	"ClassMoments-client-web/internal/schema"
	"ClassMoments-client-web/internal/service/moment"
	"github.com/gin-gonic/gin"
	"go.uber.org/zap"
	"net/http"
)

type MomentController struct {
	momentService moment.MomentService
	logger        *zap.Logger
}

func NewMomentController(
	momentService moment.MomentService,
	logger *zap.Logger) *MomentController {
	return &MomentController{
		momentService: momentService,
		logger:        logger,
	}
}

// AddMoment add moment
func (mc *MomentController) AddMoment(ctx *gin.Context) {
	req := &schema.AddMomentReq{}
	if err := ctx.ShouldBindJSON(req); err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": constant.ErrorBindReqBody})
		return
	}
	resp, err := mc.momentService.AddMoment(req)
	if err != nil {
		ctx.JSON(500, gin.H{"error": err.Error()})
		return
	}
	ctx.JSON(200, gin.H{
		"code":    "200",
		"message": "success",
		"data":    resp,
	})
}

func (mc *MomentController) GetMomentList(ctx *gin.Context) {
	req := &struct {
		ClassID uint `json:"class_id"`
	}{}
	if err := ctx.ShouldBindQuery(req); err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": constant.ErrorBindReqBody})
		return
	}
	mc.logger.Info("GetMomentList", zap.Any("req", req))
	resp, err := mc.momentService.GetMomentList(req.ClassID)
	if err != nil {
		ctx.JSON(500, gin.H{"error": err.Error()})
		return
	}
	ctx.JSON(200, gin.H{
		"code":    "200",
		"message": "success",
		"data":    resp,
	})
}
