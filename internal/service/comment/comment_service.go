package comment

import (
	"ClassMoments-client-web/internal/entity"
	"ClassMoments-client-web/internal/schema"
	usercommon "ClassMoments-client-web/internal/service/user_common"
	"go.uber.org/zap"
)

type CommentRepo interface {
	AddComment(comment *entity.Comment) error
	GetCommentList(momentID uint) ([]entity.Comment, error)
}

type CommentService interface {
	AddComment(comment *schema.CommentReq) (*schema.CommentResp, error)
	GetCommentList(momentID uint) (*schema.CommentListResp, error)
}

type commentService struct {
	commentRepo CommentRepo
	userRepo    usercommon.UserRepo
	logger      *zap.Logger
}

func NewCommentService(
	commentRepo CommentRepo,
	userRepo usercommon.UserRepo,
	logger *zap.Logger,
) CommentService {
	return &commentService{
		commentRepo: commentRepo,
		logger:      logger,
		userRepo:    userRepo,
	}
}

func (cs *commentService) AddComment(comment *schema.CommentReq) (*schema.CommentResp, error) {
	commentEntity := &entity.Comment{
		UserID:   comment.UserID,
		MomentID: comment.MomentID,
		Content:  comment.Content,
	}
	resp := &schema.CommentResp{}
	err := cs.commentRepo.AddComment(commentEntity)
	if err != nil {
		cs.logger.Error("add comment failed", zap.Error(err))
		return nil, err
	}
	resp.CommentID = commentEntity.CommentID
	resp.CreatedAt = commentEntity.CreatedAt.Format("2006-01-02 15:04:05")
	return resp, nil
}

func (cs *commentService) GetCommentList(momentID uint) (resp *schema.CommentListResp, err error) {
	comments, err := cs.commentRepo.GetCommentList(momentID)
	for _, comment := range comments {
		commentBase := schema.CommentBase{
			CommentID: comment.CommentID,
			UserID:    comment.UserID,
			Content:   comment.Content,
			CreatedAt: comment.CreatedAt.Format("2006-01-02 15:04:05"),
		}
		user, _ := cs.userRepo.GetUserByID(comment.UserID)
		commentBase.Username = user.Username
		resp.Comments = append(resp.Comments, commentBase)
	}
	if err != nil {
		cs.logger.Error("get comment list failed", zap.Error(err))
		return nil, err
	}
	return resp, nil
}
