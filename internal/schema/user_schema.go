package schema

type UserLoginReq struct {
	Name     string `json:"name"`
	Username string `validate:"required" json:"username"`
	Password string `validate:"required" json:"password"`
	ClassID  string `json:"class_id"`
	Role     string `json:"role"`
	Sex      string `json:"sex"`
}

type UserRegisterReq struct {
	Username string `validate:"required" json:"username"`
	Password string `validate:"required" json:"password"`
}

type UserLoginResp struct {
	Username string `json:"username"`
}

type UserRegisterResp struct {
	Username string `json:"username"`
}
