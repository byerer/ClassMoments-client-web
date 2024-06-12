package media

import (
	"ClassMoments-client-web/internal/base/data"
	"ClassMoments-client-web/internal/entity"
	"ClassMoments-client-web/internal/service/media"
)

type mediaRepo struct {
	data *data.Data
}

func NewMediaRepo(data *data.Data) media.MediaRepo {
	return &mediaRepo{
		data: data,
	}
}

func (mr *mediaRepo) AddMedia(media *entity.Media) error {
	result := mr.data.DB.Exec("CALL Insert_media(?,?,?)", media.MomentID, media.MediaType, media.URL)
	if result.Error != nil {
		return result.Error
	}
	return nil
}
