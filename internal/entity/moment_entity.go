package entity

import (
	"time"
)

type Moment struct {
	MomentID  uint
	UserID    uint
	Role      string
	Content   string
	Image     []string
	CreatedAt time.Time
}
