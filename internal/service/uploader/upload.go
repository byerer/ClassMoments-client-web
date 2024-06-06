package uploader

import "github.com/gin-gonic/gin"

type UploaderService interface {
	UploadFile(ctx *gin.Context) (string, error)
}

type uploaderService struct {
}

func NewUploaderService() UploaderService {
	return &uploaderService{}
}

func (us *uploaderService) UploadFile(ctx *gin.Context) (url string, err error) {
	file, _ := ctx.FormFile("file")
	dst := "uploads/" + file.Filename
	err = ctx.SaveUploadedFile(file, dst)
	if err != nil {
		return "", err
	}
	return dst, nil
}
