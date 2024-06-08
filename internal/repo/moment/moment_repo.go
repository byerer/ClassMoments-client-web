package moment

import (
	"ClassMoments-client-web/internal/base/data"
	"ClassMoments-client-web/internal/entity"
	"ClassMoments-client-web/internal/service/moment"
)

type momentRepo struct {
	data *data.Data
}

func NewMomentRepo(data *data.Data) moment.MomentRepo {
	return &momentRepo{
		data: data,
	}
}

// AddMoment add moment
func (mr *momentRepo) AddMoment(moment *entity.Moment) error {
	result := mr.data.DB.Create(moment)

	if result.Error != nil {
		return result.Error
	}
	return nil
}

// DeleteMoment delete moment
func (mr *momentRepo) DeleteMoment(momentID uint) error {
	result := mr.data.DB.Where("id = ?", momentID).Delete(&entity.Moment{})
	if result.Error != nil {
		return result.Error
	}
	return nil
}
