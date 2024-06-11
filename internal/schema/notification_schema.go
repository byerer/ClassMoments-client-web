package schema

type NotificationReq struct {
	UserID   uint   `json:"userID"`
	MomentID uint   `json:"momentID"`
	Type     string `json:"type"`
}

type NotificationResp struct {
}
