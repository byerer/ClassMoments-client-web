package schema

// AddMomentReq add moment request
type AddMomentReq struct {
	UserID  uint     `json:"userID"`
	Role    string   `json:"role"`
	Content string   `json:"content"`
	Image   []string `json:"image"`
}

// AddMomentResp add moment response
type AddMomentResp struct {
	MomentID  uint     `json:"moment_id"`
	UserID    uint     `json:"user_id"`
	Role      string   `json:"role"`
	CreatTime string   `json:"create_time"`
	Content   string   `json:"content"`
	Image     []string `json:"image"`
}
