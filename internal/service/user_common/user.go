package user_common

import (
	"ClassMoments-client-web/internal/entity"
	"context"
)

type UserRepo interface {
	AddUser(ctx context.Context, user *entity.User) error
}

type UserCommon struct {
	UserRepo UserRepo
}

func NewUserCommon(userRepo UserRepo) *UserCommon {
	return &UserCommon{
		UserRepo: userRepo,
	}
}
