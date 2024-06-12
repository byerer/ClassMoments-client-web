package entity

type School struct {
	SchoolID   uint `gorm:"primary_key;auto_increment"`
	SchoolName string
}
