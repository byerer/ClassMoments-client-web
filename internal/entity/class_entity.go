package entity

type Class struct {
	ID               uint `gorm:"primary_key;auto_increment"`
	ClassName        string
	SchoolName       string
	ClassDescription string
}
