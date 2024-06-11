package entity

type UserActivities struct {
	ID          uint `gorm:"primary_key;autoIncrement"`
	UserID      uint
	ActivityID  uint
	Description string
	CreatedAt   string
	IPAddress   string
}
