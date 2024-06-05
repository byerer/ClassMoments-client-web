package schema

type UserLoginReq struct {
	Username string `validate:"required" json:"username"`
	Password string `validate:"required" json:"password"`
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
