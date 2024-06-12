package schema

type LikeReq struct {
	LikerID  uint `json:"likerID"`
	LikeeID  uint `json:"likeeID"`
	MomentID uint `json:"momentID"`
}

type LikeResp struct {
	Liked     bool `json:"liked"`
	LikeCount int  `json:"likeCount"`
}
