package entity

import "time"

type Like struct {
	LikeID    uint `gorm:"primaryKey;autoIncrement"`
	LikerID   uint
	LikeeID   uint
	MomentID  uint
	CreatedAt time.Time
}
