package entity

import "gorm.io/gorm"

type Moment struct {
	gorm.Model
	userID  uint
	role    string
	content string
	image   []string
}
