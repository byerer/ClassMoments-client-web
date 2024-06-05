package user

import (
	"ClassMoments-client-web/internal/base/data"
	"ClassMoments-client-web/internal/entity"
	usercommon "ClassMoments-client-web/internal/service/user_common"
)

type userRepo struct {
	data *data.Data
}

func NewUserRepo(data *data.Data) usercommon.UserRepo {
	return &userRepo{
		data: data,
	}
}

func (ur *userRepo) AddUser(user *entity.User) error {
	result := ur.data.DB.Create(user)
	if result.Error != nil {
		return result.Error
	}
	return nil
}

func (ur *userRepo) GetUserByUsername(username string) (*entity.User, error) {
	var user entity.User
	result := ur.data.DB.Where("username = ?", username).First(&user)
	if result.Error != nil {
		return nil, result.Error
	}
	return &user, nil
}
