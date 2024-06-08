package schema

type LikeReq struct {
	LikerID  uint `json:"liker_id"`
	LikeeID  uint `json:"likee_id"`
	MomentID uint `json:"moment_id"`
}

type LikeResp struct {
	Liked bool `json:"liked"`
}
