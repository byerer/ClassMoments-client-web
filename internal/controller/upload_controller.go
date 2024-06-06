package controller

import (
	"ClassMoments-client-web/internal/service/uploader"
	"github.com/gin-gonic/gin"
)

type UploadController struct {
	uploaderService uploader.UploaderService
}

func NewUploadController(uploaderService uploader.UploaderService) *UploadController {
	return &UploadController{
		uploaderService: uploaderService,
	}
}

func (uc *UploadController) UploadFile(ctx *gin.Context) {
	url, err := uc.uploaderService.UploadFile(ctx)
	if err != nil {
		ctx.JSON(500, gin.H{"error": err.Error()})
		return
	}
	ctx.JSON(200, gin.H{"url": url})
}
