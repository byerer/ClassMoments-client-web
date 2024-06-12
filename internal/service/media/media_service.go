package media

import "ClassMoments-client-web/internal/entity"

type MediaRepo interface {
	AddMedia(media *entity.Media) error
}

type MediaService interface {
	AddMedia(MomentID uint, MediaType, URL string) error
}

type mediaService struct {
	mediaRepo MediaRepo
}

func NewMediaService(mediaRepo MediaRepo) MediaService {
	return &mediaService{
		mediaRepo: mediaRepo,
	}
}

func (ms *mediaService) AddMedia(MomentID uint, MediaType, URL string) error {
	media := &entity.Media{
		MomentID:  MomentID,
		MediaType: MediaType,
		URL:       URL,
	}
	err := ms.mediaRepo.AddMedia(media)
	if err != nil {
		return err
	}
	return nil
}
