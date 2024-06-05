package schema

type UserLoginReq struct {
	Username string `validate:"required" json:"username"`
	Password string `validate:"required" json:"password"`
}

type UserRegisterReq struct {
	Username string `validate:"required" json:"username"`
	Password string `validate:"required" json:"password"`
}
