package entity

import "gorm.io/gorm"

type Moment struct {
	gorm.Model
	userID   int
	identity int
	content  string
	picture  []string
}
