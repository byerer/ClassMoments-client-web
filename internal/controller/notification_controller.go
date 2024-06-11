package controller

import (
	"ClassMoments-client-web/internal/base/constant"
	"ClassMoments-client-web/internal/schema"
	"ClassMoments-client-web/internal/service/notification"
	"github.com/gin-gonic/gin"
	"net/http"
)

type NotificationController struct {
	notificationService notification.NotificationService
}

func NewNotificationController(
	notificationService notification.NotificationService,
) *NotificationController {
	return &NotificationController{
		notificationService: notificationService,
	}
}

// AddNotification add notification
func (nc *NotificationController) AddNotification(ctx *gin.Context) {
	req := &schema.NotificationReq{}
	if err := ctx.ShouldBindJSON(req); err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": constant.ErrorBindReqBody})
		return
	}
	err := nc.notificationService.AddNotification(req)
	if err != nil {
		ctx.JSON(500, gin.H{"error": err.Error()})
		return
	}
	ctx.JSON(200, gin.H{
		"code":    "200",
		"message": "success",
	})
}
