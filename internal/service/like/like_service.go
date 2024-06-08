package like

import (
	"ClassMoments-client-web/internal/entity"
	"ClassMoments-client-web/internal/schema"
)

type LikeRepo interface {
	AddLike(like *entity.Like) error
	DeleteLike(userID, momentID uint) error
}

type LikeService interface {
	Like(req *schema.LikeReq) (*schema.LikeResp, error)
	UnLike(likerID, momentID uint) error
}

type likeService struct {
	LikeRepo LikeRepo
}

func NewLikeService(likeRepo LikeRepo) LikeService {
	return &likeService{LikeRepo: likeRepo}
}

func (ls *likeService) Like(req *schema.LikeReq) (*schema.LikeResp, error) {
	like := &entity.Like{
		LikerID:  req.LikerID,
		LikeeID:  req.LikeeID,
		MomentID: req.MomentID,
	}
	err := ls.LikeRepo.AddLike(like)
	if err != nil {
		return nil, err
	}
	resp := &schema.LikeResp{}
	return resp, nil
}

func (ls *likeService) UnLike(likerID, momentID uint) error {
	err := ls.LikeRepo.DeleteLike(likerID, momentID)
	if err != nil {
		return err
	}
	return nil
}
