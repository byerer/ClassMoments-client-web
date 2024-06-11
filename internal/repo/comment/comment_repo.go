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

func (cr *commentRepo) DeleteComment(commentID uint) error {
	result := cr.data.DB.Where("id = ?", commentID).Delete(&entity.Comment{})
	if result.Error != nil {
		return result.Error
	}
	return nil
}

func (cr *commentRepo) GetCommentList(momentID uint) ([]entity.Comment, error) {
	var comments []entity.Comment
	result := cr.data.DB.Where("moment_id = ?", momentID).Find(&comments)
	if result.Error != nil {
		return nil, result.Error
	}
	return comments, nil
}
