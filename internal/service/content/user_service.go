package content

import (
	usercommon "ClassMoments-client-web/internal/service/user_common"
)

type UserService struct {
	userCommonService *usercommon.UserCommon
	userRepo          usercommon.UserRepo
}

func NewUserService(
	userCommonService *usercommon.UserCommon,
	userRepo usercommon.UserRepo,
) *UserService {
	return &UserService{
		userCommonService: userCommonService,
		userRepo:          userRepo,
	}
}
