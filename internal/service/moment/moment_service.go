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
		Role:    momentReq.Role,
		Content: momentReq.Content,
		Image:   momentReq.Image,
	}
	err := ms.momentRepo.AddMoment(moment)
	if err != nil {
		return nil, err
	}
	resp := &schema.AddMomentResp{}
	return resp, nil
}

func (ms *momentService) DeleteMoment(momentID uint) error {
	err := ms.momentRepo.DeleteMoment(momentID)
	if err != nil {
		return err
	}
	return nil
}
