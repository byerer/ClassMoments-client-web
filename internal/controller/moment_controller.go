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
	mc.logger.Info("AddMoment", zap.Any("req", req))
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
		ClassID uint `form:"classID"`
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

func (mc *MomentController) GetMomentDetail(context *gin.Context) {
	req := &struct {
		MomentID uint `form:"momentID"`
	}{}
	if err := context.ShouldBindQuery(req); err != nil {
		context.JSON(http.StatusBadRequest, gin.H{"error": constant.ErrorBindReqBody})
		return
	}
	resp, err := mc.momentService.GetMomentDetail(req.MomentID)
	mc.logger.Info("GetMomentDetail", zap.Any("req", req))
	mc.logger.Info("GetMomentDetail", zap.Any("resp", resp))
	mc.logger.Info("GetMomentDetail", zap.Any("err", err))
	if err != nil {
		context.JSON(500, gin.H{"error": err.Error()})
		return
	}
	context.JSON(200, gin.H{
		"code":    "200",
		"message": "success",
		"data":    resp,
	})
}
