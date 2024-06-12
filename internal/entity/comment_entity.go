package entity

import "time"

type Comment struct {
	CommentID uint `gorm:"primaryKey;autoIncrement"`
	UserID    uint
	MomentID  uint
	Username  string
	Content   string
	CreatedAt time.Time
}
