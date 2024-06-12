package entity

import (
	"time"
)

type Moment struct {
	MomentID uint `gorm:"primaryKey;autoIncrement"`
	UserID   uint
	ClassID  uint
	Role     uint
	Title    string
	Content  string
	Image    string
	//Image     []string `gorm:"type:text;serializer:json"`
	CreatedAt time.Time
}
