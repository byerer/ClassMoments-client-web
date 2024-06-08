package schema

type UserLoginReq struct {
	Username string `json:"username"`
	Password string `json:"password"`
}

type UserRegisterReq struct {
	Name     string `json:"name"`
	Username string `json:"username"`
	Password string `json:"password"`
	ClassID  string `json:"class_id"`
	Role     string `json:"role"`
	Sex      string `json:"sex"`
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
	Name     string `json:"name"`
}
