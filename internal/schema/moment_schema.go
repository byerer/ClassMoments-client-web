package schema

// AddMomentReq add moment request
type AddMomentReq struct {
	UserID  uint   `json:"userID"`
	ClassID uint   `json:"classID"`
	Content string `json:"content"`
	Image   string `json:"image"`
	//Image   []string `json:"image"`
}

// AddMomentResp add moment response
type AddMomentResp struct {
	MomentID  uint   `json:"momentID"`
	UserID    uint   `json:"userID"`
	Role      string `json:"role"`
	CreatTime string `json:"create_time"`
	Content   string `json:"content"`
	Image     string `json:"image"`
	//Image     []string `json:"image"`
}
