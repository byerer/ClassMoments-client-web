package moment

import (
	"ClassMoments-client-web/internal/entity"
	"ClassMoments-client-web/internal/schema"
	"ClassMoments-client-web/internal/service/comment"
	"ClassMoments-client-web/internal/service/like"
	usercommon "ClassMoments-client-web/internal/service/user_common"
	"fmt"
)

type MomentRepo interface {
	AddMoment(moment *entity.Moment) error
	DeleteMoment(momentID uint) error
	GetMomentList(classID uint) ([]entity.Moment, error)
	GetMomentDetail(momentID uint) (*entity.Moment, error)
}

type MomentService interface {
	AddMoment(momentReq *schema.AddMomentReq) (*schema.MomentBase, error)
	GetMomentList(classID uint) (*schema.MomentsResp, error)
	GetMomentDetail(momentID uint) (*schema.MomentResp, error)
}

type momentService struct {
	momentRepo     MomentRepo
	likeRepo       like.LikeRepo
	userRepo       usercommon.UserRepo
	commentService comment.CommentService
}

func NewMomentService(
	momentRepo MomentRepo,
	likeRepo like.LikeRepo,
	userRepo usercommon.UserRepo,
	commentService comment.CommentService,
) MomentService {
	return &momentService{
		momentRepo:     momentRepo,
		likeRepo:       likeRepo,
		userRepo:       userRepo,
		commentService: commentService,
	}
}

func (ms *momentService) AddMoment(momentReq *schema.AddMomentReq) (*schema.MomentBase, error) {
	moment := &entity.Moment{
		UserID:  momentReq.UserID,
		ClassID: momentReq.ClassID,
		Content: momentReq.Content,
		Image:   momentReq.Image,
	}
	err := ms.momentRepo.AddMoment(moment)
	if err != nil {
		return nil, err
	}
	resp := &schema.MomentBase{
		MomentID:  moment.MomentID,
		UserID:    moment.UserID,
		CreatTime: moment.CreatedAt.Format("2006-01-02 15:04:05"),
		Content:   moment.Content,
		Image:     moment.Image,
	}
	user, err := ms.userRepo.GetUserByID(moment.UserID)
	if err != nil {
		return nil, err
	}
	resp.Username = user.Username
	return resp, nil
}

func (ms *momentService) DeleteMoment(momentID uint) error {
	err := ms.momentRepo.DeleteMoment(momentID)
	if err != nil {
		return err
	}
	return nil
}

func (ms *momentService) GetMomentList(classID uint) (*schema.MomentsResp, error) {
	moments, err := ms.momentRepo.GetMomentList(classID)
	if err != nil {
		return nil, err
	}
	resp := &schema.MomentsResp{}
	for _, moment := range moments {
		momentResp := schema.MomentResp{
			MomentBase: schema.MomentBase{
				MomentID:  moment.MomentID,
				UserID:    moment.UserID,
				ClassID:   moment.ClassID,
				Role:      moment.Role,
				Title:     moment.Title,
				CreatTime: moment.CreatedAt.Format("2006-01-02 15:04:05"),
				Content:   moment.Content,
				Image:     moment.Image,
			},
		}
		user, err := ms.userRepo.GetUserByID(moment.UserID)
		if err != nil {
			return nil, err
		}
		momentResp.Username = user.Username
		momentResp.LikeCount, err = ms.likeRepo.GetLikeCount(moment.MomentID)
		CommentListResp, err := ms.commentService.GetCommentList(moment.MomentID)
		momentResp.CommentList = CommentListResp.Comments
		if err != nil {
			return nil, err
		}
		momentResp.CommentCount = len(momentResp.CommentList)
		resp.MomentList = append(resp.MomentList, momentResp)
		fmt.Println("momentResp:", momentResp)
	}
	return resp, nil
}

func (ms *momentService) GetMomentDetail(momentID uint) (*schema.MomentResp, error) {
	moment, err := ms.momentRepo.GetMomentDetail(momentID)
	if err != nil {
		return nil, err
	}
	fmt.Println("moment:", moment)
	resp := &schema.MomentResp{
		MomentBase: schema.MomentBase{
			MomentID:  moment.MomentID,
			UserID:    moment.UserID,
			ClassID:   moment.ClassID,
			Role:      moment.Role,
			Title:     moment.Title,
			CreatTime: moment.CreatedAt.Format("2006-01-02 15:04:05"),
			Content:   moment.Content,
			Image:     moment.Image,
		},
	}
	user, err := ms.userRepo.GetUserByID(moment.UserID)
	if err != nil {
		return nil, err
	}
	resp.Username = user.Username
	resp.LikeCount, err = ms.likeRepo.GetLikeCount(moment.MomentID)
	if err != nil {
		return nil, err
	}
	commentListResp := &schema.CommentListResp{}
	commentListResp, err = ms.commentService.GetCommentList(moment.MomentID)
	if err != nil {
		return nil, err
	}
	resp.CommentList = commentListResp.Comments
	resp.CommentCount = len(resp.CommentList)
	return resp, nil
}
