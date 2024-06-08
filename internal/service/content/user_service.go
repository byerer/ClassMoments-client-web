package content

import (
	"ClassMoments-client-web/internal/entity"
	"ClassMoments-client-web/internal/schema"
	usercommon "ClassMoments-client-web/internal/service/user_common"
	"time"
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

func (us *UserService) UserLogin(req *schema.UserLoginReq) (*schema.UserLoginResp, error) {
	user, err := us.userRepo.GetUserByUsername(req.Username)
	if err != nil {
		return nil, err
	}
	if user.Password != req.Password {
		return nil, nil
	}
	return &schema.UserLoginResp{
		Username: user.Username,
		Name:     user.Name,
		Sex:      user.Sex,
		Role:     user.Role,
		ClassID:  user.ClassID,
	}, nil
}

func (us *UserService) UserRegister(req *schema.UserRegisterReq) error {
	user := &entity.User{
		Name:      req.Name,
		Username:  req.Username,
		Password:  req.Password,
		ClassID:   req.ClassID,
		Role:      req.Role,
		Sex:       req.Sex,
		CreatedAt: time.Now(),
	}
	return us.userRepo.AddUser(user)
}
