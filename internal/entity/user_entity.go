package entity

import "time"

type User struct {
	UserID     uint `gorm:"primaryKey;autoIncrement"`
	Role       uint
	Sex        uint
	ClassID    uint
	SchoolName string
	Name       string
	Username   string `gorm:"unique"`
	Password   string
	CreatedAt  time.Time
	UpdatedAt  time.Time
}
