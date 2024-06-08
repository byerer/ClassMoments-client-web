package like

import (
	"ClassMoments-client-web/internal/base/data"
	"ClassMoments-client-web/internal/entity"
	"ClassMoments-client-web/internal/service/like"
)

type likeRepo struct {
	data *data.Data
}

func NewLikeRepo(data *data.Data) like.LikeRepo {
	return &likeRepo{data: data}
}

func (lr *likeRepo) AddLike(like *entity.Like) error {
	result := lr.data.DB.Create(like)
	if result.Error != nil {
		return result.Error
	}
	return nil
}

func (lr *likeRepo) DeleteLike(likerID, momentID uint) error {
	result := lr.data.DB.Where("liker_id = ? AND moment_id = ?", likerID, momentID).Delete(&entity.Like{})
	if result.Error != nil {
		return result.Error
	}
	return nil
}
