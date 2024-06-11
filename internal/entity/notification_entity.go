package entity

import "time"

const (
	Unread = 0
	Read   = 1
)

type Notification struct {
	ID       uint `gorm:"primary_key;auto_increment"`
	UserID   uint
	MomentID uint
	IsRead   int
	Type     string
	CreateAt time.Time
}
