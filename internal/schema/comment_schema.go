package schema

type CommentReq struct {
	MomentID uint   `json:"postID"`
	UserID   uint   `json:"userID"`
	Content  string `json:"detail"`
}

type CommentResp struct {
	CommentID uint   `json:"postID"`
	CreatedAt string `json:"postTime"`
}
