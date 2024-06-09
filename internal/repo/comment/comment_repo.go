package comment

import (
	"ClassMoments-client-web/internal/base/data"
	"ClassMoments-client-web/internal/entity"
	"ClassMoments-client-web/internal/service/comment"
	"go.uber.org/zap"
)

type commentRepo struct {
	data   *data.Data
	logger *zap.Logger
}

func NewCommentRepo(
	data *data.Data,
) comment.CommentRepo {
	return &commentRepo{
		data: data,
	}
}

func (cr *commentRepo) AddComment(comment *entity.Comment) error {
	result := cr.data.DB.Create(comment)
	if result.Error != nil {
		return result.Error
	}
	return nil
}
