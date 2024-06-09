package moment

import (
	"ClassMoments-client-web/internal/entity"
	"ClassMoments-client-web/internal/schema"
)

type MomentRepo interface {
	AddMoment(moment *entity.Moment) error
	DeleteMoment(momentID uint) error
}

type MomentService interface {
	AddMoment(momentReq *schema.AddMomentReq) (*schema.AddMomentResp, error)
}

type momentService struct {
	momentRepo MomentRepo
}

func NewMomentService(momentRepo MomentRepo) MomentService {
	return &momentService{momentRepo: momentRepo}
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
