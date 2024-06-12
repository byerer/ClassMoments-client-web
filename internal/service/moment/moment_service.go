package moment

import (
	"ClassMoments-client-web/internal/entity"
	"ClassMoments-client-web/internal/schema"
	"ClassMoments-client-web/internal/service/comment"
	"ClassMoments-client-web/internal/service/like"
)

type MomentRepo interface {
	AddMoment(moment *entity.Moment) error
	DeleteMoment(momentID uint) error
	GetMomentList(classID uint) ([]entity.Moment, error)
	GetMomentDetail(momentID uint) (*entity.Moment, error)
}

type MomentService interface {
	AddMoment(momentReq *schema.AddMomentReq) (*schema.AddMomentResp, error)
	GetMomentList(classID uint) (*schema.GetMomentsResp, error)
	GetMomentDetail(momentID uint) (*schema.Moment, error)
}

type momentService struct {
	momentRepo  MomentRepo
	likeRepo    like.LikeRepo
	commentRepo comment.CommentRepo
}

func NewMomentService(
	momentRepo MomentRepo,
	likeRepo like.LikeRepo,
	commentRepo comment.CommentRepo,
) MomentService {
	return &momentService{
		momentRepo:  momentRepo,
		likeRepo:    likeRepo,
		commentRepo: commentRepo}
}

func (ms *momentService) AddMoment(momentReq *schema.AddMomentReq) (*schema.AddMomentResp, error) {
	moment := &entity.Moment{
		UserID:  momentReq.UserID,
		Content: momentReq.Content,
		Image:   momentReq.Image,
	}
	err := ms.momentRepo.AddMoment(moment)
	if err != nil {
		return nil, err
	}
	resp := &schema.AddMomentResp{
		MomentID:  moment.MomentID,
		UserID:    moment.UserID,
		Role:      moment.Role,
		CreatTime: moment.CreatedAt.Format("2006-01-02 15:04:05"),
		Content:   moment.Content,
		Image:     moment.Image,
	}
	return resp, nil
}

func (ms *momentService) DeleteMoment(momentID uint) error {
	err := ms.momentRepo.DeleteMoment(momentID)
	if err != nil {
		return err
	}
	return nil
}

func (ms *momentService) GetMomentList(classID uint) (*schema.GetMomentsResp, error) {
	moments, err := ms.momentRepo.GetMomentList(classID)
	if err != nil {
		return nil, err
	}
	resp := &schema.GetMomentsResp{}
	for _, moment := range moments {
		momentResp := schema.Moment{
			AddMomentResp: schema.AddMomentResp{
				MomentID:  moment.MomentID,
				UserID:    moment.UserID,
				Role:      moment.Role,
				CreatTime: moment.CreatedAt.Format("2006-01-02 15:04:05"),
				Content:   moment.Content,
				Image:     moment.Image,
			},
		}
		momentResp.LikeCount, err = ms.likeRepo.GetLikeCount(moment.MomentID)
		momentResp.CommentList, err = ms.commentRepo.GetCommentList(moment.MomentID)
		if err != nil {
			return nil, err
		}
		momentResp.CommentCount = len(momentResp.CommentList)
		resp.MomentList = append(resp.MomentList, momentResp)
	}
	return resp, nil
}

func (ms *momentService) GetMomentDetail(momentID uint) (*schema.Moment, error) {
	moment, err := ms.momentRepo.GetMomentDetail(momentID)
	if err != nil {
		return nil, err
	}
	resp := &schema.Moment{
		AddMomentResp: schema.AddMomentResp{
			MomentID:  moment.MomentID,
			UserID:    moment.UserID,
			Role:      moment.Role,
			CreatTime: moment.CreatedAt.Format("2006-01-02 15:04:05"),
			Content:   moment.Content,
			Image:     moment.Image,
		},
	}
	resp.LikeCount, err = ms.likeRepo.GetLikeCount(moment.MomentID)
	resp.CommentList, err = ms.commentRepo.GetCommentList(moment.MomentID)
	if err != nil {
		return nil, err
	}
	resp.CommentCount = len(resp.CommentList)
	return resp, nil
}
