package entity

import "gorm.io/gorm"

type Like struct {
	gorm.Model
	UserID   int
	MomentID int
}
