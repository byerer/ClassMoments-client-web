package controller

import (
	"ClassMoments-client-web/internal/service/media"
	"github.com/gin-gonic/gin"
)

type MediaController struct {
	mediaService media.MediaService
}

func NewMediaController(mediaService media.MediaService) *MediaController {
	return &MediaController{
		mediaService: mediaService,
	}
}

func (mc *MediaController) AddMedia(ctx *gin.Context) {
	media := &struct {
		MomentID  uint   `json:"momentID"`
		MediaType string `json:"mediaType"`
		URL       string `json:"url"`
	}{}
	err := ctx.ShouldBindJSON(media)
	if err != nil {
		ctx.JSON(400, gin.H{"msg": "bind failed"})
		return
	}
	err = mc.mediaService.AddMedia(media.MomentID, media.MediaType, media.URL)
	if err != nil {
		ctx.JSON(500, gin.H{"error": err.Error()})
		return
	}
	ctx.JSON(200, gin.H{"code": 200})
}
