package schema

// AddMomentReq add moment request
type AddMomentReq struct {
	UserID  uint   `json:"userID"`
	ClassID uint   `json:"classID"`
	Title   string `json:"title"`
	Content string `json:"content"`
	Image   string `json:"image"`
	//Image   []string `json:"image"`
}

// AddMomentResp moment response
type MomentBase struct {
	MomentID  uint   `json:"momentID"`
	UserID    uint   `json:"userID"`
	Role      uint   `json:"role"`
	ClassID   uint   `json:"classID"`
	Username  string `json:"username"`
	Title     string `json:"title"`
	CreatTime string `json:"createTime"`
	Content   string `json:"content"`
	Image     string `json:"image"`
	//Image     []string `json:"image"`
}

type GetMomentsReq struct {
	ClassID uint `json:"classID"`
}

type MomentResp struct {
	MomentBase
	LikeCount    int           `json:"likeCount"`
	CommentCount int           `json:"commentCount"`
	CommentList  []CommentBase `json:"commentList"`
}

type MomentsResp struct {
	MomentList []MomentResp `json:"momentList"`
}
