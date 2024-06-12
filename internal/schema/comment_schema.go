package schema

type CommentReq struct {
	MomentID uint   `json:"momentID"`
	UserID   uint   `json:"userID"`
	Content  string `json:"detail"`
}

type CommentResp struct {
	CommentID uint   `json:"momentID"`
	CreatedAt string `json:"momentTime"`
}
