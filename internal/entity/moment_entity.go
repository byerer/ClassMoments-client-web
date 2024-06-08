package entity

import "gorm.io/gorm"

type Moment struct {
	gorm.Model
	UserID  uint
	Role    string
	Content string
	Image   []string
}
