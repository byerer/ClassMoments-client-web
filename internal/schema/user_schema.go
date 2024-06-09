package schema

type UserLoginReq struct {
	Username string `json:"username"`
	Password string `json:"password"`
}

type UserLoginResp struct {
	UserID   uint   `json:"userID"`
	Sex      uint   `json:"sex"`
	Role     uint   `json:"role"`
	ClassID  uint   `json:"classID"`
	Username string `json:"username"`
	Name     string `json:"name"`
}

type UserRegisterReq struct {
	Name     string `json:"name"`
	ClassID  uint   `json:"classID"`
	Role     uint   `json:"role"`
	Sex      uint   `json:"sex"`
	Username string `json:"username"`
	Password string `json:"password"`
}

type UserRegisterResp struct {
	Username string `json:"username"`
	Name     string `json:"name"`
}
