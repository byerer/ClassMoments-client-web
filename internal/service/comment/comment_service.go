package comment

import (
	"ClassMoments-client-web/internal/entity"
	"ClassMoments-client-web/internal/schema"
	"go.uber.org/zap"
)

type CommentRepo interface {
	AddComment(comment *entity.Comment) error
	GetCommentList(momentID uint) (comments []entity.Comment, err error)
}

type CommentService interface {
	AddComment(comment *schema.CommentReq) (schema.CommentResp, error)
}

type commentService struct {
	commentRepo CommentRepo
	logger      *zap.Logger
}

func NewCommentService(
	commentRepo CommentRepo,
	logger *zap.Logger,
) CommentService {
	return &commentService{
		commentRepo: commentRepo,
		logger:      logger,
	}
}

func (cs *commentService) AddComment(comment *schema.CommentReq) (schema.CommentResp, error) {
	commentEntity := &entity.Comment{
		UserID:   comment.UserID,
		MomentID: comment.MomentID,
		Content:  comment.Content,
	}
	resp := schema.CommentResp{}
	err := cs.commentRepo.AddComment(commentEntity)
	if err != nil {
		cs.logger.Error("add comment failed", zap.Error(err))
		return resp, err
	}
	resp.CommentID = commentEntity.CommentID
	resp.CreatedAt = commentEntity.CreatedAt.Format("2006-01-02 15:04:05")
	return resp, nil
}
