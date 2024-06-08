package entity

import (
	"time"
)

type Moment struct {
	MomentID  uint `gorm:"primaryKey;autoIncrement"`
	UserID    uint
	Role      string
	Content   string
	Image     []string `gorm:"type:text;serializer:json"`
	CreatedAt time.Time
}
