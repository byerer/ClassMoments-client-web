package entity

import "gorm.io/gorm"

type Comment struct {
	gorm.Model
	UserID   int
	MomentID int
	Content  string
}
