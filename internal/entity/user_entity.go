package entity

import "time"

type User struct {
	UserID    uint `gorm:"primaryKey;autoIncrement"`
	Name      string
	Username  string `gorm:"unique"`
	Password  string
	ClassID   string
	Role      string
	Sex       string
	CreatedAt time.Time
	UpdatedAt time.Time
}
