package entity

import "time"

type Notification struct {
	ID       uint `gorm:"primary_key;auto_increment"`
	UserID   uint
	MomentID uint
	Type     string
	CreateAt time.Time
}
