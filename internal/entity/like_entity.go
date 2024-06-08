package entity

import "gorm.io/gorm"

type Like struct {
	gorm.Model
	LikerID  uint
	LikeeID  uint
	MomentID uint
}
