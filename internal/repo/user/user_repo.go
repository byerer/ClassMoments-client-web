package user

import (
	"ClassMoments-client-web/internal/base/data"
	"ClassMoments-client-web/internal/entity"
	usercommon "ClassMoments-client-web/internal/service/user_common"
	"context"
)

type userRepo struct {
	data *data.Data
}

func NewUserRepo(data *data.Data) usercommon.UserRepo {
	return &userRepo{
		data: data,
	}
}

func (ur *userRepo) AddUser(ctx context.Context, user *entity.User) error {
	result := ur.data.DB.Create(user)
	if result.Error != nil {
		return result.Error
	}
	return nil
}
