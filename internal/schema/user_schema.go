package schema

type UserLoginReq struct {
	Name     string `json:"name"`
	Username string `json:"username"`
	Password string `json:"password"`
	ClassID  string `json:"class_id"`
	Role     string `json:"role"`
	Sex      string `json:"sex"`
}

type UserRegisterReq struct {
	Username string `json:"username"`
	Password string `json:"password"`
}

type UserLoginResp struct {
	Username string `json:"username"`
	Name     string `json:"name"`
	Sex      string
	Role     string
	ClassID  string
}

type UserRegisterResp struct {
	Username string `json:"username"`
}
