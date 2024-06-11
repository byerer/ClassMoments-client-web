package schema

import "ClassMoments-client-web/internal/entity"

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
	Role      uint   `json:"role"`
	CreatTime string `json:"create_time"`
	Content   string `json:"content"`
	Image     string `json:"image"`
	//Image     []string `json:"image"`
}

type GetMomentsReq struct {
	ClassID uint `json:"classID"`
}

type Moment struct {
	AddMomentResp
	LikeCount    int              `json:"likeCount"`
	CommentCount int              `json:"commentCount"`
	CommentList  []entity.Comment `json:"commentList"`
}

type GetMomentsResp struct {
	MomentList []Moment `json:"momentList"`
}
