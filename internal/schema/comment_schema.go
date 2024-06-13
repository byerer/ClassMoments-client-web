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

type CommentBase struct {
	CommentID uint   `json:"commentID"`
	UserID    uint   `json:"userID"`
	Username  string `json:"username"`
	Content   string `json:"content"`
	CreatedAt string `json:"createdAt"`
}

type CommentListResp struct {
	Comments []CommentBase `json:"comments"`
}
