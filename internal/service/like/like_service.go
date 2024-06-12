package like

import (
	"ClassMoments-client-web/internal/entity"
	"ClassMoments-client-web/internal/schema"
	"time"
)

type LikeRepo interface {
	AddLike(like *entity.Like) error
	DeleteLike(userID, momentID uint) error
	GetLikeCount(momentID uint) (int, error)
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

func (ls *likeService) Like(req *schema.LikeReq) (resp *schema.LikeResp, err error) {
	like := &entity.Like{
		LikerID:   req.LikerID,
		LikeeID:   req.LikeeID,
		MomentID:  req.MomentID,
		CreatedAt: time.Now(),
	}
	err = ls.LikeRepo.AddLike(like)
	resp = &schema.LikeResp{}
	if err != nil {
		resp.Liked = false
		return resp, err
	}
	resp.Liked = true
	resp.LikeCount, err = ls.LikeRepo.GetLikeCount(req.MomentID)
	if err != nil {
		return nil, err
	}
	return resp, nil
}

func (ls *likeService) UnLike(likerID, momentID uint) error {
	err := ls.LikeRepo.DeleteLike(likerID, momentID)
	if err != nil {
		return err
	}
	return nil
}
