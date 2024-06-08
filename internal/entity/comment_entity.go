package entity

import "time"

type Comment struct {
	CommentID uint `gorm:"primaryKey;autoIncrement"`
	UserID    int
	MomentID  int
	Content   string
	CreatedAt time.Time
}
