package user_common

import (
	"ClassMoments-client-web/internal/entity"
)

type UserRepo interface {
	AddUser(user *entity.User) error
	GetUserByUsername(username string) (*entity.User, error)
}

type UserCommon struct {
	UserRepo UserRepo
}

func NewUserCommon(userRepo UserRepo) *UserCommon {
	return &UserCommon{
		UserRepo: userRepo,
	}
}
