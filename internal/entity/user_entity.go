package entity

import (
	"gorm.io/gorm"
)

type User struct {
	gorm.Model
	//ID        string `gorm:"primaryKey;autoIncrement"`
	Name     string
	Username string `gorm:"unique"`
	Password string
	ClassID  string
	Role     string
	Sex      string

	//CreatedAt time.Time
	//UpdatedAt time.Time
}
