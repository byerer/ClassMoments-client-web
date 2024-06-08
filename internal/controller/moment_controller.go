package controller

import (
	"ClassMoments-client-web/internal/schema"
	"ClassMoments-client-web/internal/service/moment"
	"github.com/gin-gonic/gin"
)

type MomentController struct {
	momentService moment.MomentService
}

func NewMomentController(momentService moment.MomentService) *MomentController {
	return &MomentController{momentService: momentService}
}

// AddMoment add moment
func (mc *MomentController) AddMoment(ctx *gin.Context) {
	req := &schema.AddMomentReq{}
	if err := ctx.BindJSON(req); err != nil {
		ctx.JSON(400, gin.H{"error": err.Error()})
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
